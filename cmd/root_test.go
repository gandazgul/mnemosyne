package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestExecuteRootCommand(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{})
	err := Execute()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !strings.Contains(output, "Welcome to Mnemosyne!") {
		t.Errorf("Expected welcome message, got %s", output)
	}
}
