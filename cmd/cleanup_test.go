package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCleanupCmd_DefaultKeepsDB(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)
	dataDir := filepath.Join(tmpDir, "mnemosyne")

	// Create fake data directory structure.
	for _, dir := range []string{"lib", "models/embed", "models/rerank"} {
		if err := os.MkdirAll(filepath.Join(dataDir, dir), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
	}
	dbPath := filepath.Join(dataDir, "mnemosyne.db")
	if err := os.WriteFile(dbPath, []byte("fake-db"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// Default cleanup with --yes should remove lib/ and models/ but keep db.
	rootCmd.SetArgs([]string{"cleanup", "--yes"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "Removed") {
		t.Errorf("expected 'Removed' message, got: %s", output)
	}

	// lib/ and models/ should be gone.
	if _, err := os.Stat(filepath.Join(dataDir, "lib")); !os.IsNotExist(err) {
		t.Error("expected lib/ to be removed")
	}
	if _, err := os.Stat(filepath.Join(dataDir, "models")); !os.IsNotExist(err) {
		t.Error("expected models/ to be removed")
	}

	// Database should still exist.
	if _, err := os.Stat(dbPath); err != nil {
		t.Errorf("expected database to still exist, got error: %v", err)
	}
}

func TestCleanupCmd_WithDBFlag(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)
	dataDir := filepath.Join(tmpDir, "mnemosyne")

	if err := os.MkdirAll(filepath.Join(dataDir, "lib"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(dataDir, "models"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	dbPath := filepath.Join(dataDir, "mnemosyne.db")
	if err := os.WriteFile(dbPath, []byte("fake-db"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// Cleanup with --db --yes should remove everything.
	rootCmd.SetArgs([]string{"cleanup", "--db", "--yes"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "mnemosyne.db") {
		t.Errorf("expected db path in output, got: %s", output)
	}

	// Everything should be gone.
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {
		t.Error("expected database to be removed")
	}
}

func TestCleanupCmd_NothingToCleanup(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	rootCmd.SetArgs([]string{"cleanup", "--yes"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := outBuf.String()
	if !strings.Contains(output, "Nothing to clean up") {
		t.Errorf("expected 'Nothing to clean up' message, got: %s", output)
	}
}
