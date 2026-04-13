package reranker

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/embedding"
)

// testOnnxRuntimeLib returns the path to the ONNX Runtime shared library.
// Tests are skipped if the library is not available.
func testOnnxRuntimeLib(t *testing.T) string {
	t.Helper()

	if envPath := os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH"); envPath != "" {
		return envPath
	}

	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	libDir := filepath.Join(projectRoot, "lib")

	candidates := []string{
		"libonnxruntime.dylib",
		"libonnxruntime.so",
		"libonnxruntime.1.23.1.dylib",
	}

	for _, name := range candidates {
		path := filepath.Join(libDir, name)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	t.Skipf("Skipping: ONNX Runtime library not found in %s (run 'task download-onnxruntime')", libDir)
	return ""
}

// testModelsDir returns the path to the reranker model directory.
// Tests are skipped if the model is not available.
func testModelsDir(t *testing.T) string {
	t.Helper()

	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	modelDir := filepath.Join(projectRoot, "models", "ms-marco-MiniLM-L-6-v2")

	onnxPath := filepath.Join(modelDir, "onnx", "model.onnx")
	if _, err := os.Stat(onnxPath); err != nil {
		t.Skipf("Skipping: reranker model not found at %s (run 'task download-models')", modelDir)
	}

	return modelDir
}

func TestONNXCrossEncoder(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	modelDir := testModelsDir(t)

	if err := embedding.InitONNXRuntime(libPath); err != nil {
		t.Fatalf("initializing ONNX Runtime: %v", err)
	}
	defer func() { _ = embedding.DestroyONNXRuntime() }()

	cfg := config.DefaultConfig()
	cfg.Reranker.ModelPath = modelDir

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
