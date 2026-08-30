package setup

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gandazgul/mnemoteca/internal/config"
)

func TestReady_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	status := Check(tmpDir, testConfig(tmpDir))

	if status.Ready() {
		t.Error("Expected an empty directory to not be ready")
	}
}

func TestModelReady(t *testing.T) {
	tmpDir := t.TempDir()

	// Check false when missing files
	if modelReady(tmpDir, EmbeddingModel) {
		t.Error("Expected model not to be ready when missing files")
	}

	// Create all required files
	modelDir := filepath.Join(tmpDir, "models", EmbeddingModel.LocalDir)
	for file := range EmbeddingModel.Files {
		path := filepath.Join(modelDir, file)
		_ = os.MkdirAll(filepath.Dir(path), 0755)
		_ = os.WriteFile(path, []byte("dummy content"), 0644)
	}

	if !modelReady(tmpDir, EmbeddingModel) {
		t.Error("Expected model to be ready after files are created")
	}
}

func TestCheck(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := testConfig(tmpDir)

	status := Check(tmpDir, cfg)
	if status.OnnxRuntimeInstalled || status.EmbeddingModelReady || status.RerankerModelReady {
		t.Error("Expected all components to be missing in empty directory")
	}

	// Create onnxruntime
	libPath := filepath.Join(tmpDir, "lib", onnxRuntimeLibName())
	_ = os.MkdirAll(filepath.Dir(libPath), 0755)
	_ = os.WriteFile(libPath, []byte("fake dylib"), 0644)

	status = Check(tmpDir, cfg)
	if !status.OnnxRuntimeInstalled {
		t.Error("Expected ONNX runtime to be marked as installed")
	}
}

func TestCheck_CustomRerankerDisabled(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := testConfig(tmpDir)
	cfg.Reranker.Enabled = false

	status := Check(tmpDir, cfg)
	if status.RerankerEnabled {
		t.Error("Expected disabled reranker status to report RerankerEnabled=false")
	}
	if !status.RerankerModelReady {
		t.Error("Expected disabled reranker to be marked ready")
	}
}

func testConfig(dataDir string) *config.Config {
	cfg := config.DefaultConfig()
	cfg.Embedding.ModelPath = filepath.Join(dataDir, "models", EmbeddingModel.LocalDir)
	cfg.Reranker.ModelPath = filepath.Join(dataDir, "models", RerankerModel.LocalDir)
	cfg.OnnxRuntimeLib = filepath.Join(dataDir, "lib", onnxRuntimeLibName())
	return cfg
}
