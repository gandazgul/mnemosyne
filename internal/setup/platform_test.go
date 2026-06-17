package setup

import (
	"strings"
	"testing"
)

func TestPlatformFunctions(t *testing.T) {
	_, _ = onnxRuntimeURL()
	_, _ = ortOS()
	_, _ = ortArch()
	_ = hfFileURL("repo", "file.txt")
}

func TestONNXRuntimeURLForWindows(t *testing.T) {
	tests := []struct {
		name     string
		goarch   string
		contains string
	}{
		{
			name:     "amd64",
			goarch:   "amd64",
			contains: "onnxruntime-win-x64-" + ONNXRuntimeVersion + ".zip",
		},
		{
			name:     "arm64",
			goarch:   "arm64",
			contains: "onnxruntime-win-arm64-" + ONNXRuntimeVersion + ".zip",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := onnxRuntimeURLFor("windows", tt.goarch)
			if err != nil {
				t.Fatalf("onnxRuntimeURLFor: %v", err)
			}
			if !strings.Contains(url, tt.contains) {
				t.Fatalf("url %q does not contain %q", url, tt.contains)
			}
		})
	}
}

func TestONNXRuntimePlatformMappings(t *testing.T) {
	tests := []struct {
		goos    string
		goarch  string
		ortOS   string
		ortArch string
		libName string
		ext     string
	}{
		{"darwin", "amd64", "osx", "x86_64", "libonnxruntime.dylib", ".tgz"},
		{"darwin", "arm64", "osx", "arm64", "libonnxruntime.dylib", ".tgz"},
		{"linux", "amd64", "linux", "x64", "libonnxruntime.so", ".tgz"},
		{"linux", "arm64", "linux", "aarch64", "libonnxruntime.so", ".tgz"},
		{"windows", "amd64", "win", "x64", "onnxruntime.dll", ".zip"},
		{"windows", "arm64", "win", "arm64", "onnxruntime.dll", ".zip"},
	}

	for _, tt := range tests {
		t.Run(tt.goos+"/"+tt.goarch, func(t *testing.T) {
			gotOS, err := ortOSFor(tt.goos)
			if err != nil {
				t.Fatalf("ortOSFor: %v", err)
			}
			if gotOS != tt.ortOS {
				t.Fatalf("ortOSFor = %q, want %q", gotOS, tt.ortOS)
			}

			gotArch, err := ortArchFor(tt.goos, tt.goarch)
			if err != nil {
				t.Fatalf("ortArchFor: %v", err)
			}
			if gotArch != tt.ortArch {
				t.Fatalf("ortArchFor = %q, want %q", gotArch, tt.ortArch)
			}

			if got := onnxRuntimeLibNameForOS(tt.goos); got != tt.libName {
				t.Fatalf("onnxRuntimeLibNameForOS = %q, want %q", got, tt.libName)
			}
			if got := onnxRuntimeArchiveExt(tt.goos); got != tt.ext {
				t.Fatalf("onnxRuntimeArchiveExt = %q, want %q", got, tt.ext)
			}
		})
	}
}
