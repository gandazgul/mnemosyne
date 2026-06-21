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
	os.Unsetenv(key) //nolint:errcheck
	val := getEnvOrDefault(key, fallback)
	if val != fallback {
		t.Errorf("expected fallback %q, got %q", fallback, val)
	}

	// 2. Set
	expected := "env_value"
	os.Setenv(key, expected) //nolint:errcheck
	defer os.Unsetenv(key)   //nolint:errcheck

	val = getEnvOrDefault(key, fallback)
	if val != expected {
		t.Errorf("expected %q, got %q", expected, val)
	}
}

func TestLoad(t *testing.T) {
	t.Setenv(envConfigPath, "")
	t.Setenv(envConfigPathAlt, "")
	t.Setenv("HOME", t.TempDir())

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}
	if cfg == nil {
		t.Fatal("Load() returned nil")
	}
}

func TestLoad_ConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	modelPath := filepath.Join(tmpDir, "models", "custom-embed")
	rerankerPath := filepath.Join(tmpDir, "models", "custom-reranker")
	dbPath := filepath.Join(tmpDir, "custom.db")

	yaml := `
db_path: "$MNEMOSYNE_TEST_DB_DIR/custom.db"
embedding:
  model_path: "$MNEMOSYNE_TEST_MODEL_DIR/models/custom-embed"
  dimensions: 768
  max_seq_length: 8192
  pooling: mean
  query_prefix: "query: "
  document_prefix: "passage: "
  onnx_input_names: ["input_ids", "attention_mask"]
  onnx_output_names: ["last_hidden_state"]
reranker:
  model_path: "$MNEMOSYNE_TEST_MODEL_DIR/models/custom-reranker"
  enabled: false
search:
  rerank_candidates: 25
`
	if err := os.WriteFile(configPath, []byte(yaml), 0o644); err != nil {
		t.Fatal(err)
	}

	t.Setenv(envConfigPath, configPath)
	t.Setenv(envConfigPathAlt, "")
	t.Setenv("MNEMOSYNE_TEST_DB_DIR", tmpDir)
	t.Setenv("MNEMOSYNE_TEST_MODEL_DIR", tmpDir)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.DBPath != dbPath {
		t.Errorf("expected DBPath %q, got %q", dbPath, cfg.DBPath)
	}
	if cfg.Embedding.ModelPath != modelPath {
		t.Errorf("expected embedding model path %q, got %q", modelPath, cfg.Embedding.ModelPath)
	}
	if cfg.Embedding.Dimensions != 768 {
		t.Errorf("expected dimensions 768, got %d", cfg.Embedding.Dimensions)
	}
	if cfg.Embedding.MaxSeqLength != 8192 {
		t.Errorf("expected max_seq_length 8192, got %d", cfg.Embedding.MaxSeqLength)
	}
	if cfg.Embedding.Pooling != PoolingMean {
		t.Errorf("expected pooling mean, got %q", cfg.Embedding.Pooling)
	}
	if cfg.Embedding.QueryPrefix != "query: " {
		t.Errorf("expected query prefix override, got %q", cfg.Embedding.QueryPrefix)
	}
	if cfg.Embedding.DocumentPrefix != "passage: " {
		t.Errorf("expected document prefix override, got %q", cfg.Embedding.DocumentPrefix)
	}
	if cfg.Reranker.ModelPath != rerankerPath {
		t.Errorf("expected reranker model path %q, got %q", rerankerPath, cfg.Reranker.ModelPath)
	}
	if cfg.Reranker.Enabled {
		t.Error("expected reranker to be disabled")
	}
	if cfg.Search.ReRankCandidates != 25 {
		t.Errorf("expected rerank candidates 25, got %d", cfg.Search.ReRankCandidates)
	}
	if cfg.Search.TopK != 10 {
		t.Errorf("expected omitted top_k to keep default 10, got %d", cfg.Search.TopK)
	}
}

func TestLoad_ExplicitMissingConfigReturnsError(t *testing.T) {
	t.Setenv(envConfigPath, filepath.Join(t.TempDir(), "missing.yaml"))
	t.Setenv(envConfigPathAlt, "")

	if _, err := Load(); err == nil {
		t.Fatal("expected explicit missing config to return an error")
	}
}

func TestLoad_DBEnvOverridesConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")
	configDBPath := filepath.Join(tmpDir, "config.db")
	envDBPath := filepath.Join(tmpDir, "env.db")

	if err := os.WriteFile(configPath, []byte("db_path: "+configDBPath+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	t.Setenv(envConfigPath, configPath)
	t.Setenv(envConfigPathAlt, "")
	t.Setenv("MNEMOSYNE_DB_PATH", envDBPath)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}
	if cfg.DBPath != envDBPath {
		t.Errorf("expected MNEMOSYNE_DB_PATH to override config file, got %q", cfg.DBPath)
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
	os.Setenv(envKey, expectedEnvPath) //nolint:errcheck
	defer os.Unsetenv(envKey)          //nolint:errcheck

	path := findONNXRuntimeLib(tmpDir)
	if path != expectedEnvPath {
		t.Errorf("expected path from env var %q, got %q", expectedEnvPath, path)
	}
	os.Unsetenv(envKey) //nolint:errcheck

	// Test finding in dataDir/lib
	libDir := filepath.Join(tmpDir, "lib")
	_ = os.MkdirAll(libDir, 0755)

	var libName string
	if runtime.GOOS == "darwin" {
		libName = "libonnxruntime.dylib"
	} else {
		libName = "libonnxruntime.so"
	}

	expectedLibPath := filepath.Join(libDir, libName)
	_ = os.WriteFile(expectedLibPath, []byte("dummy lib"), 0755)

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

func TestFindONNXRuntimeLib_Windows(t *testing.T) {
	tmpDir := t.TempDir()
	libDir := filepath.Join(tmpDir, "lib")
	if err := os.MkdirAll(libDir, 0o755); err != nil {
		t.Fatal(err)
	}

	expectedLibPath := filepath.Join(libDir, "onnxruntime.dll")
	if err := os.WriteFile(expectedLibPath, []byte("dummy dll"), 0o755); err != nil {
		t.Fatal(err)
	}

	path := findONNXRuntimeLibForOS(tmpDir, "windows")
	if path != expectedLibPath {
		t.Errorf("expected path %q, got %q", expectedLibPath, path)
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
	_ = os.MkdirAll(modelsDir, 0755)
	path = findModelsDir(tmpDir)
	if path != modelsDir {
		t.Errorf("expected path %q, got %q", modelsDir, path)
	}
}

func TestDefaultDataDir(t *testing.T) {
	// Test XDG_DATA_HOME
	expectedXDG := "/tmp/test-xdg"
	os.Setenv("XDG_DATA_HOME", expectedXDG) //nolint:errcheck
	defer os.Unsetenv("XDG_DATA_HOME")      //nolint:errcheck

	path := defaultDataDir()
	expectedPath := filepath.Join(expectedXDG, "mnemosyne")
	if path != expectedPath {
		t.Errorf("expected path %q, got %q", expectedPath, path)
	}

	// Test home dir fallback
	os.Unsetenv("XDG_DATA_HOME") //nolint:errcheck

	// Assuming HOME is set in test environment
	path = defaultDataDir()
	if path == "" {
		t.Error("expected non-empty path from defaultDataDir")
	}
}

func TestDefaultDataDir_Windows(t *testing.T) {
	localAppData := filepath.Join(t.TempDir(), "LocalAppData")
	appData := filepath.Join(t.TempDir(), "AppData")

	t.Setenv("XDG_DATA_HOME", filepath.Join(t.TempDir(), "xdg"))
	t.Setenv("LOCALAPPDATA", localAppData)
	t.Setenv("APPDATA", appData)

	path := defaultDataDirForOS("windows")
	expectedPath := filepath.Join(localAppData, "mnemosyne")
	if path != expectedPath {
		t.Errorf("expected LOCALAPPDATA path %q, got %q", expectedPath, path)
	}

	t.Setenv("LOCALAPPDATA", "")
	path = defaultDataDirForOS("windows")
	expectedPath = filepath.Join(appData, "mnemosyne")
	if path != expectedPath {
		t.Errorf("expected APPDATA fallback path %q, got %q", expectedPath, path)
	}
}
