package loader

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FileEntry represents a discovered file in a local repository.
type FileEntry struct {
	Path string // Relative path from repo root
	Size int64  // Size in bytes
}

// LoadLocalRepository walks a local repository path and returns file entries.
// It does NOT read file contents.
func LoadLocalRepository(root string) ([]FileEntry, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("provided path is not a directory")
	}

	root = filepath.Clean(root)

	var files []FileEntry

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			// Permission denied or similar — skip gracefully
			return nil
		}

		// Skip symlinks entirely (security + determinism)
		if d.Type()&os.ModeSymlink != 0 {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Get file info
		info, err := d.Info()
		if err != nil {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return nil
		}

		// Normalize to forward slashes for cross-platform determinism
		relPath = filepath.ToSlash(relPath)

		// Ignore empty or hidden traversal artifacts
		if strings.TrimSpace(relPath) == "" {
			return nil
		}

		files = append(files, FileEntry{
			Path: relPath,
			Size: info.Size(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("no readable files found in directory")
	}

	return files, nil
}
