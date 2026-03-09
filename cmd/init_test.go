package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitCmd(t *testing.T) {
	// Setup DB in temp dir
	tmpDir := t.TempDir()
	t.Setenv("MNEMOSYNE_DB_PATH", filepath.Join(tmpDir, "mnemosyne.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	
	// Test creating new collection
	rootCmd.SetArgs([]string{"init", "--name", "testcol"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	output := outBuf.String()
	if !strings.Contains(output, "Created collection \"testcol\"") {
		t.Errorf("expected 'Created collection' message, got: %s", output)
	}
	
	// Reset buffer
	outBuf.Reset()
	
	// Test existing collection
	rootCmd.SetArgs([]string{"init", "--name", "testcol"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	output = outBuf.String()
	if !strings.Contains(output, "Collection \"testcol\" already exists") {
		t.Errorf("expected 'already exists' message, got: %s", output)
	}
}
