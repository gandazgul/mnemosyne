// Package config handles loading and providing application configuration.
//
// Configuration is resolved in this order (highest priority first):
//  1. CLI flags
//  2. Environment variables
//  3. Config file (~/.config/mnemosyne/config.yaml)
//  4. Built-in defaults
package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// Config holds all application configuration.
type Config struct {
	// DBPath is the path to the SQLite database file.
	DBPath string `yaml:"db_path"`

	// OnnxRuntimeLib is the path to the ONNX Runtime shared library.
	// If empty, the library will try default system paths.
	OnnxRuntimeLib string `yaml:"onnx_runtime_lib"`

	// Embedding holds configuration for the embedding model.
	Embedding EmbeddingConfig `yaml:"embedding"`

	// Reranker holds configuration for the reranker model.
	Reranker RerankerConfig `yaml:"reranker"`

	// Search holds configuration for search behavior.
	Search SearchConfig `yaml:"search"`
}

// PoolingMethod specifies how token-level hidden states are reduced to a
// single sentence embedding.
type PoolingMethod string

const (
	// PoolingMean averages all non-padding token vectors (weighted by attention mask).
	PoolingMean PoolingMethod = "mean"
	// PoolingCLS takes the first token's ([CLS]) output vector.
	PoolingCLS PoolingMethod = "cls"
	// PoolingNone means the model already provides a pooled sentence embedding
	// output and no additional pooling is needed.
	PoolingNone PoolingMethod = "none"
)

// EmbeddingConfig configures the embedding model.
type EmbeddingConfig struct {
	ModelPath      string        `yaml:"model_path"`
	OnnxFile       string        `yaml:"onnx_file"`
	Dimensions     int           `yaml:"dimensions"`
	MaxSeqLength   int           `yaml:"max_seq_length"`
	QueryPrefix    string        `yaml:"query_prefix"`
	DocumentPrefix string        `yaml:"document_prefix"`
	Pooling        PoolingMethod `yaml:"pooling"`

	// OnnxInputNames overrides the default ONNX input node names.
	// Default: ["input_ids", "attention_mask"] (+ "token_type_ids" for CLS pooling).
	OnnxInputNames []string `yaml:"onnx_input_names"`

	// OnnxOutputNames overrides the default ONNX output node names.
	// Default: ["last_hidden_state"] for mean/CLS pooling, ["sentence_embedding"] for none.
	OnnxOutputNames []string `yaml:"onnx_output_names"`
}

// RerankerConfig configures the cross-encoder reranker model.
type RerankerConfig struct {
	ModelPath    string `yaml:"model_path"`
	MaxSeqLength int    `yaml:"max_seq_length"`
	Enabled      bool   `yaml:"enabled"`
}

// SearchConfig configures search behavior.
type SearchConfig struct {
	RRFK             int `yaml:"rrf_k"`
	TopK             int `yaml:"top_k"`
	ReRankCandidates int `yaml:"rerank_candidates"`

	// RerankerThreshold is the minimum reranker score (logit) for a result
	// to be included. With a sigmoid function applied, values are between 0.0 and 1.0.
	// We previously used logits with a threshold of -6.0. The equivalent sigmoid threshold
	// is around 0.001 to 0.002.
	// Default: 0.001
	RerankerThreshold float64 `yaml:"reranker_threshold"`

	// RRFThreshold is the minimum RRF fusion score for a result to be included
	// when reranking is disabled. With k=60, a single-source rank-1 result
	// scores ~0.016. Default: 0.01
	RRFThreshold float64 `yaml:"rrf_threshold"`
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// DefaultConfig returns the configuration with sensible defaults.
func DefaultConfig() *Config {
	dataDir := defaultDataDir()
	modelsDir := findModelsDir(dataDir)

	return &Config{
		DBPath:         getEnvOrDefault("MNEMOSYNE_DB_PATH", filepath.Join(dataDir, "mnemosyne.db")),
		OnnxRuntimeLib: findONNXRuntimeLib(dataDir),
		Embedding: EmbeddingConfig{
			ModelPath:       filepath.Join(modelsDir, "snowflake-arctic-embed-m-v1.5"),
			OnnxFile:        "onnx/model.onnx",
			Dimensions:      256,
			MaxSeqLength:    512,
			QueryPrefix:     "Represent this sentence for searching relevant passages: ",
			DocumentPrefix:  "",
			Pooling:         PoolingNone,
			OnnxInputNames:  []string{"input_ids", "attention_mask"},
			OnnxOutputNames: []string{"sentence_embedding"},
		},
		Reranker: RerankerConfig{
			ModelPath:    filepath.Join(modelsDir, "ms-marco-MiniLM-L-6-v2"),
			MaxSeqLength: 512,
			Enabled:      true,
		},
		Search: SearchConfig{
			RRFK:              60,
			TopK:              10,
			ReRankCandidates:  50,
			RerankerThreshold: 0.01, // Corresponds roughly to a logit of -6.9
			RRFThreshold:      0.01,
		},
	}
}

// findONNXRuntimeLib searches well-known locations for the ONNX Runtime shared
// library and returns the first path found. Returns empty string if not found,
// which lets onnxruntime_go fall back to its default system-path search.
//
// Search order:
//  1. ONNXRUNTIME_SHARED_LIBRARY_PATH environment variable
//  2. Data directory lib/ (e.g. ~/.local/share/mnemosyne/lib/)
//  3. Next to the running executable
//  4. ../lib/ relative to the running executable (development layout)
func findONNXRuntimeLib(dataDir string) string {
	// Honour explicit env var (matches the convention used by tests).
	if p := os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH"); p != "" {
		return p
	}

	// Platform-specific library names.
	var names []string
	switch runtime.GOOS {
	case "darwin":
		names = []string{"libonnxruntime.dylib"}
	default: // linux, etc.
		names = []string{"libonnxruntime.so"}
	}

	// Candidate directories to search.
	dirs := []string{
		filepath.Join(dataDir, "lib"),
	}

	// Add executable-relative paths (useful during development).
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		dirs = append(dirs, exeDir, filepath.Join(exeDir, "lib"))
	}

	for _, dir := range dirs {
		for _, name := range names {
			p := filepath.Join(dir, name)
			if _, err := os.Stat(p); err == nil {
				return p
			}
		}
	}

	return ""
}

// findModelsDir returns the path to the models directory.
// It checks the data directory first, then falls back to a directory relative
// to the running executable (for development layouts where models/ is in the
// project root alongside the binary).
func findModelsDir(dataDir string) string {
	dataDirModels := filepath.Join(dataDir, "models")
	if isDir(dataDirModels) {
		return dataDirModels
	}

	// Try executable-relative paths (development layout).
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		if p := filepath.Join(exeDir, "models"); isDir(p) {
			return p
		}
	}

	// Fall back to data dir even if it doesn't exist yet; the user will get
	// a clear "model not found" error pointing them to the expected location.
	return dataDirModels
}

// isDir returns true if path exists and is a directory.
func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// defaultDataDir returns the default data directory for mnemosyne.
// On macOS/Linux: ~/.local/share/mnemosyne
// Falls back to ~/.mnemosyne if XDG_DATA_HOME is not set and home can't be resolved.
func defaultDataDir() string {
	// Check XDG_DATA_HOME first (standard on Linux).
	if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
		return filepath.Join(xdg, "mnemosyne")
	}

	// Fall back to ~/.local/share/mnemosyne.
	home, err := os.UserHomeDir()
	if err != nil {
		// Last resort: use current directory.
		return ".mnemosyne"
	}

	return filepath.Join(home, ".local", "share", "mnemosyne")
}

// Load returns the application configuration.
// For now, it returns defaults. YAML file loading will be added later.
func Load() *Config {
	return DefaultConfig()
}

// DataDir returns the default data directory path.
// This is used by the setup package to know where to download files.
func DataDir() string {
	return defaultDataDir()
}
