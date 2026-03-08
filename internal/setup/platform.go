// Package setup handles downloading ONNX Runtime and ML models on first use.
//
// It downloads:
//   - ONNX Runtime shared library from GitHub releases
//   - snowflake-arctic-embed-m-v1.5 embedding model from HuggingFace
//   - ms-marco-MiniLM-L-6-v2 reranker model from HuggingFace
//
// Files are stored in ~/.local/share/mnemosyne/{lib,models}/.
package setup

import (
	"fmt"
	"runtime"
)

// ONNXRuntimeVersion is the pinned ONNX Runtime release.
const ONNXRuntimeVersion = "1.23.1"

// onnxRuntimeURL returns the GitHub release tarball URL for the current platform.
func onnxRuntimeURL() (string, error) {
	osName, err := ortOS()
	if err != nil {
		return "", err
	}
	archName, err := ortArch()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"https://github.com/microsoft/onnxruntime/releases/download/v%s/onnxruntime-%s-%s-%s.tgz",
		ONNXRuntimeVersion, osName, archName, ONNXRuntimeVersion,
	), nil
}

// ortOS maps runtime.GOOS to ONNX Runtime release naming.
func ortOS() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		return "osx", nil
	case "linux":
		return "linux", nil
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

// ortArch maps runtime.GOARCH to ONNX Runtime release naming.
func ortArch() (string, error) {
	switch runtime.GOARCH {
	case "arm64":
		return "arm64", nil
	case "amd64":
		return "x86_64", nil
	default:
		return "", fmt.Errorf("unsupported architecture: %s", runtime.GOARCH)
	}
}

// onnxRuntimeLibName returns the expected library filename for the current platform.
func onnxRuntimeLibName() string {
	switch runtime.GOOS {
	case "darwin":
		return "libonnxruntime.dylib"
	default:
		return "libonnxruntime.so"
	}
}

// hfFileURL returns a direct download URL for a file in a HuggingFace repo.
// No authentication needed for non-gated models.
func hfFileURL(repo, filePath string) string {
	return fmt.Sprintf("https://huggingface.co/%s/resolve/main/%s", repo, filePath)
}
