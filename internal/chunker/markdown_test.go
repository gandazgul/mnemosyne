package chunker

import (
	"testing"
)

func TestParseMarkdown(t *testing.T) {
	source := []byte("# Hello\n\nThis is a test.")
	node := ParseMarkdown(source)

	if node == nil {
		t.Fatal("Expected an AST node, got nil")
	}

	if node.ChildCount() != 2 {
		t.Fatalf("Expected 2 children nodes (heading, paragraph), got %d", node.ChildCount())
	}
}
