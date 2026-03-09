package setup

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReady_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	status := Check(tmpDir)
	
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
		os.MkdirAll(filepath.Dir(path), 0755)
		os.WriteFile(path, []byte("dummy content"), 0644)
	}
	
	if !modelReady(tmpDir, EmbeddingModel) {
		t.Error("Expected model to be ready after files are created")
	}
}

func TestCheck(t *testing.T) {
	tmpDir := t.TempDir()
	
	status := Check(tmpDir)
	if status.OnnxRuntimeInstalled || status.EmbeddingModelReady || status.RerankerModelReady {
		t.Error("Expected all components to be missing in empty directory")
	}
	
	// Create onnxruntime
	libPath := filepath.Join(tmpDir, "lib", onnxRuntimeLibName())
	os.MkdirAll(filepath.Dir(libPath), 0755)
	os.WriteFile(libPath, []byte("fake dylib"), 0644)
	
	status = Check(tmpDir)
	if !status.OnnxRuntimeInstalled {
		t.Error("Expected ONNX runtime to be marked as installed")
	}
}
