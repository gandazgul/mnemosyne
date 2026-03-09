package setup

import "testing"

func TestPlatformFunctions(t *testing.T) {
	_, _ = onnxRuntimeURL()
	_, _ = ortOS()
	_, _ = ortArch()
	_ = hfFileURL("repo", "file.txt")
}
