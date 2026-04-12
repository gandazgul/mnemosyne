package setup

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mattn/go-isatty"
)

// ProgressBar renders a terminal progress bar for file downloads.
type ProgressBar struct {
	mu          sync.Mutex
	w           io.Writer
	currentFile string
	isTTY       bool
	lastUpdate  int64 // unix millis of last render
}

// NewProgressBar creates a progress bar that writes to w.
func NewProgressBar(w io.Writer) *ProgressBar {
	isTTY := false
	if f, ok := w.(*os.File); ok {
		isTTY = isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
	}
	return &ProgressBar{w: w, isTTY: isTTY}
}

// Update is a ProgressFunc-compatible callback for use with setup.Run / EnsureReady.
func (p *ProgressBar) Update(file string, written, total int64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now().UnixMilli()
	newFile := file != p.currentFile

	if newFile {
		// Finish the previous line.
		if p.currentFile != "" {
			fmt.Fprintln(p.w) //nolint:errcheck
		}
		p.currentFile = file
		p.lastUpdate = 0 // force render on new file
	}

	// Rate-limit updates to avoid flooding: render at most every 100ms,
	// but always render the first and last update for each file.
	isComplete := total > 0 && written >= total
	if !newFile && !isComplete && now-p.lastUpdate < 100 {
		return
	}
	p.lastUpdate = now

	if total <= 0 {
		// Unknown total — show bytes only.
		line := fmt.Sprintf("  Downloading %s... %s", file, formatBytes(written))
		p.writeLine(line)
		return
	}

	pct := float64(written) / float64(total) * 100
	if pct > 100 {
		pct = 100
	}

	barWidth := 30
	filled := int(float64(barWidth) * pct / 100)
	if filled > barWidth {
		filled = barWidth
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	line := fmt.Sprintf("  Downloading %s  %s / %s [%s] %.1f%%",
		file, formatBytes(written), formatBytes(total), bar, pct)
	p.writeLine(line)
}

// Finish prints a final newline if needed.
func (p *ProgressBar) Finish() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.currentFile != "" {
		fmt.Fprintln(p.w) //nolint:errcheck
		p.currentFile = ""
	}
}

// writeLine overwrites the current line on TTY, or prints a new line otherwise.
func (p *ProgressBar) writeLine(line string) {
	if p.isTTY {
		// Carriage return + clear to end of line, then write.
		fmt.Fprintf(p.w, "\r\033[K%s", line) //nolint:errcheck
	} else {
		// Non-TTY (piped output): just print with newline.
		fmt.Fprintln(p.w, line) //nolint:errcheck
	}
}

// formatBytes formats a byte count as a human-readable string.
func formatBytes(b int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case b >= GB:
		return fmt.Sprintf("%.1f GB", float64(b)/float64(GB))
	case b >= MB:
		return fmt.Sprintf("%.1f MB", float64(b)/float64(MB))
	case b >= KB:
		return fmt.Sprintf("%.1f KB", float64(b)/float64(KB))
	default:
		return fmt.Sprintf("%d B", b)
	}
}
