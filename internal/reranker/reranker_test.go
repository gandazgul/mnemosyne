package reranker

import (
	"os"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/embedding"
)

func TestMain(m *testing.M) {
	// Initialize ONNX Runtime for tests.
	// Allow overriding library path via environment variable.
	libPath := os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH")
	if libPath == "" {
		libPath = "../../lib/libonnxruntime.dylib"
		if _, err := os.Stat(libPath); err != nil {
			libPath = "../../lib/libonnxruntime.so" // fallback for Linux
		}
	}

	err := embedding.InitONNXRuntime(libPath)
	if err != nil {
		panic(err)
	}
	defer func() { _ = embedding.DestroyONNXRuntime() }()

	os.Exit(m.Run())
}

func TestONNXCrossEncoder(t *testing.T) {
	cfg := config.DefaultConfig()
	// Update path relative to the tests
	cfg.Reranker.ModelPath = "../../models/ms-marco-MiniLM-L-6-v2"

	rr, err := NewONNXCrossEncoder(cfg.Reranker)
	if err != nil {
		t.Fatalf("failed to create reranker: %v", err)
	}
	defer rr.Close() //nolint:errcheck

	query := "What is Go?"
	docs := []string{
		"Python is a high-level programming language.",
		"Go is a statically typed, compiled programming language designed at Google.",
		"JavaScript is the programming language of the Web.",
	}

	scores, err := rr.Score(query, docs)
	if err != nil {
		t.Fatalf("Score() failed: %v", err)
	}

	if len(scores) != len(docs) {
		t.Fatalf("expected %d scores, got %d", len(docs), len(scores))
	}

	// Go should have the highest score
	if scores[1] <= scores[0] || scores[1] <= scores[2] {
		t.Errorf("expected doc 1 (Go) to have highest score. Scores: %v", scores)
	}
}
