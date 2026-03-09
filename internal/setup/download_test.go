package setup

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestFileMatchesSHA256(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.txt")
	content := []byte("hello world")
	os.WriteFile(path, content, 0644)
	
	h := sha256.New()
	h.Write(content)
	expected := hex.EncodeToString(h.Sum(nil))
	
	match, err := fileMatchesSHA256(path, expected)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !match {
		t.Error("expected checksum to match")
	}
	
	match, _ = fileMatchesSHA256(path, "invalid")
	if match {
		t.Error("expected checksum to not match")
	}
	
	// test non-existent file
	_, err = fileMatchesSHA256(filepath.Join(tmpDir, "missing.txt"), expected)
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestDownloadFile(t *testing.T) {
	// Setup a dummy HTTP server
	content := []byte("downloaded content")
	h := sha256.New()
	h.Write(content)
	expectedHash := hex.EncodeToString(h.Sum(nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(content)
	}))
	defer ts.Close()

	tmpDir := t.TempDir()
	destPath := filepath.Join(tmpDir, "download.txt")

	// Test 1: Full download with checksum
	var lastWritten int64
	err := downloadFile(context.Background(), ts.URL, destPath, expectedHash, func(written, total int64) {
		lastWritten = written
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if lastWritten != int64(len(content)) {
		t.Errorf("expected %d bytes written, got %d", len(content), lastWritten)
	}

	// Test 2: Skip existing file
	lastWritten = 0 // Reset
	err = downloadFile(context.Background(), ts.URL, destPath, expectedHash, func(written, total int64) {
		lastWritten = written
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if lastWritten != 0 {
		t.Error("expected download to be skipped, but progress was called")
	}

	// Test 3: Download failure (bad checksum)
	badPath := filepath.Join(tmpDir, "bad.txt")
	err = downloadFile(context.Background(), ts.URL, badPath, "invalidhash", nil)
	if err == nil {
		t.Error("expected error for bad checksum")
	}
}

func TestDownloadFile_Errors(t *testing.T) {
	tmpDir := t.TempDir()
	
	// Test HTTP error
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()
	
	err := downloadFile(context.Background(), ts.URL, filepath.Join(tmpDir, "404.txt"), "", nil)
	if err == nil {
		t.Error("expected error for 404 response")
	}
	
	// Test invalid URL
	err = downloadFile(context.Background(), "http://invalid.url.that.doesnt.exist", filepath.Join(tmpDir, "invalid.txt"), "", nil)
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}
