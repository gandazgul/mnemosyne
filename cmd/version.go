package cmd

import (
	"github.com/spf13/cobra"
)

// These variables are set at build time via -ldflags. For example:
//
//	go build -ldflags "-X github.com/gandazgul/mnemosyne/cmd.Version=1.0.0"
//
// If not set, they default to "dev".
var (
	// Version is the semantic version of the build.
	Version = "dev"

	// Commit is the git commit hash of the build.
	Commit = "none"

	// Date is the build date.
	Date = "unknown"
)

// versionCmd prints the version information for mnemosyne.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of mnemosyne",
	Long:  "Display the version, git commit, and build date of this mnemosyne binary.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("mnemosyne %s\n", Version)
		cmd.Printf("  commit: %s\n", Commit)
		cmd.Printf("  built:  %s\n", Date)
	},
}

// init registers the version subcommand with the root command.
// In Go, init() functions run automatically when the package is loaded.
// Cobra uses this pattern to wire up the command tree before main() runs.
func init() {
	rootCmd.AddCommand(versionCmd)
}
