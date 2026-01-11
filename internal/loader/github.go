package loader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"io"

//	"encoding/base64"
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

