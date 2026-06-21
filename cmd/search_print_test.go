package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
	"time"

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
			Content:    "This is a test document",
			RRFScore:   0.95,
			Sources:    []string{"vector", "fts"},
		},
	}

	// Test empty
	if err := printSearchResults([]search.Result{}, "query", "col_a", "plain", false); err != nil {
		t.Fatalf("printSearchResults empty: %v", err)
	}

	// Test plain output
	if err := printSearchResults(results, "query", "col_a", "plain", false); err != nil {
		t.Fatalf("printSearchResults plain: %v", err)
	}

	// Test plain with debug
	if err := printSearchResults(results, "query", "col_a", "plain", true); err != nil {
		t.Fatalf("printSearchResults plain debug: %v", err)
	}

	// Test color output
	if err := printSearchResults(results, "query", "col_a", "color", false); err != nil {
		t.Fatalf("printSearchResults color: %v", err)
	}

	// Test color with debug
	if err := printSearchResults(results, "query", "col_a", "color", true); err != nil {
		t.Fatalf("printSearchResults color debug: %v", err)
	}

	w.Close() //nolint:errcheck
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
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

func TestPrintSearchResultsJSON(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	metadata := `{"benchmark_doc_id":"doc-1","tags":["eval"]}`
	results := []search.Result{
		{
			DocumentID:    1,
			CollectionID:  2,
			Content:       "This is a test document",
			Metadata:      &metadata,
			CreatedAt:     time.Date(2026, 6, 19, 12, 30, 0, 0, time.UTC),
			RRFScore:      0.032,
			FTSRank:       1.2,
			VecDistance:   0.15,
			RerankerScore: 0.98,
			IsReranked:    true,
			Sources:       []string{"fts", "vector"},
		},
	}

	if err := printSearchResults(results, "query", "col_a", "json", false); err != nil {
		t.Fatalf("printSearchResults json: %v", err)
	}

	w.Close() //nolint:errcheck
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	var output struct {
		Query      string `json:"query"`
		Collection string `json:"collection"`
		Count      int    `json:"count"`
		Results    []struct {
			Rank          int            `json:"rank"`
			DocumentID    int64          `json:"document_id"`
			CollectionID  int64          `json:"collection_id"`
			Content       string         `json:"content"`
			Metadata      map[string]any `json:"metadata"`
			CreatedAt     string         `json:"created_at"`
			RRFScore      float64        `json:"rrf_score"`
			FTSRank       float64        `json:"fts_rank"`
			VecDistance   float64        `json:"vec_distance"`
			RerankerScore float32        `json:"reranker_score"`
			IsReranked    bool           `json:"is_reranked"`
			Sources       []string       `json:"sources"`
		} `json:"results"`
	}
	if err := json.Unmarshal(buf.Bytes(), &output); err != nil {
		t.Fatalf("invalid json output: %v\n%s", err, buf.String())
	}

	if output.Query != "query" || output.Collection != "col_a" || output.Count != 1 {
		t.Fatalf("unexpected envelope: %+v", output)
	}
	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}

	got := output.Results[0]
	if got.Rank != 1 || got.DocumentID != 1 || got.CollectionID != 2 {
		t.Fatalf("unexpected result identity: %+v", got)
	}
	if got.Metadata["benchmark_doc_id"] != "doc-1" {
		t.Fatalf("expected parsed metadata, got %+v", got.Metadata)
	}
	if got.CreatedAt != "2026-06-19T12:30:00Z" {
		t.Fatalf("unexpected created_at: %s", got.CreatedAt)
	}
	if !got.IsReranked || len(got.Sources) != 2 {
		t.Fatalf("expected reranked two-source result, got %+v", got)
	}
}
