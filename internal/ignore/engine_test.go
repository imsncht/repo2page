package ignore

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIgnoreEngine(t *testing.T) {
	// Create a temp directory for testing
	tmpDir, err := os.MkdirTemp("", "ignore_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a .gitignore file
	gitignoreContent := `
node_modules/
*.log
secret.txt
`
	if err := os.WriteFile(filepath.Join(tmpDir, ".gitignore"), []byte(gitignoreContent), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		path   string
		expect bool // true = ignored, false = included
	}{
		{"main.go", false},
		{"node_modules/package.json", true},
		{"app.log", true},
		{"secret.txt", true},
		{"src/utils.go", false},
		{".git/config", true}, // Default ignore pattern
	}

	engine, err := NewEngine(tmpDir, []string{"dist/"}) // Add extra user pattern
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := engine.ShouldIgnore(tt.path); got != tt.expect {
				t.Errorf("ShouldIgnore(%q) = %v, want %v", tt.path, got, tt.expect)
			}
		})
	}

	// Test user provided pattern "dist/"
	if !engine.ShouldIgnore("dist/bundle.js") {
		t.Error("ShouldIgnore('dist/bundle.js') = false, want true (user pattern)")
	}
}
