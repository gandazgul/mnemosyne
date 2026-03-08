package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Output format constants.
const (
	formatColor = "color"
	formatPlain = "plain"
)

// validFormats lists the accepted --format values.
var validFormats = []string{formatColor, formatPlain}

// validateFormat returns an error if the given format string is not recognized.
func validateFormat(f string) error {
	for _, v := range validFormats {
		if f == v {
			return nil
		}
	}
	return fmt.Errorf("invalid format %q; must be one of: %s", f, strings.Join(validFormats, ", "))
}

// Color helpers — return plain strings when format is "plain" or NO_COLOR is set.
var (
	boldCyan   = color.New(color.Bold, color.FgCyan).SprintFunc()
	boldWhite  = color.New(color.Bold, color.FgWhite).SprintFunc()
	dimWhite   = color.New(color.Faint).SprintFunc()
	green      = color.New(color.FgGreen).SprintFunc()
	yellow     = color.New(color.FgYellow).SprintFunc()
	boldYellow = color.New(color.Bold, color.FgYellow).SprintFunc()
)

// plain returns true if the format flag disables color.
func plain(format string) bool {
	return format == formatPlain
}

// colorize returns the styled string, or the raw value if format is "plain".
func colorize(format string, styler func(a ...interface{}) string, v interface{}) string {
	if plain(format) {
		return fmt.Sprint(v)
	}
	return styler(v)
}
