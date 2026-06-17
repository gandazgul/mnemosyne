// Package setup handles downloading ONNX Runtime and ML models on first use.
//
// It downloads:
//   - ONNX Runtime shared library from GitHub releases
//   - snowflake-arctic-embed-m-v1.5 embedding model from HuggingFace
//   - ms-marco-MiniLM-L-6-v2 reranker model from HuggingFace
//
// Files are stored in the platform default data directory under {lib,models}/.
package setup

import (
	"fmt"
	"runtime"
)

// ONNXRuntimeVersion is the pinned ONNX Runtime release.
const ONNXRuntimeVersion = "1.23.1"

// onnxRuntimeURL returns the GitHub release archive URL for the current platform.
func onnxRuntimeURL() (string, error) {
	return onnxRuntimeURLFor(runtime.GOOS, runtime.GOARCH)
}

func onnxRuntimeURLFor(goos, goarch string) (string, error) {
	osName, err := ortOSFor(goos)
	if err != nil {
		return "", err
	}
	archName, err := ortArchFor(goos, goarch)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"https://github.com/microsoft/onnxruntime/releases/download/v%s/onnxruntime-%s-%s-%s%s",
		ONNXRuntimeVersion, osName, archName, ONNXRuntimeVersion, onnxRuntimeArchiveExt(goos),
	), nil
}

// ortOS maps runtime.GOOS to ONNX Runtime release naming.
func ortOS() (string, error) {
	return ortOSFor(runtime.GOOS)
}

func ortOSFor(goos string) (string, error) {
	switch goos {
	case "darwin":
		return "osx", nil
	case "linux":
		return "linux", nil
	case "windows":
		return "win", nil
	default:
		return "", fmt.Errorf("unsupported OS: %s", goos)
	}
}

// ortArch maps runtime.GOARCH to ONNX Runtime release naming.
func ortArch() (string, error) {
	return ortArchFor(runtime.GOOS, runtime.GOARCH)
}

func ortArchFor(goos, goarch string) (string, error) {
	switch goarch {
	case "arm64":
		if goos == "linux" {
			return "aarch64", nil
		}
		return "arm64", nil
	case "amd64":
		if goos == "linux" || goos == "windows" {
			return "x64", nil
		}
		return "x86_64", nil
	default:
		return "", fmt.Errorf("unsupported architecture: %s", goarch)
	}
}

// onnxRuntimeLibName returns the expected library filename for the current platform.
func onnxRuntimeLibName() string {
	return onnxRuntimeLibNameForOS(runtime.GOOS)
}

func onnxRuntimeLibNameForOS(goos string) string {
	switch goos {
	case "darwin":
		return "libonnxruntime.dylib"
	case "windows":
		return "onnxruntime.dll"
	default:
		return "libonnxruntime.so"
	}
}

func onnxRuntimeArchiveExt(goos string) string {
	if goos == "windows" {
		return ".zip"
	}
	return ".tgz"
}

// hfFileURL returns a direct download URL for a file in a HuggingFace repo.
// No authentication needed for non-gated models.
func hfFileURL(repo, filePath string) string {
	return fmt.Sprintf("https://huggingface.co/%s/resolve/main/%s", repo, filePath)
}
