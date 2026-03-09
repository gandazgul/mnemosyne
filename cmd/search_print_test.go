package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	
	"github.com/gandazgul/mnemosyne/internal/search"
)

func TestPrintSearchResults(t *testing.T) {
	// Setup capture
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	results := []search.Result{
		{
			DocumentID: 1,
			Content: "This is a test document",
			RRFScore: 0.95,
			Sources: []string{"vector", "fts"},
		},
	}
	
	// Test empty
	printSearchResults([]search.Result{}, "query", "col_a", "plain", false)
	
	// Test plain output
	printSearchResults(results, "query", "col_a", "plain", false)
	
	// Test plain with debug
	printSearchResults(results, "query", "col_a", "plain", true)
	
	// Test color output
	printSearchResults(results, "query", "col_a", "color", false)
	
	// Test color with debug
	printSearchResults(results, "query", "col_a", "color", true)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()
	
	if !strings.Contains(output, "No results for \"query\"") {
		t.Error("Missing empty results message")
	}
	
	if !strings.Contains(output, "This is a test document") {
		t.Error("Missing result content")
	}
	
	if !strings.Contains(output, "score:") {
		t.Error("Missing debug info")
	}
}
