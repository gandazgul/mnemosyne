// Package cmd contains all CLI commands for mnemosyne.
//
// This package uses the Cobra library to define a tree of commands.
// The root command is the parent of all subcommands (add, search, list, etc.).
// Each command lives in its own file for clarity.
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const description = "A local document store with hybrid search"

// rootCmd is the base command when called without any subcommands.
// Running `mnemosyne` with no args prints a welcome message and usage info.
var rootCmd = &cobra.Command{
	// Use is how the command appears in help text.
	Use:   "mnemosyne",
	Short: description,
	Long: `Mnemosyne - Local Document Storage & Retrieval

Store small documents (sentences to paragraphs) and retrieve them
using hybrid search: full-text (BM25) + vector similarity (cosine),
combined with Reciprocal Rank Fusion and local cross-encoder reranking.

All ML inference runs locally via ONNX Runtime. No cloud APIs needed.`,

	// Run is called when the command is executed with no subcommands.
	// This is a good place for a welcome message or default behavior.
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Mnemosyne!")
		fmt.Println()
		fmt.Println(description)
		fmt.Println("Run 'mnemosyne --help' to see available commands.")
	},
}

// Execute runs the root command. This is called from main.go.
// It parses command-line arguments and dispatches to the right subcommand.
// Returns an error if command execution fails.
func Execute() error {
	return rootCmd.Execute()
}
