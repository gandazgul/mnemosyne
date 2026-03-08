package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd creates or confirms a collection in the database.
// If --name is not provided, uses the current directory's base name.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new collection",
	Long: `Create a new collection for storing documents.

If --name is not provided, the base name of the current working directory
is used as the collection name.

This command is idempotent: running it again for an existing collection
simply confirms it exists.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")

		collectionName, err := resolveCollectionName(nameFlag)
		if err != nil {
			return err
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close()

		collection, created, err := database.GetOrCreateCollection(collectionName)
		if err != nil {
			return fmt.Errorf("initializing collection: %w", err)
		}

		if created {
			fmt.Printf("Created collection %q (id: %d)\n", collection.Name, collection.ID)
		} else {
			fmt.Printf("Collection %q already exists (id: %d)\n", collection.Name, collection.ID)
		}

		return nil
	},
}

func init() {
	initCmd.Flags().String("name", "", "collection name (defaults to current directory name)")
	rootCmd.AddCommand(initCmd)
}
