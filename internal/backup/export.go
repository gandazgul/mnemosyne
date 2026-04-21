package backup

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gandazgul/mnemosyne/internal/db"
)

// ExportCollection streams a single collection to a JSONL writer.
// When skipEmbeddings is true, vectors are omitted from the output.
// Returns the number of documents exported.
func ExportCollection(w io.Writer, database *db.DB, collectionName string, skipEmbeddings bool) (int64, error) {
	collection, err := database.GetCollectionByName(collectionName)
	if err != nil {
		return 0, fmt.Errorf("looking up collection: %w", err)
	}
	if collection == nil {
		return 0, fmt.Errorf("collection %q does not exist", collectionName)
	}

	docCount, err := database.CountDocuments(collection.ID, nil)
	if err != nil {
		return 0, fmt.Errorf("counting documents: %w", err)
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	// Write header line.
	header := Header{
		Version:    FormatVersion,
		ExportedAt: time.Now().UTC().Format(time.RFC3339),
		Collection: collectionName,
		DocCount:   docCount,
	}
	if err := enc.Encode(header); err != nil {
		return 0, fmt.Errorf("writing header: %w", err)
	}

	// Stream documents.
	var exported int64
	if skipEmbeddings {
		err = database.StreamDocuments(collection.ID, func(rec db.ExportRecord) error {
			doc := DocRecord{
				Content:  rec.Content,
				Metadata: rec.Metadata,
			}
			if err := enc.Encode(doc); err != nil {
				return fmt.Errorf("writing document: %w", err)
			}
			exported++
			return nil
		})
	} else {
		err = database.StreamDocumentsWithVectors(collection.ID, func(rec db.ExportRecord) error {
			doc := DocRecord{
				Content:  rec.Content,
				Metadata: rec.Metadata,
				Vector:   rec.Vector,
			}
			if err := enc.Encode(doc); err != nil {
				return fmt.Errorf("writing document: %w", err)
			}
			exported++
			return nil
		})
	}
	if err != nil {
		return exported, fmt.Errorf("streaming documents: %w", err)
	}

	return exported, nil
}
