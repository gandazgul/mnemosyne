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
	Long:  `Remove a document from the database by its numeric ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid document ID %q: must be a number", args[0])
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer func() {
			if cerr := database.Close(); cerr != nil && err == nil {
				err = cerr
			}
		}()

		// Fetch the document first so we can show what was deleted.
		doc, err := database.GetDocumentByID(id)
		if err != nil {
			return fmt.Errorf("looking up document: %w", err)
		}
		if doc == nil {
			return fmt.Errorf("document %d not found", id)
		}

		if err := database.DeleteDocument(id); err != nil {
			return fmt.Errorf("deleting document: %w", err)
		}

		preview := doc.Content
		if len(preview) > 80 {
			preview = preview[:80] + "..."
		}

		fmt.Printf("Deleted document %d\n", id)
		fmt.Printf("  %s\n", preview)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
