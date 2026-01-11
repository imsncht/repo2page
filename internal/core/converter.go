package core

import "time"
import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"repo2page/internal/ignore"
	"repo2page/internal/loader"
)

// Convert is the primary entry point for repository conversion.
func Convert(input RepoInput, options ConvertOptions) (*ConvertResult, error) {
//	start := nowUTC()

	// 1. Detect source type (if not provided)
	if input.Source == SourceUnset {
	detected, err := AutoDetectSource(input.PathOrURL)
	if err != nil {
		return nil, err
	}
	input.Source = detected
}

	// 2. Load repository file list
	var (
		paths    []string
		rootPath string
		source   string
	)

	var (
		owner string
		repo  string
		ref   string
	)


	switch input.Source {
	case SourceLocal:
		rootPath = filepath.Clean(input.PathOrURL)
		files, err := loader.LoadLocalRepository(rootPath)
		if err != nil {
			return nil, err
		}

		for _, f := range files {
			paths = append(paths, f.Path)
		}
		source = rootPath

	case SourceGitHub:
		if !IsOnline() {
			return nil, errors.New("GitHub repositories require an internet connection")
		}

		var err error
		owner, repo, err := parseGitHubRepo(input.PathOrURL)
		if err != nil {
			return nil, err
		}

//		ref = input.Commit
//		if ref == "" {
//			ref = input.Branch
//		}
		
		ref = input.Commit
		if ref == "" {
			ref = input.Branch
		}
		if ref == "" {
			ref = "main"
		}


		files, err := loader.LoadGitHubRepository(owner, repo, ref)
		if err != nil {
			return nil, err
		}

		for _, f := range files {
			paths = append(paths, f.Path)
		}

		source = fmt.Sprintf("https://github.com/%s/%s", owner, repo)
		rootPath = "" // GitHub files are fetched virtually

	default:
		return nil, errors.New("unsupported source type")
	}

	if len(paths) == 0 {
		return nil, errors.New("no files found to convert")
	}

	// 3. Deterministic ordering
	tree := BuildTree(paths)
	orderedPaths := FlattenTree(tree)

	// 4. Ignore engine
	var ign *ignore.Engine
	if input.Source == SourceLocal {
		engine, err := ignore.NewEngine(rootPath, options.ExcludePatterns)
		if err != nil {
			return nil, err
		}
		ign = engine
	}

	// 5. Assemble
	assembled := Assemble(AssembleInput{
		RepoName:     repoNameFromSource(input.PathOrURL),
		Source:       source,
		RootPath:     rootPath,
		Files:        orderedPaths,
		Options:      options,
		GitHubOwner:  owner,
		GitHubRepo:   repo,
		GitHubRef:    ref,
		IsGitHub:     input.Source == SourceGitHub,
	}, ign)


	// 6. Format output
//	var content string

	

	// 7. Final result
return &ConvertResult{
	Format:    options.Format,
	Stats:     assembled.Stats,
	Warnings:  assembled.Warnings,
	Tree:      assembled.Tree,
	Files:     assembled.Files,
	Source:    source,
	RepoName:  repoNameFromSource(input.PathOrURL),
}, nil
}

// --------------------
// Helpers
// --------------------

func parseGitHubRepo(input string) (owner string, repo string, err error) {
	s := strings.TrimPrefix(input, "https://github.com/")
	parts := strings.Split(s, "/")
	if len(parts) < 2 {
		return "", "", errors.New("invalid GitHub repository format")
	}
	return parts[0], parts[1], nil
}

func repoNameFromSource(source string) string {
	s := strings.TrimSuffix(source, "/")
	if strings.Contains(s, "/") {
		return filepath.Base(s)
	}
	return s
}

// Time helpers (isolated for testability)
func nowUTC() int64 {
	return timeNow().UnixNano()
}

func elapsedMs(start int64) int64 {
	return (timeNow().UnixNano() - start) / int64(1e6)
}

var timeNow = func() time.Time {
	return time.Now().UTC()
}
