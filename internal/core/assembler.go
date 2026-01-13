package core

import (
	"path/filepath"
	"strings"
	"time"

	"repo2page/internal/ignore"
)

// AssembleInput represents all data needed by formatters.
type AssembleInput struct {
	RepoName string
	Source   string
	RootPath string   // Local path to repository root (also used for extracted GitHub repos)
	Files    []string // Ordered file paths (from tree resolver)
	Options  ConvertOptions
}

// AssembleResult is the assembled, formatter-ready output.
type AssembleResult struct {
	Tree     string
	Files    []RenderedFile
	Stats    Statistics
	Warnings []string
}

// Assemble builds tree text, reads files, applies ignores, and collects stats.
func Assemble(input AssembleInput, ign *ignore.Engine) AssembleResult {
	var (
		renderedFiles []RenderedFile
		warnings      []string
		totalLines    int
		totalSizeKB   int
		skipped       int
	)

	start := time.Now()

	// Build tree text
	treeText := buildTreeText(input.Files)

	for _, relPath := range input.Files {
		if ign != nil && ign.ShouldIgnore(relPath) {
			skipped++
			continue
		}

		// Read file from local filesystem
		// (GitHub repos are now extracted locally via tarball)
		absPath := filepath.Join(input.RootPath, filepath.FromSlash(relPath))

		read := ReadFileSafe(absPath, input.Options.MaxFileSizeKB)
		if read.Skipped {
			skipped++
			if read.Warning != "" {
				warnings = append(warnings, relPath+": "+read.Warning)
			}
			continue
		}

		renderedFiles = append(renderedFiles, RenderedFile{
			Path:    relPath,
			Content: read.Content,
		})

		totalLines += read.Lines
		totalSizeKB += read.SizeKB
	}

	stats := Statistics{
		FileCount:        len(renderedFiles),
		TotalLines:       totalLines,
		TotalSizeKB:      totalSizeKB,
		SkippedFiles:     skipped,
		ProcessingTimeMs: time.Since(start).Milliseconds(),
		GeneratedAt:      time.Now().UTC(),
	}

	return AssembleResult{
		Tree:     treeText,
		Files:    renderedFiles,
		Stats:    stats,
		Warnings: warnings,
	}
}

// --------------------
// Helpers
// --------------------

func buildTreeText(paths []string) string {
	if len(paths) == 0 {
		return ""
	}

	root := BuildTree(paths)

	var b strings.Builder
	var walk func(node *TreeNode, prefix string)

	walk = func(node *TreeNode, prefix string) {
		for i, child := range node.Children {
			connector := "├── "
			nextPrefix := prefix + "│   "
			if i == len(node.Children)-1 {
				connector = "└── "
				nextPrefix = prefix + "    "
			}

			b.WriteString(prefix + connector + child.Name + "\n")

			if child.IsDir {
				walk(child, nextPrefix)
			}
		}
	}

	walk(root, "")
	return b.String()
}
