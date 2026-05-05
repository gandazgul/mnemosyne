package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List all tags used in a collection",
	Long: `List all distinct tags used in documents within the specified collection.

If --name is not provided, the current directory name is used.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		nameFlag, _ := cmd.Flags().GetString("name")
		globalFlag, _ := cmd.Flags().GetBool("global")
		formatFlag, _ := cmd.Flags().GetString("format")

		if err := validateFormat(formatFlag); err != nil {
			return err
		}
		if plain(formatFlag) {
			color.NoColor = true
		}

		collectionName, err := resolveCollectionName(nameFlag, globalFlag)
		if err != nil {
			return err
		}

		database, err := openDB()
		if err != nil {
			return err
		}
		defer database.Close() //nolint:errcheck

		collection, err := database.GetCollectionByName(collectionName)
		if err != nil {
			return fmt.Errorf("looking up collection: %w", err)
		}
		if collection == nil {
			return fmt.Errorf("collection %q does not exist; run 'mnemosyne init --name %s' first",
				collectionName, collectionName)
		}

		tags, err := database.GetTags(collection.ID)
		if err != nil {
			return fmt.Errorf("listing tags: %w", err)
		}

		if len(tags) == 0 {
			cmd.Printf("No tags found in collection %q\n", collectionName)
			return nil
		}

		if plain(formatFlag) {
			cmd.Printf("Tags in %q:\n", collectionName)
			for _, tag := range tags {
				cmd.Printf("- %s\n", tag)
			}
		} else {
			cmd.Printf("%s %s:\n", boldWhite("Tags in"), boldCyan(collectionName))
			cmd.Println(dimWhite(strings.Repeat("─", 40)))
			for _, tag := range tags {
				cmd.Printf("  %s\n", green(tag))
			}
		}

		return nil
	},
}

func init() {
	tagsCmd.Flags().StringP("name", "n", "", "collection name (defaults to current directory name)")
	tagsCmd.Flags().BoolP("global", "g", false, "use the global collection")
	tagsCmd.Flags().StringP("format", "f", "color", "output format: color (default) or plain")
	rootCmd.AddCommand(tagsCmd)
}
