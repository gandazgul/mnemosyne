package cmd

import (
	"fmt"

	"github.com/gandazgul/mnemoteca/internal/config"
	"github.com/gandazgul/mnemoteca/internal/setup"
	"github.com/spf13/cobra"
)

// setupCmd downloads ONNX Runtime and ML models.
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Download ONNX Runtime and ML models",
	Long: `Download the required ONNX Runtime shared library and ML models.

This downloads:
  - ONNX Runtime (~38 MB) from GitHub releases
  - default embedding/reranker models when the config points at built-in paths

Files are stored in ~/.local/share/mnemoteca/.
This command is idempotent — it skips files that are already downloaded.
Custom model paths are validated but not auto-downloaded.

Note: This happens automatically on first use of 'add' or 'search'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dataDir := config.DataDir()
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		// Check current status.
		status := setup.Check(dataDir, cfg)
		if status.Ready() {
			fmt.Println("All components are already installed.")
			printStatus(status)
			return nil
		}

		printStatus(status)
		fmt.Println()

		// Run setup with progress bar.
		progress := setup.NewProgressBar(cmd.OutOrStdout())
		err = setup.Run(cmd.Context(), dataDir, cfg, progress.Update)
		progress.Finish()
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println("Setup complete!")
		printStatus(setup.Check(dataDir, cfg))
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
	fmt.Printf("  Embedding model:  %s%s\n", check(s.EmbeddingModelReady), autoInstallSuffix(s.EmbeddingAutoInstall))
	if !s.RerankerEnabled {
		fmt.Println("  Reranker model:   disabled")
		return
	}
	fmt.Printf("  Reranker model:   %s%s\n", check(s.RerankerModelReady), autoInstallSuffix(s.RerankerAutoInstall))
}

func autoInstallSuffix(autoInstall bool) string {
	if autoInstall {
		return " (default auto-install)"
	}
	return " (custom path)"
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
