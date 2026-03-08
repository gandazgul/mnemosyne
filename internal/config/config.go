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

// EmbeddingConfig configures the embedding model.
type EmbeddingConfig struct {
	ModelPath      string `yaml:"model_path"`
	OnnxFile       string `yaml:"onnx_file"`
	Dimensions     int    `yaml:"dimensions"`
	MaxSeqLength   int    `yaml:"max_seq_length"`
	QueryPrefix    string `yaml:"query_prefix"`
	DocumentPrefix string `yaml:"document_prefix"`
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
}

// DefaultConfig returns the configuration with sensible defaults.
func DefaultConfig() *Config {
	dataDir := defaultDataDir()

	return &Config{
		DBPath:         filepath.Join(dataDir, "mnemosyne.db"),
		OnnxRuntimeLib: "", // empty means try default system paths
		Embedding: EmbeddingConfig{
			ModelPath:      filepath.Join(dataDir, "models", "embeddinggemma-300m"),
			OnnxFile:       "onnx/model_quantized.onnx",
			Dimensions:     768,
			MaxSeqLength:   2048,
			QueryPrefix:    "task: search result | query: ",
			DocumentPrefix: "title: none | text: ",
		},
		Reranker: RerankerConfig{
			ModelPath:    filepath.Join(dataDir, "models", "ms-marco-MiniLM-L-6-v2"),
			MaxSeqLength: 512,
			Enabled:      true,
		},
		Search: SearchConfig{
			RRFK:             60,
			TopK:             10,
			ReRankCandidates: 50,
		},
	}
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
