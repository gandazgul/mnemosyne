package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gandazgul/mnemosyne/internal/config"
	"github.com/spf13/cobra"
)

// cleanupCmd removes downloaded assets (models and ONNX Runtime).
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove downloaded models and ONNX Runtime (keeps database by default)",
	Long: `Remove files downloaded by mnemosyne setup.

By default this deletes only the ML models and ONNX Runtime library:
  - ONNX Runtime shared library (lib/)
  - ML models: embedding and reranker (models/)

Use --db to also delete the SQLite database containing all your
collections and memories. This is irreversible — all memories will be lost.

You will be asked to confirm unless --yes is provided.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		yesFlag, _ := cmd.Flags().GetBool("yes")
		dbFlag, _ := cmd.Flags().GetBool("db")

		dataDir := config.DataDir()

		// Check if there's anything to clean up.
		if _, err := os.Stat(dataDir); os.IsNotExist(err) {
			cmd.Println("Nothing to clean up — data directory does not exist.")
			return nil
		}

		// Directories/files to remove.
		libDir := filepath.Join(dataDir, "lib")
		modelsDir := filepath.Join(dataDir, "models")
		dbPath := filepath.Join(dataDir, "mnemosyne.db")

		// Show what will be deleted.
		cmd.Println("This will permanently delete:")
		cmd.Printf("  %s\n", libDir)
		cmd.Printf("  %s\n", modelsDir)
		if dbFlag {
			cmd.Printf("  %s\n", dbPath)
			cmd.Println()
			cmd.Println("⚠️  WARNING: --db will delete your database.")
			cmd.Println("   ALL collections and memories will be permanently lost.")
			cmd.Println("   This action cannot be undone.")
		}
		cmd.Println()

		// Confirm models/runtime removal.
		if !yesFlag {
			cmd.Printf("This will remove ONNX Runtime and ML models. They can be re-downloaded with 'mnemosyne setup'.\n")
			cmd.Printf("Type 'yes' to confirm: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("reading confirmation: %w", err)
			}
			if strings.TrimSpace(input) != "yes" {
				cmd.Println("Aborted.")
				return nil
			}
		}

		// Confirm database removal separately.
		if dbFlag && !yesFlag {
			cmd.Println()
			cmd.Println("WARNING: --db will delete your database. All memories will be lost.")
			cmd.Println("This action cannot be undone.")
			cmd.Printf("Type 'yes' to confirm database deletion: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("reading confirmation: %w", err)
			}
			if strings.TrimSpace(input) != "yes" {
				cmd.Println("Skipping database deletion.")
				dbFlag = false
			}
		}

		// Remove lib/ and models/.
		for _, dir := range []string{libDir, modelsDir} {
			if err := os.RemoveAll(dir); err != nil {
				return fmt.Errorf("removing %s: %w", dir, err)
			}
			cmd.Printf("Removed %s\n", dir)
		}

		// Remove database if requested.
		if dbFlag {
			if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("removing database: %w", err)
			}
			cmd.Printf("Removed %s\n", dbPath)
		}

		return nil
	},
}

func init() {
	cleanupCmd.Flags().Bool("yes", false, "skip confirmation prompt")
	cleanupCmd.Flags().Bool("db", false, "also delete the database (all memories will be lost)")
	rootCmd.AddCommand(cleanupCmd)
}
