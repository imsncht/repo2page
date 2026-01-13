package core

import (
	"reflect"
	"testing"
)

func TestTreeBuilder(t *testing.T) {
	paths := []string{
		"src/utils.go",
		"README.md",
		"src/main.go",
		"go.mod",
		"docs/index.md",
	}

	// Build the tree
	root := BuildTree(paths)

	// Flatten it back to verify deterministic order
	// Expected order:
	// 1. README.md (Special priority)
	// 2. docs/ (Directory) -> docs/index.md
	// 3. src/ (Directory) -> src/main.go, src/utils.go (Alphabetical)
	// 4. go.mod (File)
	
	// Wait, check SortTree logic in tree.go:
	// "README* first, then directories, then files (alphabetical)"

	got := FlattenTree(root)

	want := []string{
		"README.md",
		"docs/index.md",
		"src/main.go",
		"src/utils.go",
		"go.mod",
	}

	// Note: FlattenTree returns full paths.
	// Let's check if my manual sort expectation matches the implementation.
	// docs/ vs go.mod? "directories before files" -> docs/ comes before go.mod?
	// Let's run the test and see. If it fails, I'll align with the actual implementation logic.

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlattenTree() order mismatch.\nGot:  %v\nWant: %v", got, want)
	}
}

func TestTreeStructure(t *testing.T) {
	paths := []string{"a/b/c.txt"}
	root := BuildTree(paths)
	
	if len(root.Children) != 1 {
		t.Fatalf("Root should have 1 child (a), got %d", len(root.Children))
	}
	
	nodeA := root.Children[0]
	if nodeA.Name != "a" || !nodeA.IsDir {
		t.Errorf("Expected dir 'a', got %v", nodeA)
	}
	
	if len(nodeA.Children) != 1 {
		t.Fatalf("Node 'a' should have 1 child (b), got %d", len(nodeA.Children))
	}
	
	nodeB := nodeA.Children[0]
	if nodeB.Name != "b" || !nodeB.IsDir {
		t.Errorf("Expected dir 'b', got %v", nodeB)
	}
	
	if len(nodeB.Children) != 1 {
		t.Fatalf("Node 'b' should have 1 child (c.txt), got %d", len(nodeB.Children))
	}
	
	nodeC := nodeB.Children[0]
	if nodeC.Name != "c.txt" || nodeC.IsDir {
		t.Errorf("Expected file 'c.txt', got %v", nodeC)
	}
}
