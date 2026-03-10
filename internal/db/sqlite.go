// Package db provides SQLite database access for mnemosyne.
//
// This package manages the database connection, schema migrations, and provides
// repositories for collections and documents. It uses mattn/go-sqlite3 with CGO,
// which requires a C compiler (GCC/Clang) to build.
//
// The database uses three main components:
//   - collections table: groups of documents
//   - documents table: the actual stored content, linked to a collection
//   - FTS5 + sqlite-vec virtual tables: added in later phases for search
package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	// Import sqlite-vec CGO bindings. Auto() registers the sqlite-vec extension
	// as an auto-loaded extension with mattn/go-sqlite3, so every new connection
	// gets vector search functions (vec0, vec_version, etc.) available.
	sqlite_vec "github.com/asg017/sqlite-vec-go-bindings/cgo"

	// Import the SQLite driver. The underscore means we import it only for its
	// side effect: registering itself as a database/sql driver named "sqlite3".
	// This is a common Go pattern for database drivers.
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Register sqlite-vec as an auto-loaded extension. This must happen before
	// any sql.Open("sqlite3", ...) call. Auto() tells mattn/go-sqlite3 to load
	// the sqlite-vec extension into every new SQLite connection automatically.
	sqlite_vec.Auto()
}

// DB wraps a *sql.DB connection and provides access to repositories.
type DB struct {
	conn *sql.DB
}

// Open creates or opens a SQLite database at the given path.
// It creates parent directories if they don't exist, runs migrations,
// and enables WAL mode and foreign keys for better performance and data integrity.
func Open(dbPath string) (*DB, error) {
	// Ensure the parent directory exists.
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("creating database directory %s: %w", dir, err)
	}

	// Open the database connection.
	// The _journal_mode=WAL and _foreign_keys=on pragmas are set via DSN parameters
	// so they apply immediately when the connection is established.
	dsn := fmt.Sprintf("file:%s?_journal_mode=WAL&_foreign_keys=on", dbPath)
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	// Verify the connection works.
	if err := conn.Ping(); err != nil {
		conn.Close() //nolint:errcheck
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	db := &DB{conn: conn}

	// Run schema migrations to ensure tables exist.
	if err := db.migrate(); err != nil {
		conn.Close() //nolint:errcheck
		return nil, fmt.Errorf("running migrations: %w", err)
	}

	return db, nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.conn.Close()
}

// Conn returns the underlying *sql.DB connection.
// This is useful for tests and for packages that need direct SQL access.
func (db *DB) Conn() *sql.DB {
	return db.conn
}

// migrate runs all schema migrations.
// For now this is a simple "create if not exists" approach.
// A proper migration system (versioned migrations) can be added later if needed.
func (db *DB) migrate() error {
	migrations := []string{
		// Collections table: each collection is a named group of documents.
		`CREATE TABLE IF NOT EXISTS collections (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,

		// Documents table: stores the actual content, linked to a collection.
		// ON DELETE CASCADE means deleting a collection removes all its documents.
		`CREATE TABLE IF NOT EXISTS documents (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			collection_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			metadata TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE
		)`,

		// Index on collection_id for faster lookups when listing/searching
		// within a collection.
		`CREATE INDEX IF NOT EXISTS idx_documents_collection
			ON documents(collection_id)`,

		// FTS5 virtual table for full-text search with BM25 ranking.
		// content=documents means this is a "content sync" (external content) table
		// that mirrors the documents table. content_rowid=id maps FTS rowids to
		// documents.id so we can join back to the source table.
		`CREATE VIRTUAL TABLE IF NOT EXISTS docs_fts USING fts5(
			content,
			content=documents,
			content_rowid=id
		)`,

		// Trigger: keep FTS index in sync after INSERT on documents.
		`CREATE TRIGGER IF NOT EXISTS docs_fts_ai AFTER INSERT ON documents BEGIN
			INSERT INTO docs_fts(rowid, content) VALUES (new.id, new.content);
		END`,

		// Trigger: keep FTS index in sync after DELETE on documents.
		// The special 'delete' command tells FTS5 to remove the old entry.
		`CREATE TRIGGER IF NOT EXISTS docs_fts_ad AFTER DELETE ON documents BEGIN
			INSERT INTO docs_fts(docs_fts, rowid, content) VALUES('delete', old.id, old.content);
		END`,

		// Trigger: keep FTS index in sync after UPDATE on documents.
		// Delete the old entry, then insert the new one.
		`CREATE TRIGGER IF NOT EXISTS docs_fts_au AFTER UPDATE ON documents BEGIN
			INSERT INTO docs_fts(docs_fts, rowid, content) VALUES('delete', old.id, old.content);
			INSERT INTO docs_fts(rowid, content) VALUES (new.id, new.content);
		END`,
	}

	for _, m := range migrations {
		if _, err := db.conn.Exec(m); err != nil {
			return fmt.Errorf("executing migration: %w\nSQL: %s", err, m)
		}
	}

	return nil
}
