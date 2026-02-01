package repo2page

import "github.com/imsncht/repo2page/internal/core"

// Convert converts a repository (local path or GitHub URL)
// into a single, portable document.
//
// This is a stable public API intended for library use.
// It mirrors the CLI behavior without performing any I/O.
func Convert(input core.RepoInput, options core.ConvertOptions) (*core.ConvertResult, error) {
	return core.Convert(input, options)
}
