package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func buildCLI(t *testing.T) string {
	t.Helper()

	exePath := filepath.Join(t.TempDir(), "mnemosyne-test-bin")

	// Get current working directory for CGO_LDFLAGS
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}

	cmd := exec.Command("go", "build", "-tags", "sqlite_fts5", "-o", exePath, "main.go")

	// Add required CGO_LDFLAGS to find libtokenizers
	cmd.Env = append(os.Environ(), fmt.Sprintf("CGO_LDFLAGS=-L%s", filepath.Join(cwd, "lib")))

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build mnemosyne: %v\nOutput:\n%s", err, string(out))
	}

	return exePath
}

func runCLI(t *testing.T, binPath string, env []string, args ...string) (string, string) {
	t.Helper()

	cmd := exec.Command(binPath, args...)
	cmd.Env = append(os.Environ(), env...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %s %v\nError: %v\nStdout: %s\nStderr: %s", binPath, args, err, stdout.String(), stderr.String())
	}

	return stdout.String(), stderr.String()
}

func TestIntegrationPipeline(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Build the CLI tool
	binPath := buildCLI(t)

	// Prepare a temporary directory for the isolated database and models
	tempDir := t.TempDir()
	dbPath := filepath.Join(tempDir, "test.db")

	// We use the default data dir so we don't have to download models again in CI!
	// Just isolate the DB path.
	env := []string{
		fmt.Sprintf("MNEMOSYNE_DB_PATH=%s", dbPath),
	}

	// Test 1: Initialize a new collection
	colName := "test_collection"
	t.Log("Initializing collection...")
	stdout, stderr := runCLI(t, binPath, env, "init", "--name", colName)
	output := stdout + stderr
	if !strings.Contains(output, "Created collection") && !strings.Contains(output, "already exists") {
		t.Fatalf("Expected 'Created collection' or 'already exists' in output, got: %s", output)
	}

	// Test 2: Add document
	t.Log("Adding document...")
	stdout, stderr = runCLI(t, binPath, env, "add", "--name", colName, "This is a test document about artificial intelligence.")
	output = stdout + stderr
	if !strings.Contains(output, "Added document") {
		t.Fatalf("Expected 'Added document' in output, got: %s", output)
	}

	// Test 3: Search document
	t.Log("Searching document...")
	stdout, stderr = runCLI(t, binPath, env, "search", "--name", colName, "artificial intelligence")
	output = stdout + stderr
	if !strings.Contains(output, "This is a test document") {
		t.Fatalf("Expected to find the document in search results, got: %s", output)
	}

	// Test 4: Collections list
	t.Log("Listing collections...")
	stdout, stderr = runCLI(t, binPath, env, "collections")
	output = stdout + stderr
	if !strings.Contains(output, colName) {
		t.Fatalf("Expected to find the collection in list, got: %s", output)
	}

	// Test 5: Stats
	t.Log("Getting stats...")
	stdout, stderr = runCLI(t, binPath, env, "stats")
	output = stdout + stderr
	if !strings.Contains(output, "Total Documents:") {
		t.Fatalf("Expected 'Total Documents:', got: %s", output)
	}
}

func TestMainCoverage(t *testing.T) {
	// Calling main will execute os.Exit, we just want to ensure it compiles correctly
	// Since this tool executes the root command it's better to just skip main.go for direct coverage tests
	// or test the actual `cmd.Execute()` independently
}
