package integration

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var binaryPath string

func TestMain(m *testing.M) {
	// 1. Build the binary
	if runtime.GOOS == "windows" {
		binaryPath = filepath.Join(os.TempDir(), "repo2page.exe")
	} else {
		binaryPath = filepath.Join(os.TempDir(), "repo2page")
	}

	// Clean up previous build if exists
	os.Remove(binaryPath)

	// Get project root (assuming we run from project root or checks relative)
	// We assume "go test ./tests/integration/..." is run from project root?
	// Or we locate cmd/repo2page relative to this test file.
	// This test file is in tests/integration/
	// Project root is ../../
	
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get working directory: %v\n", err)
		os.Exit(1)
	}
	
	projectRoot := filepath.Join(wd, "..", "..")
	// If running from project root using go test ./...
	if filepath.Base(wd) == "repo2page" {
		projectRoot = "."
	}

	cmdPath := filepath.Join(projectRoot, "cmd", "repo2page")

	fmt.Printf("Building binary from %s to %s...\n", cmdPath, binaryPath)
	buildCmd := exec.Command("go", "build", "-o", binaryPath, cmdPath)
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to build binary: %v\nOutput: %s\n", err, out)
		os.Exit(1)
	}

	// 2. Run tests
	code := m.Run()

	// 3. Cleanup
	os.Remove(binaryPath)
	os.Exit(code)
}

func TestCLI_Version(t *testing.T) {
	out, err := exec.Command(binaryPath, "--version").CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}
	output := string(out)
	if output == "" {
		t.Error("Expected version output, got empty string")
	}
}

func TestCLI_SmallRepo(t *testing.T) {
	// Locate fixture
	wd, _ := os.Getwd()
	// Adjust based on where test is run. Safe assumption: absolute path relative to known location.
	// If running from repo root: tests/fixtures/small-repo
	projectRoot := filepath.Join(wd, "..", "..")
	if filepath.Base(wd) == "repo2page" {
		projectRoot = "."
	}
	
	fixturePath := filepath.Join(projectRoot, "tests", "fixtures", "small-repo")
	outputPath := filepath.Join(os.TempDir(), "small-repo-output.md")
	defer os.Remove(outputPath)

	cmd := exec.Command(binaryPath, "--output", outputPath, fixturePath)
	out, err := cmd.CombinedOutput()
	
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, string(out))
	}

	// Check output file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Output file was not created: %s", outputPath)
	}
}

func TestCLI_InvalidSource(t *testing.T) {
	cmd := exec.Command(binaryPath, "./non-existent-path")
	err := cmd.Run()
	
	if err == nil {
		t.Error("Expected error for non-existent path, got success")
	}
	// Check exit code? 
	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() == 0 {
			t.Error("Expected non-zero exit code")
		}
	}
}

func TestCLI_JSONOutput(t *testing.T) {
	wd, _ := os.Getwd()
	projectRoot := filepath.Join(wd, "..", "..")
	if filepath.Base(wd) == "repo2page" {
		projectRoot = "."
	}
	fixturePath := filepath.Join(projectRoot, "tests", "fixtures", "small-repo")
	outputPath := filepath.Join(os.TempDir(), "output.json")
	defer os.Remove(outputPath)

	cmd := exec.Command(binaryPath, "--format", "json", "--output", outputPath, fixturePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, string(out))
	}

	// Verify JSON content
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(content, &data); err != nil {
		t.Errorf("Output is not valid JSON: %v", err)
	}

	if data["files"] == nil {
		t.Error("JSON output missing 'files' field")
	}
}
