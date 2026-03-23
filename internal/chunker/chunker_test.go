package chunker

import (
	"strings"
	"testing"
)

func TestChunkDocument(t *testing.T) {
	source := []byte(`Intro text.

# H1
Paragraph 1.

## H2
Paragraph 2.

- item 1
- item 2

# H1-B
Paragraph 3.`)

	chunks := ChunkDocument(source, 2000)

	if len(chunks) != 4 {
		t.Fatalf("Expected 4 chunks, got %d", len(chunks))
	}

	// Intro
	if len(chunks[0].Path) != 0 {
		t.Errorf("Chunk 0 path should be empty")
	}
	if !strings.Contains(chunks[0].Content, "Intro text.") {
		t.Errorf("Chunk 0 content mismatch")
	}

	// H1
	if strings.Join(chunks[1].Path, ">") != "H1" {
		t.Errorf("Chunk 1 path mismatch, got %v", chunks[1].Path)
	}
	if !strings.HasPrefix(chunks[1].Content, "[Path: H1]\n\n# H1") {
		t.Errorf("Chunk 1 formatting mismatch: %s", chunks[1].Content)
	}

	// H2
	if strings.Join(chunks[2].Path, ">") != "H1>H2" {
		t.Errorf("Chunk 2 path mismatch, got %v", chunks[2].Path)
	}
	if !strings.Contains(chunks[2].Content, "- item 1") {
		t.Errorf("Chunk 2 missing list item")
	}

	// H1-B
	if strings.Join(chunks[3].Path, ">") != "H1-B" {
		t.Errorf("Chunk 3 path mismatch, got %v", chunks[3].Path)
	}
}

func TestChunkDocument_SizeFallback(t *testing.T) {
	source := []byte(`# Large Section
This is a really large paragraph that needs to be split up. It has multiple sentences. One sentence. Two sentences. Three sentences.
`)
	// We set maxChars to a very small number to force splitting
	// len("[Path: Large Section]\n\n") is 23
	// So maxChars=50 gives limit=27
	chunks := ChunkDocument(source, 50)

	if len(chunks) < 2 {
		t.Fatalf("Expected chunk to be split into multiple sub-chunks, got %d", len(chunks))
	}

	for i, c := range chunks {
		if !strings.HasPrefix(c.Content, "[Path: Large Section]\n\n") {
			t.Errorf("Sub-chunk %d missing heading path prefix: %q", i, c.Content)
		}
	}
}
