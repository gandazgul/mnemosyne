package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/db"
)

// resolveCollectionName returns the collection name from the --name flag,
// or the base name of the current working directory if the flag is empty.
func resolveCollectionName(name string) (string, error) {
	if name != "" {
		return name, nil
	}

	// Use the base name of the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getting current directory: %w", err)
	}

	return filepath.Base(cwd), nil
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
