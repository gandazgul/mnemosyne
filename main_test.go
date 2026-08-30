package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func buildCLI(t *testing.T) string {
	t.Helper()

	exePath := filepath.Join(t.TempDir(), "mnemoteca-test-bin")

	// Get current working directory for CGO_LDFLAGS
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}

	cmd := exec.Command("go", "build", "-tags", "sqlite_fts5", "-o", exePath, "main.go")

	// Add CGO_LDFLAGS so tests can find native libraries in the local lib dir.
	cmd.Env = append(os.Environ(), fmt.Sprintf("CGO_LDFLAGS=-L%s", filepath.Join(cwd, "lib")))

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build mnemoteca: %v\nOutput:\n%s", err, string(out))
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

func prepareIntegrationDataDir(t *testing.T, tempDir string) {
	t.Helper()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}

	dataDir := filepath.Join(tempDir, "mnemoteca")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Fatalf("Failed to create test data dir: %v", err)
	}

	for _, name := range []string{"lib", "models"} {
		src := filepath.Join(cwd, name)
		if info, err := os.Stat(src); err != nil || !info.IsDir() {
			t.Fatalf("Expected %s directory to exist; run task setup or the CI setup steps before integration tests", src)
		}

		dst := filepath.Join(dataDir, name)
		if err := os.Symlink(src, dst); err != nil {
			t.Fatalf("Failed to link %s into test data dir: %v", name, err)
		}
	}
}

func addedDocumentID(t *testing.T, output string) string {
	t.Helper()

	for _, line := range strings.Split(output, "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 3 && fields[0] == "Added" && fields[1] == "document" {
			if _, err := strconv.Atoi(fields[2]); err != nil {
				t.Fatalf("Expected added document ID to be numeric in output, got: %s", output)
			}
			return fields[2]
		}
	}

	t.Fatalf("Expected add output to include an 'Added document <id>' line, got: %s", output)
	return ""
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
	prepareIntegrationDataDir(t, tempDir)

	env := []string{
		fmt.Sprintf("XDG_DATA_HOME=%s", tempDir),
		fmt.Sprintf("MNEMOTECA_DB_PATH=%s", dbPath),
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

	docID := addedDocumentID(t, output)

	// Test 3: Search document
	t.Log("Searching document...")
	stdout, stderr = runCLI(t, binPath, env, "search", "--name", colName, "artificial intelligence")
	output = stdout + stderr
	if !strings.Contains(output, "This is a test document") {
		t.Fatalf("Expected to find the document in search results, got: %s", output)
	}

	// Test 4: Update document
	t.Log("Updating document...")
	stdout, stderr = runCLI(t, binPath, env, "update", "--name", colName, docID, "This revised document is about knowledge management.")
	output = stdout + stderr
	if !strings.Contains(output, "Updated document") {
		t.Fatalf("Expected 'Updated document' in output, got: %s", output)
	}

	// Test 5: Search updated document
	t.Log("Searching updated document...")
	stdout, stderr = runCLI(t, binPath, env, "search", "--name", colName, "knowledge management")
	output = stdout + stderr
	if !strings.Contains(output, "This revised document") {
		t.Fatalf("Expected to find the updated document in search results, got: %s", output)
	}

	// Test 6: Collections list
	t.Log("Listing collections...")
	stdout, stderr = runCLI(t, binPath, env, "collections")
	output = stdout + stderr
	if !strings.Contains(output, colName) {
		t.Fatalf("Expected to find the collection in list, got: %s", output)
	}

	// Test 7: Stats
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
