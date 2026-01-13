package formatter

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"repo2page/internal/core"
)

func TestRenderMarkdown(t *testing.T) {
	files := []core.RenderedFile{
		{Path: "main.go", Content: "package main"},
	}
	stats := core.Statistics{
		FileCount:   1,
		GeneratedAt: time.Now(),
	}

	output := RenderMarkdown("test-repo", "local", "├── main.go", files, stats, nil)

	if !strings.Contains(output, "# Repository: test-repo") {
		t.Error("Markdown output missing title")
	}
	if !strings.Contains(output, "```go") {
		t.Error("Markdown output missing syntax highlighting")
	}
	if !strings.Contains(output, "package main") {
		t.Error("Markdown output missing file content")
	}
}

func TestRenderJSON(t *testing.T) {
	files := []core.RenderedFile{
		{Path: "README.md", Content: "# Hello"},
	}
	stats := core.Statistics{
		FileCount:  1,
		TotalLines: 1,
	}

	output, err := RenderJSON("test-repo", "local", "tree", files, stats, []string{"warn1"})
	if err != nil {
		t.Fatalf("RenderJSON failed: %v", err)
	}

	var parsed JSONOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatalf("Failed to parse generated JSON: %v", err)
	}

	if parsed.Metadata.RepoName != "test-repo" {
		t.Errorf("JSON RepoName = %v, want test-repo", parsed.Metadata.RepoName)
	}
	if parsed.Files[0].Content != "# Hello" {
		t.Errorf("JSON Content mismatch")
	}
	if len(parsed.Warnings) != 1 {
		t.Errorf("JSON Warnings count = %d, want 1", len(parsed.Warnings))
	}
}
