package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd updates an existing document in the specified collection.
var updateCmd = &cobra.Command{
	Use:   "update <id> [text]",
	Short: "Update an existing document by ID",
	Long: `Update an existing document in the selected collection.

Text can be provided as positional arguments, read from a file with --file,
or piped via stdin with --stdin. Replacement text is stored as one document
and is not chunked.

The command is strict: the document must already exist in the selected
collection. It never inserts a new document or moves documents between
collections.

Tags are additive by default with --tag. Use --replace-tags to replace the
existing tag set; --replace-tags without --tag clears all tags. Non-tag metadata
is preserved.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a document ID")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid document ID %q: must be a number", args[0])
		}

		nameFlag, _ := cmd.Flags().GetString("name")
		globalFlag, _ := cmd.Flags().GetBool("global")
		fileFlag, _ := cmd.Flags().GetString("file")
		stdinFlag, _ := cmd.Flags().GetBool("stdin")
		tagsFlag, _ := cmd.Flags().GetStringSlice("tag")
		replaceTagsFlag, _ := cmd.Flags().GetBool("replace-tags")

		replacementContent, hasContent, err := readUpdateContent(args[1:], fileFlag, stdinFlag)
		if err != nil {
			return err
		}

		tagOperation := len(tagsFlag) > 0 || replaceTagsFlag
		if !hasContent && !tagOperation {
			return fmt.Errorf("no update provided; pass text, use --file/--stdin, add --tag, or use --replace-tags")
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

		collection, err := getSelectedCollection(database, collectionName, globalFlag)
		if err != nil {
			return err
		}

		doc, err := database.GetDocumentByID(id)
		if err != nil {
			return fmt.Errorf("looking up document: %w", err)
		}
		if doc == nil || doc.CollectionID != collection.ID {
			return fmt.Errorf("document %d not found in collection %q", id, collectionName)
		}

		metadata := doc.Metadata
		if tagOperation {
			metadata, err = updateMetadataTags(doc.Metadata, tagsFlag, replaceTagsFlag)
			if err != nil {
				return err
			}
		}

		preview := doc.Content
		if hasContent {
			embedder, cfg, err := openEmbedder(cmd.Context())
			if err != nil {
				return fmt.Errorf("loading embedding model: %w", err)
			}
			defer embedder.Close() //nolint:errcheck

			if err := database.EnsureVectorTable(cfg.Embedding.Dimensions); err != nil {
				return fmt.Errorf("ensuring vector table: %w", err)
			}

			vec, err := embedder.EmbedDocument(replacementContent)
			if err != nil {
				return fmt.Errorf("embedding document: %w", err)
			}

			if err := database.UpdateDocumentWithVector(id, collection.ID, replacementContent, metadata, vec); err != nil {
				return fmt.Errorf("updating document: %w", err)
			}
			preview = replacementContent
		} else {
			if err := database.UpdateDocumentMetadata(id, collection.ID, metadata); err != nil {
				return fmt.Errorf("updating document: %w", err)
			}
		}

		if len(preview) > 80 {
			preview = strings.ReplaceAll(preview[:80], "\n", " ") + "..."
		}

		cmd.Printf("Updated document %d in collection %q\n", id, collectionName)
		cmd.Printf("  %s\n", preview)

		return nil
	},
}

func readUpdateContent(textArgs []string, fileFlag string, stdinFlag bool) (content string, hasContent bool, err error) {
	sources := 0
	if len(textArgs) > 0 {
		sources++
	}
	if fileFlag != "" {
		sources++
	}
	if stdinFlag {
		sources++
	}
	if sources == 0 {
		return "", false, nil
	}
	if sources > 1 {
		return "", false, fmt.Errorf("provide only one content source: positional text, --file, or --stdin")
	}

	switch {
	case fileFlag != "":
		data, err := os.ReadFile(fileFlag)
		if err != nil {
			return "", false, fmt.Errorf("reading file %s: %w", fileFlag, err)
		}
		content = strings.TrimSpace(string(data))

	case stdinFlag:
		var lines []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return "", false, fmt.Errorf("reading stdin: %w", err)
		}
		content = strings.TrimSpace(strings.Join(lines, "\n"))

	default:
		content = strings.TrimSpace(strings.Join(textArgs, " "))
	}

	if content == "" {
		return "", false, fmt.Errorf("content is empty")
	}
	return content, true, nil
}

func updateMetadataTags(existing *string, tags []string, replace bool) (*string, error) {
	metadata := make(map[string]interface{})
	if existing != nil && strings.TrimSpace(*existing) != "" {
		if err := json.Unmarshal([]byte(*existing), &metadata); err != nil {
			return nil, fmt.Errorf("parsing document metadata: %w", err)
		}
		if metadata == nil {
			return nil, fmt.Errorf("document metadata must be a JSON object")
		}
	}

	var existingTags []string
	if value, ok := metadata["tags"]; ok {
		tagValues, ok := value.([]interface{})
		if !ok {
			return nil, fmt.Errorf("document metadata tags must be an array of strings")
		}
		for _, tagValue := range tagValues {
			tag, ok := tagValue.(string)
			if !ok {
				return nil, fmt.Errorf("document metadata tags must be an array of strings")
			}
			existingTags = append(existingTags, tag)
		}
	}

	merged := dedupeTags(tags)
	if !replace {
		merged = appendUniqueTags(existingTags, tags)
	}

	if len(merged) == 0 {
		delete(metadata, "tags")
	} else {
		metadata["tags"] = merged
	}

	if len(metadata) == 0 {
		return nil, nil
	}

	b, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("encoding metadata: %w", err)
	}
	s := string(b)
	return &s, nil
}

func appendUniqueTags(existing, additions []string) []string {
	seen := make(map[string]bool, len(existing)+len(additions))
	merged := make([]string, 0, len(existing)+len(additions))
	for _, tag := range existing {
		if seen[tag] {
			continue
		}
		seen[tag] = true
		merged = append(merged, tag)
	}
	for _, tag := range additions {
		if seen[tag] {
			continue
		}
		seen[tag] = true
		merged = append(merged, tag)
	}
	return merged
}

func dedupeTags(tags []string) []string {
	return appendUniqueTags(nil, tags)
}

func init() {
	updateCmd.Flags().StringP("name", "n", "", "collection name (defaults to current directory name)")
	updateCmd.Flags().BoolP("global", "g", false, "use the global collection")
	updateCmd.Flags().String("file", "", "read replacement content from a file")
	updateCmd.Flags().Bool("stdin", false, "read replacement content from stdin")
	updateCmd.Flags().StringSliceP("tag", "t", nil, "add one or more tags to the document (or replace with --replace-tags)")
	updateCmd.Flags().Bool("replace-tags", false, "replace existing tags with --tag values; without --tag, clear all tags")
	rootCmd.AddCommand(updateCmd)
}
