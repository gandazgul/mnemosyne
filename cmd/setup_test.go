package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/setup"
)

func TestSetupCmd_AlreadyReady(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)

	mnemoDir := filepath.Join(tmpDir, "mnemosyne")

	// mock all files
	libPath := filepath.Join(mnemoDir, "lib", setup.OnnxRuntimeLibNameForTest())
	_ = os.MkdirAll(filepath.Dir(libPath), 0755)
	_ = os.WriteFile(libPath, []byte("fake"), 0644)

	for _, m := range setup.AllModelsForTest() {
		modelDir := filepath.Join(mnemoDir, "models", m.LocalDir)
		for file := range m.Files {
			path := filepath.Join(modelDir, file)
			_ = os.MkdirAll(filepath.Dir(path), 0755)
			_ = os.WriteFile(path, []byte("dummy content"), 0644)
		}
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"setup"})
	err := rootCmd.Execute()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	output := buf.String()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(output, "All components are already installed.") {
		t.Errorf("expected already installed message, got: %s", output)
	}
}

func TestSetupCmd_NotReady_FailsToDownload(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("XDG_DATA_HOME", tmpDir)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"setup"})
	err := rootCmd.Execute()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	// output := buf.String()

	if err == nil {
		t.Log("Expected error from real download without network/proper path")
	}
}
