package version

import "fmt"

// These variables are set at build time using -ldflags.
// Do NOT modify them at runtime.
var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

// Full returns a human-readable version string.
func Full() string {
	return fmt.Sprintf(
		"repo2page %s (commit %s, built %s)",
		Version,
		Commit,
		Date,
	)
}
