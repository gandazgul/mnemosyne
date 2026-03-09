package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveCollectionName(t *testing.T) {
	// 1. Name provided explicitly
	name, err := resolveCollectionName("explicit_name")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if name != "explicit_name" {
		t.Errorf("expected 'explicit_name', got %v", name)
	}

	// 2. No name provided (uses cwd)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %v", err)
	}
	expected := filepath.Base(cwd)
	
	name, err = resolveCollectionName("")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if name != expected {
		t.Errorf("expected '%s', got '%s'", expected, name)
	}
}

func TestOpenDB(t *testing.T) {
	// Temporarily override DBPath to a test file so we don't clobber the real DB
	tmpDB := filepath.Join(t.TempDir(), "test.db")
	t.Setenv("MNEMOSYNE_DB_PATH", tmpDB)

	db, err := openDB()
	if err != nil {
		t.Fatalf("expected no error opening test DB, got %v", err)
	}
	defer db.Close()
	
	if db == nil {
		t.Error("expected valid DB instance, got nil")
	}
}
