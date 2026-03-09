package cmd

import (
	"bytes"
	"testing"
)

func TestSearchCmd_Empty(t *testing.T) {
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
