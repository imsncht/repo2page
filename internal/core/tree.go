package core

import (
	"path"
	"sort"
	"strings"
)

// TreeNode represents a node in the repository tree.
type TreeNode struct {
	Name     string
	Path     string
	IsDir    bool
	Children []*TreeNode
}

// BuildTree constructs an in-memory tree from a list of file paths.
// Paths must be forward-slash normalized.
func BuildTree(paths []string) *TreeNode {
	root := &TreeNode{
		Name:  ".",
		Path:  "",
		IsDir: true,
	}

	for _, p := range paths {
		parts := strings.Split(p, "/")
		current := root

		for i, part := range parts {
			isDir := i < len(parts)-1
			var next *TreeNode

			for _, child := range current.Children {
				if child.Name == part && child.IsDir == isDir {
					next = child
					break
				}
			}

			if next == nil {
				nodePath := part
				if current.Path != "" {
					nodePath = path.Join(current.Path, part)
				}

				next = &TreeNode{
					Name:  part,
					Path:  nodePath,
					IsDir: isDir,
				}
				current.Children = append(current.Children, next)
			}

			current = next
		}
	}

	sortTree(root)
	return root
}

// sortTree applies deterministic ordering rules recursively.
func sortTree(node *TreeNode) {
	if len(node.Children) == 0 {
		return
	}

	sort.SliceStable(node.Children, func(i, j int) bool {
		a := node.Children[i]
		b := node.Children[j]

		// Root-level README first
		if node.Path == "" {
			if isReadme(a.Name) && !isReadme(b.Name) {
				return true
			}
			if !isReadme(a.Name) && isReadme(b.Name) {
				return false
			}
		}

		// Directories before files
		if a.IsDir != b.IsDir {
			return a.IsDir
		}

		// Case-insensitive alphabetical order
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)
	})

	for _, child := range node.Children {
		if child.IsDir {
			sortTree(child)
		}
	}
}

// FlattenTree returns a depth-first ordered list of file paths.
func FlattenTree(node *TreeNode) []string {
	var result []string

	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		for _, child := range n.Children {
			if child.IsDir {
				walk(child)
			} else {
				result = append(result, child.Path)
			}
		}
	}

	walk(node)
	return result
}

func isReadme(name string) bool {
	lower := strings.ToLower(name)
	return strings.HasPrefix(lower, "readme")
}
