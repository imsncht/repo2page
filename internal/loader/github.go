package loader

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GitHubFileEntry represents a file returned from GitHub.
type GitHubFileEntry struct {
	Path string
	Size int64
	SHA  string
}

// githubTreeResponse models the GitHub Trees API response.
type githubTreeResponse struct {
	Tree []struct {
		Path string `json:"path"`
		Type string `json:"type"` // "blob" or "tree"
		SHA  string `json:"sha"`
		Size int64  `json:"size,omitempty"`
	} `json:"tree"`
	Truncated bool `json:"truncated"`
}

// LoadGitHubRepository fetches a GitHub repository tree (recursive).
// Supports branch or commit selection.
// Public repositories only (private support is post-MVP).
func LoadGitHubRepository(owner, repo, ref string) ([]GitHubFileEntry, error) {
	if owner == "" || repo == "" {
		return nil, errors.New("invalid GitHub repository identifier")
	}

	if ref == "" {
		ref = "main"
	}

	apiURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1",
		owner,
		repo,
		ref,
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Optional token for higher rate limits (public repos only)
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "repo2page")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed to reach GitHub API")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// continue
	case http.StatusNotFound:
		return nil, errors.New("GitHub repository or ref not found")
	case http.StatusForbidden:
		return nil, errors.New("GitHub API access forbidden or rate-limited")
	default:
		return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var tree githubTreeResponse
	if err := json.NewDecoder(resp.Body).Decode(&tree); err != nil {
		return nil, err
	}

	if tree.Truncated {
		return nil, errors.New("GitHub tree response truncated; repository too large")
	}

	var files []GitHubFileEntry

	for _, node := range tree.Tree {
		if node.Type != "blob" {
			continue
		}

		path := strings.TrimSpace(node.Path)
		if path == "" {
			continue
		}

		files = append(files, GitHubFileEntry{
			Path: path,
			Size: node.Size,
			SHA:  node.SHA,
		})
	}

	if len(files) == 0 {
		return nil, errors.New("no readable files found in GitHub repository")
	}

	return files, nil
}

func FetchDefaultBranch(owner, repo string) (string, error) {
	apiURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s",
		owner,
		repo,
	)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "repo2page")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch repository metadata")
	}

	var data struct {
		DefaultBranch string `json:"default_branch"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.DefaultBranch == "" {
		return "", errors.New("default branch not found")
	}

	return data.DefaultBranch, nil
}

// ReadGitHubFile fetches and decodes a file's content from GitHub.
func ReadGitHubFile(owner, repo, path, ref string) (string, int, error) {
	rawURL := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/%s/%s",
		owner,
		repo,
		ref,
		path,
	)

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Get(rawURL)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("failed to fetch GitHub file: %s", path)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	content := string(data)
	lines := strings.Count(content, "\n") + 1

	return content, lines, nil
}

// DownloadRepoAsArchive downloads a GitHub repository as a tarball and extracts it.
// Returns the path to the extracted directory and a cleanup function.
// This is much more efficient than fetching files individually.
func DownloadRepoAsArchive(owner, repo, ref string, progressCallback func(int64, io.Reader) io.Reader) (extractedPath string, cleanup func(), err error) {
	if owner == "" || repo == "" {
		return "", nil, errors.New("invalid GitHub repository identifier")
	}

	if ref == "" {
		ref = "main"
	}

	// GitHub tarball URL
	tarballURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/tarball/%s",
		owner,
		repo,
		ref,
	)

	client := &http.Client{
		Timeout: 120 * time.Second, // Longer timeout for large repos
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// GitHub redirects to codeload.github.com - follow it
			return nil
		},
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, tarballURL, nil)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Optional token for higher rate limits and private repos
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "repo2page")

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("failed to download tarball: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// continue
	case http.StatusNotFound:
		return "", nil, fmt.Errorf("repository or branch not found (%s)", ref)
	default:
		return "", nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	tempDir, err := os.MkdirTemp("", "repo2page-gh-")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp dir: %w", err)
	}

	cleanup = func() {
		os.RemoveAll(tempDir)
	}

	var r io.Reader = resp.Body
	if progressCallback != nil {
		r = progressCallback(resp.ContentLength, resp.Body)
	}

	// Extract
	extractedRoot, err := extractTarGz(r, tempDir)
	if err != nil {
		cleanup() // Clean up on failure
		return "", nil, fmt.Errorf("failed to extract archive: %w", err)
	}

	return extractedRoot, cleanup, nil
}

// extractTarGz extracts a gzip-compressed tar archive to the destination directory.
// Returns the path to the root directory inside the archive (GitHub adds a prefix dir).
func extractTarGz(r io.Reader, destDir string) (string, error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	var rootDir string

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read tar header: %w", err)
		}

		// Skip PAX global headers and other special entries
		if header.Typeflag == tar.TypeXGlobalHeader || header.Typeflag == tar.TypeXHeader {
			continue
		}

		// Skip pax_global_header file (sometimes appears as a regular file)
		if header.Name == "pax_global_header" || strings.HasPrefix(header.Name, "pax_global_header") {
			continue
		}

		// GitHub tarballs have a root directory like "owner-repo-commit/"
		// We need to track it to return the correct path
		// Only set rootDir from actual directories, not special files
		if header.Typeflag == tar.TypeDir || header.Typeflag == tar.TypeReg {
			parts := strings.SplitN(header.Name, "/", 2)
			if rootDir == "" && len(parts) > 0 && parts[0] != "" {
				rootDir = parts[0]
			}
		}

		// Sanitize path to prevent path traversal attacks
		target := filepath.Join(destDir, header.Name)
		if !strings.HasPrefix(target, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return "", fmt.Errorf("invalid tar path: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return "", fmt.Errorf("failed to create directory: %w", err)
			}

		case tar.TypeReg:
			// Create parent directories if needed
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return "", fmt.Errorf("failed to create parent directory: %w", err)
			}

			// Create file with safe permissions
			mode := os.FileMode(header.Mode)
			if mode == 0 {
				mode = 0644 // Default to readable file
			}

			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
			if err != nil {
				return "", fmt.Errorf("failed to create file: %w", err)
			}

			// Copy contents
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return "", fmt.Errorf("failed to write file: %w", err)
			}
			f.Close()

		case tar.TypeSymlink:
			// Skip symlinks for security
			continue

		default:
			// Skip other types (devices, etc.)
			continue
		}
	}

	if rootDir == "" {
		return "", errors.New("empty tarball")
	}

	return filepath.Join(destDir, rootDir), nil
}
