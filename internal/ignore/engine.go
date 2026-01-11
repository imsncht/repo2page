package ignore

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Engine evaluates ignore rules against paths.
type Engine struct {
	patterns []string
}

// NewEngine creates a new ignore engine.
// Priority:
// 1. Default ignore patterns
// 2. .gitignore patterns
// 3. User-provided patterns (highest)
func NewEngine(root string, userPatterns []string) (*Engine, error) {
	var patterns []string

	// 1. Defaults
	patterns = append(patterns, DefaultIgnorePatterns...)

	// 2. .gitignore (if exists)
	gitignorePath := filepath.Join(root, ".gitignore")
	if _, err := os.Stat(gitignorePath); err == nil {
		gitPatterns, err := parseGitignore(gitignorePath)
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, gitPatterns...)
	}

	// 3. User overrides
	patterns = append(patterns, userPatterns...)

	return &Engine{patterns: patterns}, nil
}

// ShouldIgnore returns true if the path matches any ignore pattern.
func (e *Engine) ShouldIgnore(path string) bool {
	normalized := filepath.ToSlash(path)

	for _, pattern := range e.patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" || strings.HasPrefix(pattern, "#") {
			continue
		}

		// Directory ignore
		if strings.HasSuffix(pattern, "/") {
			if strings.HasPrefix(normalized, strings.TrimSuffix(pattern, "/")) {
				return true
			}
			continue
		}

		matched, err := filepath.Match(pattern, normalized)
		if err == nil && matched {
			return true
		}
	}

	return false
}

// --------------------
// Helpers
// --------------------

func parseGitignore(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Negation (!pattern) not supported in MVP
		if strings.HasPrefix(line, "!") {
			continue
		}

		patterns = append(patterns, line)
	}

	return patterns, scanner.Err()
}
