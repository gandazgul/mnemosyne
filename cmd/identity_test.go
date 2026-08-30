package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func executeRootForIdentityTest(t *testing.T, args ...string) string {
	t.Helper()

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs(args)

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("creating stdout pipe: %v", err)
	}
	os.Stdout = w

	execErr := rootCmd.Execute()
	resetHelpFlagsForIdentityTest(rootCmd)

	w.Close() //nolint:errcheck
	os.Stdout = oldStdout

	var stdout bytes.Buffer
	_, _ = io.Copy(&stdout, r)

	if execErr != nil {
		t.Fatalf("root command %v returned error: %v\nOutput:\n%s%s", args, execErr, stdout.String(), outBuf.String())
	}

	return stdout.String() + outBuf.String()
}

func resetHelpFlagsForIdentityTest(command *cobra.Command) {
	if command == nil {
		return
	}
	if command.Flags().Lookup("help") != nil {
		command.Flags().Set("help", "false") //nolint:errcheck
	}
	if command.Flags().Lookup("version") != nil {
		command.Flags().Set("version", "false") //nolint:errcheck
	}
	for _, child := range command.Commands() {
		resetHelpFlagsForIdentityTest(child)
	}
}

func assertMnemotecaIdentity(t *testing.T, surface, output string) {
	t.Helper()
	if !strings.Contains(output, "Mnemoteca") && !strings.Contains(output, "mnemoteca") && !strings.Contains(output, "MNEMOTECA_") {
		t.Fatalf("%s output does not identify Mnemoteca:\n%s", surface, output)
	}
	for _, legacy := range []string{"Mnemosyne", "mnemosyne", "MNEMOSYNE_"} {
		if strings.Contains(output, legacy) {
			t.Fatalf("%s output contains legacy identity %q:\n%s", surface, legacy, output)
		}
	}
}

func TestCommandIdentitySurfacesUseMnemotecaOnly(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{name: "no args", args: []string{}},
		{name: "root help", args: []string{"--help"}},
		{name: "root version flag", args: []string{"--version"}},
		{name: "version command", args: []string{"version"}},
		{name: "setup help", args: []string{"setup", "--help"}},
		{name: "cleanup help", args: []string{"cleanup", "--help"}},
		{name: "stats help", args: []string{"stats", "--help"}},
		{name: "add help", args: []string{"add", "--help"}},
		{name: "search help", args: []string{"search", "--help"}},
		{name: "export help", args: []string{"export", "--help"}},
		{name: "import help", args: []string{"import", "--help"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertMnemotecaIdentity(t, tc.name, executeRootForIdentityTest(t, tc.args...))
		})
	}
}

func TestRepresentativeErrorsUseMnemotecaGuidanceOnly(t *testing.T) {
	t.Setenv("MNEMOTECA_DB_PATH", filepath.Join(t.TempDir(), "mnemoteca.db"))

	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(outBuf)
	rootCmd.SetArgs([]string{"list", "--name", "missing"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected missing named collection error")
	}
	output := err.Error() + "\n" + outBuf.String()
	assertMnemotecaIdentity(t, "missing collection error", output)
	if !strings.Contains(output, "mnemoteca init --name missing") {
		t.Fatalf("expected Mnemoteca init guidance, got:\n%s", output)
	}
}
