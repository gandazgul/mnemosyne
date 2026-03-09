package cmd

import (
	"fmt"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/gandazgul/mnemosyne/internal/setup"
	"github.com/spf13/cobra"
)

// setupCmd downloads ONNX Runtime and ML models.
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Download ONNX Runtime and ML models",
	Long: `Download the required ONNX Runtime shared library and ML models.

This downloads:
  - ONNX Runtime (~38 MB) from GitHub releases
  - snowflake-arctic-embed-m-v1.5 embedding model (~420 MB) from HuggingFace
  - ms-marco-MiniLM-L-6-v2 reranker model (~80 MB) from HuggingFace

Files are stored in ~/.local/share/mnemosyne/.
This command is idempotent — it skips files that are already downloaded.

Note: This happens automatically on first use of 'add' or 'search'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dataDir := config.DataDir()

		// Check current status.
		status := setup.Check(dataDir)
		if status.Ready() {
			fmt.Println("All components are already installed.")
			printStatus(status)
			return nil
		}

		printStatus(status)
		fmt.Println()

		// Run setup with simple progress output.
		var currentFile string
		err := setup.Run(cmd.Context(), dataDir, func(file string, written, total int64) {
			if file != currentFile {
				if currentFile != "" {
					fmt.Println(" done")
				}
				currentFile = file
				fmt.Printf("  Downloading %s...", file)
			}
		})
		if err != nil {
			return err
		}
		if currentFile != "" {
			fmt.Println(" done")
		}

		fmt.Println()
		fmt.Println("Setup complete!")
		printStatus(setup.Check(dataDir))
		return nil
	},
}

func printStatus(s setup.Status) {
	check := func(ok bool) string {
		if ok {
			return "installed"
		}
		return "missing"
	}
	fmt.Printf("  ONNX Runtime:     %s\n", check(s.OnnxRuntimeInstalled))
	fmt.Printf("  Embedding model:  %s\n", check(s.EmbeddingModelReady))
	fmt.Printf("  Reranker model:   %s\n", check(s.RerankerModelReady))
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
