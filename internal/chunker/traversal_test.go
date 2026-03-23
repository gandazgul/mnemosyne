package chunker

import (
	"reflect"
	"testing"
)

func TestExtractContext(t *testing.T) {
	source := []byte(`
# Project Title
Some introductory paragraph.
## Setup
Setup instructions here.
### Environment
Env vars go here.
## Usage
How to use it.
`)
	root := ParseMarkdown(source)
	contexts := ExtractContext(root, source)

	// We expect multiple contexts
	if len(contexts) == 0 {
		t.Fatal("Expected non-empty contexts")
	}

	expectedPaths := [][]string{
		{"Project Title"},                         // The H1 node itself
		{"Project Title"},                         // The paragraph
		{"Project Title", "Setup"},                // The H2 node
		{"Project Title", "Setup"},                // The paragraph
		{"Project Title", "Setup", "Environment"}, // The H3 node
		{"Project Title", "Setup", "Environment"}, // The paragraph
		{"Project Title", "Usage"},                // The H2 node
		{"Project Title", "Usage"},                // The paragraph
	}

	if len(contexts) != len(expectedPaths) {
		t.Fatalf("Expected %d contexts, got %d", len(expectedPaths), len(contexts))
	}

	for i, expected := range expectedPaths {
		if !reflect.DeepEqual(contexts[i].Path, expected) {
			t.Errorf("At index %d, expected path %v, got %v", i, expected, contexts[i].Path)
		}
	}
}

func TestExtractContext_EmptyHeading(t *testing.T) {
	source := []byte(`
# 
No title
`)
	root := ParseMarkdown(source)
	contexts := ExtractContext(root, source)

	if len(contexts) != 2 {
		t.Fatalf("Expected 2 contexts, got %d", len(contexts))
	}

	if contexts[1].Path[0] != "" {
		t.Errorf("Expected empty path for empty heading, got '%s'", contexts[1].Path[0])
	}
}
