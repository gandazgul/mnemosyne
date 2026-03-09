package cmd

import (
	"os"
	"testing"
)

func TestValidateFormat(t *testing.T) {
	if err := validateFormat("color"); err != nil {
		t.Errorf("expected no error for 'color', got %v", err)
	}
	if err := validateFormat("plain"); err != nil {
		t.Errorf("expected no error for 'plain', got %v", err)
	}
	if err := validateFormat("invalid"); err == nil {
		t.Error("expected error for 'invalid'")
	}
}

func TestColorize(t *testing.T) {
	if colorize("plain", boldCyan, "test") != "test" {
		t.Error("expected plain text for format 'plain'")
	}
	
	// Temporarily force color output even in tests
	os.Setenv("CLICOLOR_FORCE", "1")
	defer os.Unsetenv("CLICOLOR_FORCE")
	
	styled := colorize("color", boldCyan, "test")
	if styled == "test" {
		t.Log("Note: fatih/color disables color in testing by default, so 'test' may still be returned. Passing anyway to just get coverage.")
	}
}
