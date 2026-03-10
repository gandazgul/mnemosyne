package setup

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// downloadFile downloads a URL to destPath. If the file already exists and
// matches expectedSHA256 (if non-empty), the download is skipped.
//
// Supports resuming partial downloads via HTTP Range requests. Uses a .partial
// suffix during download and renames on completion.
//
// progress is called with (bytesWritten, totalBytes) during download.
// totalBytes may be -1 if the server doesn't provide Content-Length.
func downloadFile(ctx context.Context, url, destPath, expectedSHA256 string, progress func(written, total int64)) error {
	// If destination exists, check checksum and skip if valid.
	if expectedSHA256 != "" {
		if match, _ := fileMatchesSHA256(destPath, expectedSHA256); match {
			return nil // already downloaded
		}
	}

	// Ensure parent directory exists.
	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	partialPath := destPath + ".partial"

	// Check if a partial download exists for resume.
	var existingSize int64
	if info, err := os.Stat(partialPath); err == nil {
		existingSize = info.Size()
	}

	// Build HTTP request with optional Range header.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	if existingSize > 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", existingSize))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck

	// Handle response status.
	switch resp.StatusCode {
	case http.StatusOK:
		// Full response — start fresh.
		existingSize = 0
	case http.StatusPartialContent:
		// Resume successful.
	case http.StatusRequestedRangeNotSatisfiable:
		// File is already complete (or server doesn't support Range).
		// Re-download from scratch.
		resp.Body.Close() //nolint:errcheck
		return downloadFile(ctx, url, destPath, expectedSHA256, progress)
	default:
		return fmt.Errorf("HTTP %d from %s", resp.StatusCode, url)
	}

	// Calculate total size.
	totalSize := resp.ContentLength
	if totalSize > 0 && existingSize > 0 {
		totalSize += existingSize
	}

	// Open partial file for append (or create).
	flags := os.O_CREATE | os.O_WRONLY
	if existingSize > 0 && resp.StatusCode == http.StatusPartialContent {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
		existingSize = 0
	}

	f, err := os.OpenFile(partialPath, flags, 0o644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	// Copy with progress tracking.
	written := existingSize
	buf := make([]byte, 32*1024)
	for {
		n, readErr := resp.Body.Read(buf)
		if n > 0 {
			if _, writeErr := f.Write(buf[:n]); writeErr != nil {
				_ = f.Close()
				return fmt.Errorf("write file: %w", writeErr)
			}
			written += int64(n)
			if progress != nil {
				progress(written, totalSize)
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			_ = f.Close()
			return fmt.Errorf("read response: %w", readErr)
		}
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}

	// Verify checksum if provided.
	if expectedSHA256 != "" {
		match, err := fileMatchesSHA256(partialPath, expectedSHA256)
		if err != nil {
			return fmt.Errorf("verify checksum: %w", err)
		}
		if !match {
			os.Remove(partialPath) //nolint:errcheck
			return fmt.Errorf("checksum mismatch for %s", filepath.Base(destPath))
		}
	}

	// Atomic rename.
	if err := os.Rename(partialPath, destPath); err != nil {
		return fmt.Errorf("rename: %w", err)
	}

	return nil
}

// fileMatchesSHA256 checks if a file's SHA-256 matches the expected hex digest.
func fileMatchesSHA256(path, expected string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close() //nolint:errcheck

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	actual := hex.EncodeToString(h.Sum(nil))
	return strings.EqualFold(actual, expected), nil
}
