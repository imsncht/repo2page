package formatter

import (
	"encoding/json"
	
	"repo2page/internal/core"
)

// JSONOutput represents the schema for the JSON output format.
type JSONOutput struct {
	Metadata Metadata             `json:"metadata"`
	Tree     string               `json:"tree"`
	Files    []core.RenderedFile  `json:"files"`
	Warnings []string             `json:"warnings,omitempty"`
}

type Metadata struct {
	RepoName string          `json:"repo_name"`
	Source   string          `json:"source"`
	Stats    core.Statistics `json:"stats"`
}

// RenderJSON generates a JSON representation of the repository.
func RenderJSON(repoName, source, tree string, files []core.RenderedFile, stats core.Statistics, warnings []string) (string, error) {
	out := JSONOutput{
		Metadata: Metadata{
			RepoName: repoName,
			Source:   source,
			Stats:    stats,
		},
		Tree:     tree,
		Files:    files,
		Warnings: warnings,
	}

	// efficient serialization
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
