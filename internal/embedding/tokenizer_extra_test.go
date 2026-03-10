package embedding

import (
	"testing"
)

func TestTokenizer_EncodeBatch_Empty(t *testing.T) {
	tk := &Tokenizer{}
	_, _, err := tk.EncodeBatch([]string{})
	if err == nil {
		t.Error("expected error for empty batch")
	}
}

func TestEmbedder_EmptyTexts(t *testing.T) {
	e := &ONNXEmbedder{}
	_, err := e.embedTexts([]string{})
	if err == nil {
		t.Error("expected error for empty texts")
	}
}

func TestEmbedder_CloseErrors(t *testing.T) {
	e := &ONNXEmbedder{}
	err := e.Close()
	if err != nil {
		t.Errorf("expected no error closing uninitialized embedder, got %v", err)
	}
}

func TestDestroyONNXRuntime(t *testing.T) {
	_ = InitONNXRuntime("")
	err := DestroyONNXRuntime()
	if err != nil {
		t.Logf("DestroyONNXRuntime returned: %v", err)
	}
}

func TestTokenizer_FromFile_Error(t *testing.T) {
	_, err := NewTokenizer("/path/that/does/not/exist", 512)
	if err == nil {
		t.Error("expected error loading tokenizer from non-existent path")
	}
}

func TestEmbedder_Dimensions(t *testing.T) {
	e := &ONNXEmbedder{dimensions: 123}
	if e.Dimensions() != 123 {
		t.Errorf("expected 123, got %d", e.Dimensions())
	}
}
