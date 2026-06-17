package cmd

import (
	"fmt"
	"io"
	"runtime"

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

	// Release is set to "true" for goreleaser CI builds, "false" otherwise.
	Release = "false"
)

// printVersion writes the version information to w.
// For CI/release builds (Release="true") it prints a single-line
// "mnemosyne vVERSION (platform)". For local builds it prints the
// full multi-line output with commit, build date, and platform.
func printVersion(w io.Writer) {
	platform := runtime.GOOS + "/" + runtime.GOARCH
	if Release == "true" {
		_, _ = fmt.Fprintf(w, "mnemosyne %s (%s)\n", Version, platform)
		return
	}
	_, _ = fmt.Fprintf(w, "mnemosyne %s (%s)\n", Version, platform)
	_, _ = fmt.Fprintf(w, "  commit: %s\n", Commit)
	_, _ = fmt.Fprintf(w, "  built:  %s\n", Date)
}

// versionCmd prints the version information for mnemosyne.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of mnemosyne",
	Long:  "Display the version, git commit, and build date of this mnemosyne binary.",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion(cmd.OutOrStdout())
	},
}

// init registers the version subcommand with the root command.
// In Go, init() functions run automatically when the package is loaded.
// Cobra uses this pattern to wire up the command tree before main() runs.
func init() {
	rootCmd.AddCommand(versionCmd)
}
