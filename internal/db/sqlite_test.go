package db_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/db"
)

// testDB creates a temporary database for testing and returns it along with
// a cleanup function. The database is created in a temp directory and
// automatically removed when cleanup is called.
func testDB(t *testing.T) *db.DB {
	t.Helper()

	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	database, err := db.Open(dbPath)
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	t.Cleanup(func() {
		database.Close()
	})

	return database
}

func TestOpen_CreatesDirectory(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "subdir", "nested", "test.db")

	database, err := db.Open(dbPath)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer database.Close()

	// Verify the file was created.
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Error("database file was not created")
	}
}

func TestOpen_WALMode(t *testing.T) {
	database := testDB(t)

	var journalMode string
	err := database.Conn().QueryRow("PRAGMA journal_mode").Scan(&journalMode)
	if err != nil {
		t.Fatalf("querying journal_mode: %v", err)
	}

	if journalMode != "wal" {
		t.Errorf("journal_mode = %q, want %q", journalMode, "wal")
	}
}

func TestOpen_ForeignKeysEnabled(t *testing.T) {
	database := testDB(t)

	var fk int
	err := database.Conn().QueryRow("PRAGMA foreign_keys").Scan(&fk)
	if err != nil {
		t.Fatalf("querying foreign_keys: %v", err)
	}

	if fk != 1 {
		t.Errorf("foreign_keys = %d, want 1", fk)
	}
}

func TestOpen_TablesCreated(t *testing.T) {
	database := testDB(t)

	// Check that the collections table exists.
	var name string
	err := database.Conn().QueryRow(
		"SELECT name FROM sqlite_master WHERE type='table' AND name='collections'",
	).Scan(&name)
	if err != nil {
		t.Fatalf("collections table not found: %v", err)
	}

	// Check that the documents table exists.
	err = database.Conn().QueryRow(
		"SELECT name FROM sqlite_master WHERE type='table' AND name='documents'",
	).Scan(&name)
	if err != nil {
		t.Fatalf("documents table not found: %v", err)
	}
}

func TestOpen_InvalidDirectory(t *testing.T) {
	dir := t.TempDir()
	// Create a file where a directory should be.
	filePath := filepath.Join(dir, "notadir")
	os.WriteFile(filePath, []byte("test"), 0644)
	
	dbPath := filepath.Join(filePath, "test.db")
	
	_, err := db.Open(dbPath)
	if err == nil {
		t.Error("expected error when opening DB with invalid directory path")
	}
}

func TestOpen_InvalidDatabaseFile(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "corrupt.db")
	
	// Create a corrupted file.
	os.WriteFile(dbPath, []byte("this is not a sqlite database"), 0644)
	
	_, err := db.Open(dbPath)
	if err == nil {
		t.Error("expected error when opening corrupted DB file")
	}
}
