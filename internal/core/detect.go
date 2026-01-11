package core

import (
	"errors"
	"context"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// AutoDetectSource determines whether the input is a local path or a GitHub repository.
func AutoDetectSource(input string) (SourceType, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return SourceLocal, errors.New("empty source input")
	}

	// 1. Full URL (https://github.com/user/repo)
	if isGitHubURL(input) {
		return SourceGitHub, nil
	}

	// 2. GitHub shorthand (user/repo)
	if isGitHubShorthand(input) {
		return SourceGitHub, nil
	}

	// 3. Local filesystem path
	if existsAsPath(input) {
		return SourceLocal, nil
	}

	return SourceLocal, errors.New("unable to determine source type: invalid path or GitHub repository")
}

// IsOnline checks whether GitHub is reachable.
// This is a lightweight DNS-based check with a hard timeout.
func IsOnline() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resolver := net.Resolver{}
	_, err := resolver.LookupIPAddr(ctx, "github.com")
	return err == nil
}


// --------------------
// Helper functions
// --------------------

func isGitHubURL(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	return strings.EqualFold(u.Host, "github.com")
}

func isGitHubShorthand(input string) bool {
	// Matches: user/repo
	if strings.Count(input, "/") != 1 {
		return false
	}
	parts := strings.Split(input, "/")
	if parts[0] == "" || parts[1] == "" {
		return false
	}
	// Prevent false positives like ./path/file
	return !strings.Contains(parts[0], ".") && !strings.Contains(parts[1], ".")
}

func existsAsPath(path string) bool {
	_, err := os.Stat(filepath.Clean(path))
	return err == nil
}

// withTimeout creates a cancellable context with timeout.
// Defined locally to avoid importing context everywhere else.
func withTimeout(d time.Duration) (func() <-chan struct{}, func()) {
	done := make(chan struct{})
	timer := time.AfterFunc(d, func() {
		close(done)
	})
	cancel := func() {
		if timer.Stop() {
			close(done)
		}
	}
	return func() <-chan struct{} { return done }, cancel
}
