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
		globalFlag, _ := cmd.Flags().GetBool("global")
		isDefault := nameFlag == "" && !globalFlag

		collectionName, err := resolveCollectionName(nameFlag, globalFlag)
		if err != nil {
			return err
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close() //nolint:errcheck

		// Prevent accidentally linking to an existing collection when using the directory name.
		if isDefault {
			existing, err := database.GetCollectionByName(collectionName)
			if err != nil {
				return fmt.Errorf("checking collection existence: %w", err)
			}
			if existing != nil {
				cmd.SilenceUsage = true // Don't show help for this semantic error
				return fmt.Errorf("collection %q (derived from current directory) already exists. Please pass --name or -n to specify a different name, or pass explicitly to confirm", collectionName)
			}
		}

		collection, created, err := database.GetOrCreateCollection(collectionName)
		if err != nil {
			return fmt.Errorf("initializing collection: %w", err)
		}

		if created {
			cmd.Printf("Created collection %q (id: %d)\n", collection.Name, collection.ID)
		} else {
			cmd.Printf("Collection %q already exists (id: %d)\n", collection.Name, collection.ID)
		}

		return nil
	},
}

func init() {
	initCmd.Flags().StringP("name", "n", "", "collection name (defaults to current directory name)")
	initCmd.Flags().BoolP("global", "g", false, "use the global collection")
	rootCmd.AddCommand(initCmd)
}
