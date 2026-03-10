package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// forgetCmd deletes a collection and all its documents permanently.
var forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Permanently delete a collection and all its documents",
	Long: `Permanently delete a collection and all of its documents.

This action is irreversible. You will be asked to confirm by typing
the collection name unless --yes is provided.

If --name is not provided, the base name of the current working directory
is used as the collection name.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		yesFlag, _ := cmd.Flags().GetBool("yes")

		collectionName, err := resolveCollectionName(nameFlag)
		if err != nil {
			return err
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close() //nolint:errcheck

		// Verify the collection exists and show what will be deleted.
		collection, err := database.GetCollectionByName(collectionName)
		if err != nil {
			return fmt.Errorf("looking up collection: %w", err)
		}
		if collection == nil {
			return fmt.Errorf("collection %q does not exist", collectionName)
		}

		docCount, err := database.CountDocuments(collection.ID)
		if err != nil {
			return fmt.Errorf("counting documents: %w", err)
		}

		cmd.Printf("This will permanently delete collection %q and all %d of its documents.\n", collectionName, docCount)
		cmd.Println("This action cannot be undone.")

		if !yesFlag {
			cmd.Printf("\nType the collection name to confirm: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("reading confirmation: %w", err)
			}

			input = strings.TrimSpace(input)
			if input != collectionName {
				return fmt.Errorf("confirmation did not match; expected %q, got %q", collectionName, input)
			}
		}

		// Delete vectors first. sqlite-vec virtual tables don't participate in
		// CASCADE deletes, so we must clean them up explicitly before the
		// collection (and its documents) are removed.
		// We ignore errors here because the vector table may not exist yet
		// (e.g. if no documents were ever added with embeddings).
		_ = database.DeleteVectorsByCollection(collection.ID)

		if err := database.DeleteCollection(collectionName); err != nil {
			return fmt.Errorf("deleting collection: %w", err)
		}

		cmd.Printf("Deleted collection %q and %d documents.\n", collectionName, docCount)

		return nil
	},
}

func init() {
	forgetCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	forgetCmd.Flags().Bool("yes", false, "skip confirmation prompt")
	rootCmd.AddCommand(forgetCmd)
}
