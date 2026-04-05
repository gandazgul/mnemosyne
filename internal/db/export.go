package db

import (
	"encoding/binary"
	"fmt"
	"math"
)

// ExportRecord holds one document and its raw vector for export.
type ExportRecord struct {
	Content  string
	Metadata *string
	Vector   []float32
}

// StreamDocumentsWithVectors calls fn for each document in the collection,
// joining with the vector table to include the raw embedding.
// Streaming avoids loading the full collection into memory.
func (db *DB) StreamDocumentsWithVectors(collectionID int64, fn func(ExportRecord) error) (err error) {
	rows, err := db.conn.Query(`
		SELECT d.content, d.metadata, v.embedding
		FROM documents d
		LEFT JOIN docs_vec v ON v.document_id = d.id
		WHERE d.collection_id = ?
		ORDER BY d.id`, collectionID)
	if err != nil {
		return fmt.Errorf("querying documents with vectors: %w", err)
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var rec ExportRecord
		var vecBlob []byte
		if err := rows.Scan(&rec.Content, &rec.Metadata, &vecBlob); err != nil {
			return fmt.Errorf("scanning export record: %w", err)
		}
		if vecBlob != nil {
			rec.Vector = DeserializeFloat32(vecBlob)
		}
		if err := fn(rec); err != nil {
			return fmt.Errorf("processing export record: %w", err)
		}
	}

	return rows.Err()
}

// DeserializeFloat32 converts a little-endian byte slice back to []float32.
// This is the inverse of SerializeFloat32.
func DeserializeFloat32(buf []byte) []float32 {
	n := len(buf) / 4
	v := make([]float32, n)
	for i := 0; i < n; i++ {
		v[i] = math.Float32frombits(binary.LittleEndian.Uint32(buf[i*4:]))
	}
	return v
}
