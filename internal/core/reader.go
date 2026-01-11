package core

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)

// ReadResult represents the outcome of reading a file.
type ReadResult struct {
	Content string
	Lines   int
	SizeKB  int
	Skipped bool
	Warning string
}

// ReadFileSafe reads a file with safety checks.
// It enforces size limits, detects binary files, and normalizes line endings.
func ReadFileSafe(path string, maxSizeKB int) ReadResult {
	info, err := os.Stat(path)
	if err != nil {
		return ReadResult{
			Skipped: true,
			Warning: "cannot stat file: " + err.Error(),
		}
	}

	sizeKB := int(info.Size() / 1024)
	if maxSizeKB > 0 && sizeKB > maxSizeKB {
		return ReadResult{
			Skipped: true,
			SizeKB:  sizeKB,
			Warning: "file exceeds size limit",
		}
	}

	f, err := os.Open(path)
	if err != nil {
		return ReadResult{
			Skipped: true,
			Warning: "cannot open file: " + err.Error(),
		}
	}
	defer f.Close()

	// Read first 512 bytes for binary detection
	head := make([]byte, 512)
	n, _ := f.Read(head)
	if isBinary(head[:n]) {
		return ReadResult{
			Skipped: true,
			SizeKB:  sizeKB,
			Warning: "binary file detected",
		}
	}

	// Reset reader
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return ReadResult{
			Skipped: true,
			Warning: "failed to reset file reader",
		}
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return ReadResult{
			Skipped: true,
			Warning: "failed to read file",
		}
	}

	// Normalize line endings
	content := normalizeLineEndings(string(data))
	lines := strings.Count(content, "\n") + 1

	return ReadResult{
		Content: content,
		Lines:   lines,
		SizeKB: sizeKB,
	}
}

// --------------------
// Helpers
// --------------------

func isBinary(data []byte) bool {
	return bytes.IndexByte(data, 0x00) != -1
}

func normalizeLineEndings(s string) string {
	// Convert CRLF and CR to LF
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return s
}

// Guard against accidental misuse
var ErrBinaryFile = errors.New("binary file detected")
