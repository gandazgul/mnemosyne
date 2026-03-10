// Package embedding provides text-to-vector embedding using ONNX Runtime.
//
// It wraps a HuggingFace tokenizer (via daulet/tokenizers) and an ONNX model
// session (via yalue/onnxruntime_go) to generate dense vector embeddings
// from text. The primary use case is encoding documents and queries for
// semantic vector search.
package embedding

import (
	"fmt"
	"path/filepath"

	"github.com/daulet/tokenizers"
)

// Tokenizer wraps a HuggingFace tokenizer loaded from a tokenizer.json file.
// It handles text encoding (tokenization) and provides input tensors suitable
// for ONNX model inference: input_ids and attention_mask.
type Tokenizer struct {
	inner     *tokenizers.Tokenizer
	maxSeqLen int
}

// NewTokenizer loads a HuggingFace tokenizer from the given model directory.
// It expects a tokenizer.json file in the directory. maxSeqLen limits the
// maximum number of tokens per sequence (truncating longer inputs).
func NewTokenizer(modelDir string, maxSeqLen int) (*Tokenizer, error) {
	tokenizerPath := filepath.Join(modelDir, "tokenizer.json")

	tk, err := tokenizers.FromFile(tokenizerPath)
	if err != nil {
		return nil, fmt.Errorf("load tokenizer from %s: %w", tokenizerPath, err)
	}

	return &Tokenizer{
		inner:     tk,
		maxSeqLen: maxSeqLen,
	}, nil
}

// EncodedInput holds the tokenized output for a single text, ready for
// ONNX model input.
type EncodedInput struct {
	InputIDs      []int64
	AttentionMask []int64
	TokenTypeIDs  []int64 // all zeros for single-sentence encoding (BERT models)
	Length        int     // actual token count (before padding)
}

// Encode tokenizes a single text string and returns input_ids and
// attention_mask as int64 slices. Special tokens (BOS/EOS) are added.
// Sequences longer than maxSeqLen are truncated.
func (t *Tokenizer) Encode(text string) (*EncodedInput, error) {
	opts := []tokenizers.EncodeOption{
		tokenizers.WithReturnAttentionMask(),
	}

	result := t.inner.EncodeWithOptions(text, true, opts...)
	if result.IDs == nil {
		return nil, fmt.Errorf("tokenization returned nil IDs for text: %.50s", text)
	}

	ids := result.IDs
	mask := result.AttentionMask

	// Truncate to maxSeqLen if needed.
	if len(ids) > t.maxSeqLen {
		ids = ids[:t.maxSeqLen]
		mask = mask[:t.maxSeqLen]
	}

	tokenLen := len(ids)

	// Convert uint32 -> int64 for ONNX Runtime.
	inputIDs := make([]int64, tokenLen)
	attentionMask := make([]int64, tokenLen)
	tokenTypeIDs := make([]int64, tokenLen) // all zeros for single-sentence
	for i := 0; i < tokenLen; i++ {
		inputIDs[i] = int64(ids[i])
		attentionMask[i] = int64(mask[i])
	}

	return &EncodedInput{
		InputIDs:      inputIDs,
		AttentionMask: attentionMask,
		TokenTypeIDs:  tokenTypeIDs,
		Length:        tokenLen,
	}, nil
}

// EncodeBatch tokenizes multiple texts and pads them to the same length.
// Returns a slice of EncodedInputs and the maximum sequence length in the batch.
// All sequences are padded to maxLen with 0s for both input_ids and attention_mask.
func (t *Tokenizer) EncodeBatch(texts []string) ([]*EncodedInput, int, error) {
	if len(texts) == 0 {
		return nil, 0, fmt.Errorf("empty text batch")
	}

	// Encode each text individually.
	encoded := make([]*EncodedInput, len(texts))
	maxLen := 0
	for i, text := range texts {
		enc, err := t.Encode(text)
		if err != nil {
			return nil, 0, fmt.Errorf("encode text %d: %w", i, err)
		}
		encoded[i] = enc
		if enc.Length > maxLen {
			maxLen = enc.Length
		}
	}

	// Pad all sequences to maxLen.
	for _, enc := range encoded {
		if enc.Length < maxLen {
			padLen := maxLen - enc.Length
			enc.InputIDs = append(enc.InputIDs, make([]int64, padLen)...)
			enc.AttentionMask = append(enc.AttentionMask, make([]int64, padLen)...)
			enc.TokenTypeIDs = append(enc.TokenTypeIDs, make([]int64, padLen)...)
		}
	}

	return encoded, maxLen, nil
}

// Close releases the tokenizer's native resources.
func (t *Tokenizer) Close() {
	if t.inner != nil {
		t.inner.Close() //nolint:errcheck
	}
}
