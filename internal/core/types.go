package core

import (
	"io"
	"time"
)

// SourceType represents where the repository is loaded from.
type SourceType int

const (
	SourceUnset SourceType = iota
	SourceLocal
	SourceGitHub
)

// RepoInput defines the repository source and selection parameters.
type RepoInput struct {
	Source    SourceType // Local filesystem or GitHub
	PathOrURL string     // Filesystem path or GitHub URL/shorthand
	Branch    string     // Optional, default: "main"
	Commit    string     // Optional, specific commit hash
}

// OutputFormat defines the supported output formats.
type OutputFormat string

const (
	FormatMarkdown OutputFormat = "md"
	FormatHTML     OutputFormat = "html"
	FormatText     OutputFormat = "txt"
	FormatJSON     OutputFormat = "json"
)

// ConvertOptions controls conversion behavior.
type ConvertOptions struct {
	Format             OutputFormat // md, html, txt, json
	IncludeTree        bool         // Include directory tree (default: true)
	IncludeReadmeFirst bool         // Render README first if present (default: true)
	ExcludePatterns    []string     // Additional ignore patterns
	MaxFileSizeKB      int          // Max file size per file (default: 500)
	SummaryMode        bool         // Structure-only mode (no file contents)

	OutputPath         string       // Optional output file path
	ProgressCallback   func(int64, io.Reader) io.Reader // Optional callback for tracking progress
}

// ConvertResult is the final output of a conversion operation.
type ConvertResult struct {
	Content  string      // Final rendered document
	Format   OutputFormat
	Stats    Statistics
	Warnings []string    // Non-fatal issues encountered during processing
	
	// NEW (raw data for formatting)
	Tree  string
	Files []RenderedFile
	Source string
	RepoName string	
}

// Statistics contains conversion metrics.
type Statistics struct {
	FileCount        int           // Number of files included
	TotalLines       int           // Total lines across all files
	TotalSizeKB      int           // Total size of included files
	SkippedFiles     int           // Number of skipped files
	ProcessingTimeMs int64         // Total processing time in milliseconds
	GeneratedAt      time.Time     // Timestamp of generation
}

// RenderedFile represents a file with content ready for formatting.
type RenderedFile struct {
	Path    string
	Content string
}
