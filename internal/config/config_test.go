package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg == nil {
		t.Fatal("DefaultConfig() returned nil")
	}

	if cfg.Embedding.Dimensions != 256 {
		t.Errorf("expected dimensions 256, got %d", cfg.Embedding.Dimensions)
	}
	if cfg.Search.TopK != 10 {
		t.Errorf("expected TopK 10, got %d", cfg.Search.TopK)
	}
	if cfg.DBPath == "" {
		t.Error("expected DBPath to be set")
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	key := "MNEMOSYNE_TEST_ENV_KEY"
	fallback := "fallback_value"

	// 1. Not set
	os.Unsetenv(key)
	val := getEnvOrDefault(key, fallback)
	if val != fallback {
		t.Errorf("expected fallback %q, got %q", fallback, val)
	}

	// 2. Set
	expected := "env_value"
	os.Setenv(key, expected)
	defer os.Unsetenv(key)

	val = getEnvOrDefault(key, fallback)
	if val != expected {
		t.Errorf("expected %q, got %q", expected, val)
	}
}

func TestLoad(t *testing.T) {
	cfg := Load()
	if cfg == nil {
		t.Fatal("Load() returned nil")
	}
}

func TestDataDir(t *testing.T) {
	dir := DataDir()
	if dir == "" {
		t.Error("DataDir() returned empty string")
	}
}

func TestIsDir(t *testing.T) {
	// Create a temp dir
	tmpDir := t.TempDir()

	if !isDir(tmpDir) {
		t.Errorf("isDir(%q) returned false, expected true", tmpDir)
	}

	// Test non-existent dir
	if isDir(filepath.Join(tmpDir, "does-not-exist")) {
		t.Error("isDir() returned true for non-existent directory")
	}

	// Test file (not dir)
	tmpFile := filepath.Join(tmpDir, "file.txt")
	if err := os.WriteFile(tmpFile, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}
	if isDir(tmpFile) {
		t.Errorf("isDir(%q) returned true for a file", tmpFile)
	}
}

func TestFindONNXRuntimeLib(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with explicit env var
	envKey := "ONNXRUNTIME_SHARED_LIBRARY_PATH"
	expectedEnvPath := filepath.Join(tmpDir, "env-libonnxruntime.so")
	os.Setenv(envKey, expectedEnvPath)
	defer os.Unsetenv(envKey)

	path := findONNXRuntimeLib(tmpDir)
	if path != expectedEnvPath {
		t.Errorf("expected path from env var %q, got %q", expectedEnvPath, path)
	}
	os.Unsetenv(envKey)

	// Test finding in dataDir/lib
	libDir := filepath.Join(tmpDir, "lib")
	os.MkdirAll(libDir, 0755)

	var libName string
	if runtime.GOOS == "darwin" {
		libName = "libonnxruntime.dylib"
	} else {
		libName = "libonnxruntime.so"
	}

	expectedLibPath := filepath.Join(libDir, libName)
	os.WriteFile(expectedLibPath, []byte("dummy lib"), 0755)

	path = findONNXRuntimeLib(tmpDir)
	if path != expectedLibPath {
		t.Errorf("expected path %q, got %q", expectedLibPath, path)
	}

	// Test fallback (empty string)
	emptyDir := t.TempDir()
	path = findONNXRuntimeLib(emptyDir)
	if path != "" {
		t.Errorf("expected empty string when not found, got %q", path)
	}
}

func TestFindModelsDir(t *testing.T) {
	tmpDir := t.TempDir()
	modelsDir := filepath.Join(tmpDir, "models")

	// 1. Not exists yet -> should return dataDir/models anyway as fallback
	path := findModelsDir(tmpDir)
	if path != modelsDir {
		t.Errorf("expected fallback path %q, got %q", modelsDir, path)
	}

	// 2. Create the dir -> should return it
	os.MkdirAll(modelsDir, 0755)
	path = findModelsDir(tmpDir)
	if path != modelsDir {
		t.Errorf("expected path %q, got %q", modelsDir, path)
	}
}

func TestDefaultDataDir(t *testing.T) {
	// Test XDG_DATA_HOME
	expectedXDG := "/tmp/test-xdg"
	os.Setenv("XDG_DATA_HOME", expectedXDG)
	defer os.Unsetenv("XDG_DATA_HOME")

	path := defaultDataDir()
	expectedPath := filepath.Join(expectedXDG, "mnemosyne")
	if path != expectedPath {
		t.Errorf("expected path %q, got %q", expectedPath, path)
	}

	// Test home dir fallback
	os.Unsetenv("XDG_DATA_HOME")

	// Assuming HOME is set in test environment
	path = defaultDataDir()
	if path == "" {
		t.Error("expected non-empty path from defaultDataDir")
	}
}
