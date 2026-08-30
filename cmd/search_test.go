package cmd

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestSearchCmd_GlobalFreshDatabaseFTSOnlyLazilyInitializes(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(tmpDir, "mnemoteca.db"))
	resetSearchFlagsForTest(t)

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"search", "--global", "--fts-only", "--no-rerank", "--format", "plain", "anything"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("fresh global fts-only search error: %v", err)
	}

	database, err := openDB()
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer database.Close() //nolint:errcheck
	if got := countCollectionsNamedForTest(t, database, "global"); got != 1 {
		t.Fatalf("global collection count = %d, want 1", got)
	}
}

func TestSearchCmd_Empty(t *testing.T) {
	resetSearchFlagsForTest(t)
	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// Error: Empty query
	rootCmd.SetArgs([]string{"search", ""})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error for empty search query")
	}
}

func TestSearchCmd_InvalidFormat(t *testing.T) {
	resetSearchFlagsForTest(t)
	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)

	// Error: Invalid format
	rootCmd.SetArgs([]string{"search", "query", "--format", "invalid-fmt"})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error for invalid format")
	}
}

func resetSearchFlagsForTest(t *testing.T) {
	t.Helper()
	for name, value := range map[string]string{
		"name":              "",
		"global":            "false",
		"limit":             "3",
		"rrf-k":             "0",
		"rerank-candidates": "0",
		"no-rerank":         "false",
		"fts-only":          "false",
		"vector-only":       "false",
		"no-threshold":      "false",
		"debug":             "false",
		"format":            "color",
	} {
		if err := searchCmd.Flags().Set(name, value); err != nil {
			t.Fatalf("reset flag %s: %v", name, err)
		}
	}

	flag := searchCmd.Flags().Lookup("tag")
	if flag == nil {
		t.Fatal("tag flag not found")
	}
	sliceValue, ok := flag.Value.(interface{ Replace([]string) error })
	if !ok {
		t.Fatal("tag flag does not support slice replacement")
	}
	if err := sliceValue.Replace(nil); err != nil {
		t.Fatalf("reset tag flag: %v", err)
	}
}
