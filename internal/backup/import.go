package backup

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gandazgul/mnemosyne/internal/db"
)

// batchSize is the number of documents to insert per transaction.
const batchSize = 500

// ImportCollection reads a JSONL file and inserts documents into the database.
// If overrideName is non-empty, it is used instead of the collection name from
// the file header. Creates the collection if it doesn't exist.
// Returns the header and the number of imported documents.
func ImportCollection(r io.Reader, database *db.DB, overrideName string) (*Header, int64, error) {
	scanner := bufio.NewScanner(r)

	// Increase scanner buffer for potentially large lines (vectors can be big).
	scanner.Buffer(make([]byte, 0, 1024*1024), 10*1024*1024)

	// Read header line.
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, 0, fmt.Errorf("reading header: %w", err)
		}
		return nil, 0, fmt.Errorf("empty file: no header line")
	}

	var header Header
	if err := json.Unmarshal(scanner.Bytes(), &header); err != nil {
		return nil, 0, fmt.Errorf("parsing header: %w", err)
	}
	if header.Version != FormatVersion {
		return nil, 0, fmt.Errorf("unsupported format version %d (expected %d)", header.Version, FormatVersion)
	}

	// Determine collection name.
	collectionName := header.Collection
	if overrideName != "" {
		collectionName = overrideName
	}
	if collectionName == "" {
		return nil, 0, fmt.Errorf("no collection name in header and no override provided")
	}

	collection, _, err := database.GetOrCreateCollection(collectionName)
	if err != nil {
		return nil, 0, fmt.Errorf("getting or creating collection %q: %w", collectionName, err)
	}

	// Stream and insert documents in batches.
	var imported int64
	var batch []DocRecord

	for scanner.Scan() {
		var doc DocRecord
		if err := json.Unmarshal(scanner.Bytes(), &doc); err != nil {
			return &header, imported, fmt.Errorf("parsing document at line %d: %w", imported+2, err)
		}
		batch = append(batch, doc)

		if len(batch) >= batchSize {
			n, err := insertBatch(database, collection.ID, batch)
			imported += n
			if err != nil {
				return &header, imported, err
			}
			batch = batch[:0]
		}
	}
	if err := scanner.Err(); err != nil {
		return &header, imported, fmt.Errorf("reading file: %w", err)
	}

	// Flush remaining batch.
	if len(batch) > 0 {
		n, err := insertBatch(database, collection.ID, batch)
		imported += n
		if err != nil {
			return &header, imported, err
		}
	}

	return &header, imported, nil
}

// insertBatch inserts a batch of documents with their vectors into the database.
// Each document is inserted individually using InsertDocumentWithVector which
// handles its own transaction. Returns the number of successfully inserted documents.
func insertBatch(database *db.DB, collectionID int64, docs []DocRecord) (int64, error) {
	var inserted int64
	for i, doc := range docs {
		_, err := database.InsertDocumentWithVector(collectionID, doc.Content, doc.Metadata, doc.Vector)
		if err != nil {
			return inserted, fmt.Errorf("inserting document %d in batch: %w", i+1, err)
		}
		inserted++
	}
	return inserted, nil
}
