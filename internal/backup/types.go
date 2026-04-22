// Package backup provides JSONL-based export and import of mnemosyne collections.
//
// The JSONL format uses one JSON object per line:
//   - Line 1: Header with version, timestamp, collection name, and document count.
//   - Lines 2–N: One DocRecord per document with content, metadata, and raw vector.
package backup

// FormatVersion is the current JSONL export format version.
const FormatVersion = 1

// Header is the first line of a JSONL export file.
type Header struct {
	Version    int    `json:"version"`
	ExportedAt string `json:"exported_at"` // RFC 3339
	Collection string `json:"collection"`
	DocCount   int64  `json:"doc_count"`
}

// DocRecord is one document line in a JSONL export file.
type DocRecord struct {
	Content            string    `json:"content"`
	Metadata           *string   `json:"metadata,omitempty"`
	Vector             []float32 `json:"vector,omitempty"`
	OriginalDocumentID int64     `json:"original_document_id,omitempty"`
}
