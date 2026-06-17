package setup

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// EmbedModel describes a HuggingFace model to download.
type EmbedModel struct {
	// Repo is the HuggingFace repository (e.g., "Snowflake/snowflake-arctic-embed-m-v1.5").
	Repo string
	// LocalDir is the directory name under models/ (e.g., "snowflake-arctic-embed-m-v1.5").
	LocalDir string
	// Files maps relative file paths to their expected SHA-256 checksums.
	// An empty checksum means no verification (useful during development).
	Files map[string]string
}

// Models to download. Checksums will be populated after initial testing.
var (
	EmbeddingModel = EmbedModel{
		Repo:     "Snowflake/snowflake-arctic-embed-m-v1.5",
		LocalDir: "snowflake-arctic-embed-m-v1.5",
		Files: map[string]string{
			"onnx/model.onnx": "", // ~420 MB - checksum TBD
			"tokenizer.json":  "",
			"config.json":     "",
		},
	}

	RerankerModel = EmbedModel{
		Repo:     "cross-encoder/ms-marco-MiniLM-L-6-v2",
		LocalDir: "ms-marco-MiniLM-L-6-v2",
		Files: map[string]string{
			"onnx/model.onnx": "", // ~80 MB - checksum TBD
			"tokenizer.json":  "",
			"config.json":     "",
		},
	}
)

// ProgressFunc is called during download with the current file name,
// bytes written, and total bytes (-1 if unknown).
type ProgressFunc func(file string, written, total int64)

// Status describes the current state of required files.
type Status struct {
	OnnxRuntimeInstalled bool
	EmbeddingModelReady  bool
	RerankerModelReady   bool
}

// Check returns the current setup status for the given data directory.
func Check(dataDir string) Status {
	var s Status

	libPath := filepath.Join(dataDir, "lib", onnxRuntimeLibName())
	if _, err := os.Stat(libPath); err == nil {
		s.OnnxRuntimeInstalled = true
	}

	s.EmbeddingModelReady = modelReady(dataDir, EmbeddingModel)
	s.RerankerModelReady = modelReady(dataDir, RerankerModel)
	return s
}

// Ready returns true if all required components are installed.
func (s Status) Ready() bool {
	return s.OnnxRuntimeInstalled && s.EmbeddingModelReady && s.RerankerModelReady
}

// modelReady checks if all files for a model exist in the data directory.
func modelReady(dataDir string, model EmbedModel) bool {
	modelDir := filepath.Join(dataDir, "models", model.LocalDir)
	for file := range model.Files {
		if _, err := os.Stat(filepath.Join(modelDir, file)); err != nil {
			return false
		}
	}
	return true
}

// Run downloads all required components to dataDir.
// progress is called for each file being downloaded.
func Run(ctx context.Context, dataDir string, progress ProgressFunc) error {
	status := Check(dataDir)

	if !status.OnnxRuntimeInstalled {
		if err := installONNXRuntime(ctx, dataDir, progress); err != nil {
			return fmt.Errorf("install ONNX Runtime: %w", err)
		}
	}

	if !status.EmbeddingModelReady {
		if err := downloadModel(ctx, dataDir, EmbeddingModel, progress); err != nil {
			return fmt.Errorf("download embedding model: %w", err)
		}
	}

	if !status.RerankerModelReady {
		if err := downloadModel(ctx, dataDir, RerankerModel, progress); err != nil {
			return fmt.Errorf("download reranker model: %w", err)
		}
	}

	return nil
}

// EnsureReady checks if setup is complete and runs it if not.
// This is the auto-download entry point called by add/search commands.
func EnsureReady(ctx context.Context, dataDir string, progress ProgressFunc) error {
	status := Check(dataDir)
	if status.Ready() {
		return nil
	}

	fmt.Println("First-time setup: downloading ONNX Runtime and ML models...")
	fmt.Println("This is a one-time download (~500 MB total).")
	fmt.Println()

	bar := NewProgressBar(os.Stdout)
	err := Run(ctx, dataDir, bar.Update)
	bar.Finish()
	return err
}

// installONNXRuntime downloads and extracts the ONNX Runtime shared library.
func installONNXRuntime(ctx context.Context, dataDir string, progress ProgressFunc) error {
	url, err := onnxRuntimeURL()
	if err != nil {
		return err
	}

	libDir := filepath.Join(dataDir, "lib")
	if err := os.MkdirAll(libDir, 0o755); err != nil {
		return fmt.Errorf("create lib dir: %w", err)
	}

	// Download archive to a temp file, then extract the library.
	archiveName := "onnxruntime" + onnxRuntimeArchiveExt(runtime.GOOS)
	archivePath := filepath.Join(libDir, archiveName)
	if err := downloadFile(ctx, url, archivePath, "", func(written, total int64) {
		if progress != nil {
			progress(archiveName, written, total)
		}
	}); err != nil {
		return err
	}
	defer os.Remove(archivePath) //nolint:errcheck

	// Extract runtime library files from the archive.
	if err := extractONNXRuntimeLib(archivePath, libDir); err != nil {
		return fmt.Errorf("extract: %w", err)
	}

	return nil
}

// extractONNXRuntimeLib extracts runtime library files from the ONNX Runtime archive.
func extractONNXRuntimeLib(archivePath, destDir string) error {
	if strings.HasSuffix(archivePath, ".zip") {
		return extractONNXRuntimeZipLib(archivePath, destDir)
	}
	return extractONNXRuntimeTarLib(archivePath, destDir)
}

// extractONNXRuntimeTarLib extracts libonnxruntime* files from an ONNX Runtime tarball.
func extractONNXRuntimeTarLib(tgzPath, destDir string) error {
	f, err := os.Open(tgzPath)
	if err != nil {
		return err
	}
	defer f.Close() //nolint:errcheck

	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gz.Close() //nolint:errcheck

	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Only extract lib/libonnxruntime* files.
		base := filepath.Base(hdr.Name)
		if !strings.HasPrefix(base, "libonnxruntime") {
			continue
		}
		if hdr.Typeflag == tar.TypeDir {
			continue
		}

		destPath := filepath.Join(destDir, base)

		// Handle symlinks.
		if hdr.Typeflag == tar.TypeSymlink {
			_ = os.Remove(destPath) // remove existing symlink if any
			if err := os.Symlink(hdr.Linkname, destPath); err != nil {
				return fmt.Errorf("create symlink %s -> %s: %w", base, hdr.Linkname, err)
			}
			continue
		}

		out, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			return err
		}
		if _, err := io.Copy(out, tr); err != nil {
			_ = out.Close()
			return err
		}
		if err := out.Close(); err != nil {
			return err
		}
	}

	return nil
}

// extractONNXRuntimeZipLib extracts onnxruntime.dll from a Windows ONNX Runtime zip archive.
func extractONNXRuntimeZipLib(zipPath, destDir string) error {
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zr.Close() //nolint:errcheck

	for _, file := range zr.File {
		if file.FileInfo().IsDir() {
			continue
		}

		base := filepath.Base(file.Name)
		if base != onnxRuntimeLibNameForOS("windows") {
			continue
		}

		in, err := file.Open()
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, base)
		out, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			_ = in.Close()
			return err
		}
		if _, err := io.Copy(out, in); err != nil {
			_ = in.Close()
			_ = out.Close()
			return err
		}
		if err := in.Close(); err != nil {
			_ = out.Close()
			return err
		}
		if err := out.Close(); err != nil {
			return err
		}
	}

	return nil
}

// downloadModel downloads all files for a HuggingFace model.
func downloadModel(ctx context.Context, dataDir string, model EmbedModel, progress ProgressFunc) error {
	modelDir := filepath.Join(dataDir, "models", model.LocalDir)

	for file, sha := range model.Files {
		destPath := filepath.Join(modelDir, file)
		url := hfFileURL(model.Repo, file)

		if err := downloadFile(ctx, url, destPath, sha, func(written, total int64) {
			if progress != nil {
				progress(model.LocalDir+"/"+file, written, total)
			}
		}); err != nil {
			return fmt.Errorf("download %s/%s: %w", model.LocalDir, file, err)
		}
	}

	return nil
}
