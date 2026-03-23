package chunker

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

// NodeContext associates an AST node with its heading hierarchy path.
type NodeContext struct {
	Path []string
	Node ast.Node
}

// ExtractContext walks the AST (assuming root is a Document) and extracts top-level blocks
// with their associated heading hierarchy.
func ExtractContext(root ast.Node, source []byte) []NodeContext {
	var contexts []NodeContext
	var currentPath []string

	for child := root.FirstChild(); child != nil; child = child.NextSibling() {
		if heading, ok := child.(*ast.Heading); ok {
			level := heading.Level
			headingText := extractText(heading, source)

			// Truncate path to level-1
			if level-1 < len(currentPath) {
				currentPath = currentPath[:level-1]
			}
			// Pad if there are skipped levels
			for len(currentPath) < level-1 {
				currentPath = append(currentPath, "")
			}
			currentPath = append(currentPath, headingText)
		}

		pathCopy := make([]string, len(currentPath))
		copy(pathCopy, currentPath)

		contexts = append(contexts, NodeContext{
			Path: pathCopy,
			Node: child,
		})
	}

	return contexts
}

func extractText(n ast.Node, source []byte) string {
	var sb strings.Builder
	_ = ast.Walk(n, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			if node.Kind() == ast.KindText {
				textNode := node.(*ast.Text)
				sb.Write(textNode.Segment.Value(source))
			} else if node.Kind() == ast.KindString {
				strNode := node.(*ast.String)
				sb.Write(strNode.Value)
			}
			// Other text-bearing nodes could be added here if needed.
		}
		return ast.WalkContinue, nil
	})
	return strings.TrimSpace(sb.String())
}
