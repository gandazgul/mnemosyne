package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/db"
	"github.com/gandazgul/mnemosyne/internal/embedding"
	"github.com/gandazgul/mnemosyne/internal/reranker"
	"github.com/gandazgul/mnemosyne/internal/setup"
)

// resolveCollectionName returns the collection name from the --name flag,
// the --global flag (which forces "global"), or the base name of the current working directory.
func resolveCollectionName(name string, global bool) (string, error) {
	if global {
		if name != "" {
			return "", fmt.Errorf("cannot use both --global and --name flags")
		}
		return "global", nil
	}

	if name == "global" {
		return "", fmt.Errorf("the name 'global' is reserved; use the --global or -g flag instead")
	}

	if name != "" {
		return name, nil
	}

	// Use the base name of the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getting current directory: %w", err)
	}

	base := filepath.Base(cwd)
	if base == "global" {
		return "", fmt.Errorf("current directory is named 'global', which is reserved; use the --global or -g flag instead")
	}

	return base, nil
}

// openDB loads config and opens the database connection.
// The caller is responsible for closing the returned *db.DB.
func openDB() (*db.DB, error) {
	cfg := config.Load()

	database, err := db.Open(cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	return database, nil
}

// openEmbedder initializes ONNX Runtime and creates an embedder from the config.
// The caller is responsible for calling Close() on the returned embedder.
//
// On first use, this triggers automatic download of ONNX Runtime and models
// if they are not already present.
//
// This is expensive (loads model into memory), so it should only be called by
// commands that need embeddings (add, search). Commands like list, delete, init
// should not call this.
func openEmbedder(ctx context.Context) (embedding.Embedder, *config.Config, error) {
	// Auto-download ONNX Runtime and models if not present.
	dataDir := config.DataDir()
	if err := setup.EnsureReady(ctx, dataDir, func(file string, written, total int64) {
		// Auto-download uses simple inline progress.
		// The setup command uses a full progress bar.
	}); err != nil {
		return nil, nil, fmt.Errorf("setup: %w", err)
	}

	// Re-load config after setup (paths may now resolve to downloaded files).
	cfg := config.Load()

	if err := embedding.InitONNXRuntime(cfg.OnnxRuntimeLib); err != nil {
		return nil, nil, fmt.Errorf("initializing ONNX Runtime: %w", err)
	}

	embedder, err := embedding.NewONNXEmbedder(cfg.Embedding)
	if err != nil {
		return nil, nil, fmt.Errorf("creating embedder: %w", err)
	}

	return embedder, cfg, nil
}

// openReranker initializes a cross-encoder reranker from the config.
// The caller is responsible for calling Close() on the returned reranker.
// Ensure openEmbedder (or InitONNXRuntime) is called before this.
func openReranker(cfg *config.Config) (reranker.Reranker, error) {
	if !cfg.Reranker.Enabled {
		return nil, nil
	}

	rr, err := reranker.NewONNXCrossEncoder(cfg.Reranker)
	if err != nil {
		return nil, fmt.Errorf("creating reranker: %w", err)
	}

	return rr, nil
}
