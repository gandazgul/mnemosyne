// Package cmd contains all CLI commands for mnemoteca.
//
// This package uses the Cobra library to define a tree of commands.
// The root command is the parent of all subcommands (add, search, list, etc.).
// Each command lives in its own file for clarity.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const description = "A local document store with hybrid search"

// rootCmd is the base command when called without any subcommands.
// Running `mnemoteca` with no args prints a welcome message and usage info.
var rootCmd = &cobra.Command{
	// Use is how the command appears in help text.
	Use:           "mnemoteca",
	Short:         description,
	SilenceUsage:  true, // Don't print usage when RunE returns an error
	SilenceErrors: true, // We handle printing errors in main.go
	Long: `Mnemoteca - Local Document Storage & Retrieval

Store small documents (sentences to paragraphs) and retrieve them
using hybrid search: vector similarity (cosine) with a small BM25
lexical boost, plus optional local cross-encoder reranking.

All ML inference runs locally via ONNX Runtime. No cloud APIs needed.

Configuration:
  Mnemoteca reads ~/.config/mnemoteca/config.yaml by default.
  Set MNEMOTECA_CONFIG=/path/to/config.yaml to use another config file.
  Set MNEMOTECA_DB_PATH=/path/to/mnemoteca.db to override the database path.`,

	// Run is called when the command is executed with no subcommands.
	// This is a good place for a welcome message or default behavior.
	Run: func(cmd *cobra.Command, args []string) {
		// Check if --version or -v flag was passed
		if version, _ := cmd.Flags().GetBool("version"); version {
			printVersion(cmd.OutOrStdout())
			return
		}
		fmt.Println("Welcome to Mnemoteca!")
		fmt.Println()
		fmt.Println(description)
		fmt.Println("Run 'mnemoteca --help' to see available commands.")
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version of mnemoteca")
}

// Execute runs the root command. This is called from main.go.
// It parses command-line arguments and dispatches to the right subcommand.
// Returns an error if command execution fails.
func Execute() error {
	rootCmd.SetOut(os.Stdout)
	return rootCmd.Execute()
}
