package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd removes a document by its ID.
var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a document by ID",
	Long:  `Remove a document and its embedding vector from the database by its numeric ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid document ID %q: must be a number", args[0])
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close() //nolint:errcheck

		// Fetch the document first so we can show what was deleted.
		doc, err := database.GetDocumentByID(id)
		if err != nil {
			return fmt.Errorf("looking up document: %w", err)
		}
		if doc == nil {
			return fmt.Errorf("document %d not found", id)
		}

		// Delete the embedding vector first (sqlite-vec doesn't support CASCADE).
		// We ignore errors here because the document may not have a vector
		// (e.g. documents added before Phase 5).
		_ = database.DeleteVector(id)

		if err := database.DeleteDocument(id); err != nil {
			return fmt.Errorf("deleting document: %w", err)
		}

		preview := doc.Content
		if len(preview) > 80 {
			preview = preview[:80] + "..."
		}

		cmd.Printf("Deleted document %d\n", id)
		cmd.Printf("  %s\n", preview)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
