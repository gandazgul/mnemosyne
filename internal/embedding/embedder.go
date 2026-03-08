package embedding

import (
	"fmt"
	"math"
	"path/filepath"
	"sync"

	ort "github.com/yalue/onnxruntime_go"

	"github.com/gandazgul/mnemosyne/internal/config"
)

// Embedder generates vector embeddings from text.
type Embedder interface {
	// Embed generates an embedding for a single text (no prefix added).
	Embed(text string) ([]float32, error)

	// EmbedBatch generates embeddings for multiple texts (no prefix added).
	EmbedBatch(texts []string) ([][]float32, error)

	// EmbedQuery generates an embedding with the query prefix prepended.
	EmbedQuery(query string) ([]float32, error)

	// EmbedDocument generates an embedding with the document prefix prepended.
	EmbedDocument(doc string) ([]float32, error)

	// Dimensions returns the configured embedding dimensions.
	Dimensions() int

	// Close releases all resources (ONNX session, tokenizer).
	Close() error
}

// --- ONNX Runtime global lifecycle (shared across embedder + future reranker) ---

var (
	ortOnce    sync.Once
	ortInitErr error
)

// InitONNXRuntime initializes the ONNX Runtime environment.
// libraryPath is the path to the ONNX Runtime shared library (.so/.dylib/.dll).
// If empty, the default system path is used.
// Safe to call multiple times; initialization happens only once.
func InitONNXRuntime(libraryPath string) error {
	ortOnce.Do(func() {
		if libraryPath != "" {
			ort.SetSharedLibraryPath(libraryPath)
		}
		ortInitErr = ort.InitializeEnvironment()
	})
	return ortInitErr
}

// DestroyONNXRuntime cleans up the global ONNX Runtime environment.
// Call once when the application exits.
func DestroyONNXRuntime() error {
	return ort.DestroyEnvironment()
}

// --- ONNXEmbedder implementation ---

// ONNXEmbedder implements Embedder using an ONNX model via ONNX Runtime.
// It performs tokenization, inference, pooling (mean or CLS), and L2 normalization.
type ONNXEmbedder struct {
	session        *ort.DynamicAdvancedSession
	tokenizer      *Tokenizer
	dimensions     int
	queryPrefix    string
	documentPrefix string
	pooling        config.PoolingMethod
	hasTokenTypes  bool // true if session was created with token_type_ids input
}

// NewONNXEmbedder creates a new ONNX-based embedder from the given config.
// InitONNXRuntime must be called before this.
func NewONNXEmbedder(cfg config.EmbeddingConfig) (*ONNXEmbedder, error) {
	// Load tokenizer from model directory.
	tokenizer, err := NewTokenizer(cfg.ModelPath, cfg.MaxSeqLength)
	if err != nil {
		return nil, fmt.Errorf("load tokenizer: %w", err)
	}

	// Resolve ONNX model file path.
	modelFile := filepath.Join(cfg.ModelPath, cfg.OnnxFile)

	// Determine ONNX input names from config or defaults.
	inputNames := cfg.OnnxInputNames
	if len(inputNames) == 0 {
		inputNames = []string{"input_ids", "attention_mask"}
		if cfg.Pooling == config.PoolingCLS {
			inputNames = append(inputNames, "token_type_ids")
		}
	}

	// Check if token_type_ids is among the inputs.
	hasTokenTypes := false
	for _, name := range inputNames {
		if name == "token_type_ids" {
			hasTokenTypes = true
			break
		}
	}

	// Determine ONNX output names from config or defaults.
	outputNames := cfg.OnnxOutputNames
	if len(outputNames) == 0 {
		if cfg.Pooling == config.PoolingNone {
			outputNames = []string{"sentence_embedding"}
		} else {
			outputNames = []string{"last_hidden_state"}
		}
	}

	// Create a dynamic ONNX session (supports variable-length inputs).
	session, err := ort.NewDynamicAdvancedSession(
		modelFile,
		inputNames,
		outputNames,
		nil, // default session options
	)
	if err != nil {
		tokenizer.Close()
		return nil, fmt.Errorf("create ONNX session for %s: %w", modelFile, err)
	}

	pooling := cfg.Pooling
	if pooling == "" {
		pooling = config.PoolingMean // backward compat
	}

	return &ONNXEmbedder{
		session:        session,
		tokenizer:      tokenizer,
		dimensions:     cfg.Dimensions,
		queryPrefix:    cfg.QueryPrefix,
		documentPrefix: cfg.DocumentPrefix,
		pooling:        pooling,
		hasTokenTypes:  hasTokenTypes,
	}, nil
}

// Embed generates an embedding for a single text without any prefix.
func (e *ONNXEmbedder) Embed(text string) ([]float32, error) {
	results, err := e.embedTexts([]string{text})
	if err != nil {
		return nil, err
	}
	return results[0], nil
}

// EmbedBatch generates embeddings for multiple texts without any prefix.
func (e *ONNXEmbedder) EmbedBatch(texts []string) ([][]float32, error) {
	return e.embedTexts(texts)
}

// EmbedQuery generates an embedding with the configured query prefix.
func (e *ONNXEmbedder) EmbedQuery(query string) ([]float32, error) {
	return e.Embed(e.queryPrefix + query)
}

// EmbedDocument generates an embedding with the configured document prefix.
func (e *ONNXEmbedder) EmbedDocument(doc string) ([]float32, error) {
	return e.Embed(e.documentPrefix + doc)
}

// Dimensions returns the configured embedding dimensions.
func (e *ONNXEmbedder) Dimensions() int {
	return e.dimensions
}

// Close releases the ONNX session and tokenizer resources.
func (e *ONNXEmbedder) Close() error {
	var errs []error
	if e.session != nil {
		if err := e.session.Destroy(); err != nil {
			errs = append(errs, fmt.Errorf("destroy session: %w", err))
		}
	}
	if e.tokenizer != nil {
		e.tokenizer.Close()
	}
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

// embedTexts is the core implementation that handles tokenization, inference,
// pooling (mean or CLS), L2 normalization, and optional MRL dimension truncation.
func (e *ONNXEmbedder) embedTexts(texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, fmt.Errorf("empty text batch")
	}

	// Tokenize all texts with padding to the same length.
	encoded, maxLen, err := e.tokenizer.EncodeBatch(texts)
	if err != nil {
		return nil, fmt.Errorf("tokenize: %w", err)
	}

	batchSize := int64(len(texts))
	seqLen := int64(maxLen)

	// Flatten encoded inputs into contiguous arrays for ONNX tensors.
	flatIDs := make([]int64, batchSize*seqLen)
	flatMask := make([]int64, batchSize*seqLen)
	flatTokenTypes := make([]int64, batchSize*seqLen) // all zeros
	for i, enc := range encoded {
		copy(flatIDs[int64(i)*seqLen:], enc.InputIDs)
		copy(flatMask[int64(i)*seqLen:], enc.AttentionMask)
		copy(flatTokenTypes[int64(i)*seqLen:], enc.TokenTypeIDs)
	}

	// Create ONNX input tensors.
	inputShape := ort.NewShape(batchSize, seqLen)

	idsTensor, err := ort.NewTensor(inputShape, flatIDs)
	if err != nil {
		return nil, fmt.Errorf("create input_ids tensor: %w", err)
	}
	defer idsTensor.Destroy()

	maskTensor, err := ort.NewTensor(inputShape, flatMask)
	if err != nil {
		return nil, fmt.Errorf("create attention_mask tensor: %w", err)
	}
	defer maskTensor.Destroy()

	// Build input list: always input_ids + attention_mask, optionally token_type_ids.
	inputs := []ort.Value{idsTensor, maskTensor}
	if e.hasTokenTypes {
		tokenTypeTensor, err := ort.NewTensor(inputShape, flatTokenTypes)
		if err != nil {
			return nil, fmt.Errorf("create token_type_ids tensor: %w", err)
		}
		defer tokenTypeTensor.Destroy()
		inputs = append(inputs, tokenTypeTensor)
	}

	// Run inference. Output is auto-allocated by ONNX Runtime.
	outputs := []ort.Value{nil}
	err = e.session.Run(inputs, outputs)
	if err != nil {
		return nil, fmt.Errorf("ONNX inference: %w", err)
	}
	defer func() {
		if outputs[0] != nil {
			outputs[0].Destroy()
		}
	}()

	// Extract output data.
	outputShape := outputs[0].GetShape()

	// Type-assert to get the float32 data.
	outputTensor, ok := outputs[0].(*ort.Tensor[float32])
	if !ok {
		return nil, fmt.Errorf("unexpected output tensor type (expected float32)")
	}
	outputData := outputTensor.GetData()

	// Perform pooling and L2 normalization for each text in the batch.
	results := make([][]float32, batchSize)

	switch {
	case e.pooling == config.PoolingNone && len(outputShape) == 2:
		// Pre-pooled output: shape [batch_size, hidden_dim].
		hiddenDim := int(outputShape[1])
		for i := int64(0); i < batchSize; i++ {
			embedding := make([]float32, hiddenDim)
			copy(embedding, outputData[i*int64(hiddenDim):(i+1)*int64(hiddenDim)])
			l2Normalize(embedding)

			if e.dimensions < hiddenDim {
				embedding = embedding[:e.dimensions]
				l2Normalize(embedding)
			}
			results[i] = embedding
		}

	case len(outputShape) == 3:
		// Token-level output: shape [batch_size, seq_len, hidden_dim].
		hiddenDim := int(outputShape[2])
		for i := int64(0); i < batchSize; i++ {
			var embedding []float32
			switch e.pooling {
			case config.PoolingCLS:
				embedding = clsPool(outputData, i, seqLen, hiddenDim)
			default: // PoolingMean
				embedding = meanPool(outputData, flatMask, i, seqLen, hiddenDim)
			}
			l2Normalize(embedding)

			// MRL dimension truncation: if configured dimensions < hidden_dim,
			// truncate the embedding and re-normalize.
			if e.dimensions < hiddenDim {
				embedding = embedding[:e.dimensions]
				l2Normalize(embedding)
			}
			results[i] = embedding
		}

	default:
		return nil, fmt.Errorf("unexpected output shape: %v (pooling=%s)", outputShape, e.pooling)
	}

	return results, nil
}

// clsPool extracts the [CLS] token's hidden state (first token, index 0)
// as the sentence embedding. Used by BERT-family models like Snowflake Arctic.
func clsPool(hiddenStates []float32, batchIdx, seqLen int64, hiddenDim int) []float32 {
	pooled := make([]float32, hiddenDim)
	baseOffset := batchIdx * seqLen * int64(hiddenDim)
	copy(pooled, hiddenStates[baseOffset:baseOffset+int64(hiddenDim)])
	return pooled
}

// meanPool performs mean pooling over the sequence dimension, using the
// attention mask to exclude padding tokens.
//
// For each dimension d:
//
//	pooled[d] = sum(hidden_state[t][d] * mask[t]) / sum(mask[t])
func meanPool(hiddenStates []float32, mask []int64, batchIdx, seqLen int64, hiddenDim int) []float32 {
	pooled := make([]float32, hiddenDim)

	// Calculate sum of attention mask values (number of real tokens).
	var maskSum float32
	for t := int64(0); t < seqLen; t++ {
		maskSum += float32(mask[batchIdx*seqLen+t])
	}

	// Avoid division by zero.
	if maskSum == 0 {
		return pooled
	}

	// Accumulate weighted hidden states.
	baseOffset := batchIdx * seqLen * int64(hiddenDim)
	for t := int64(0); t < seqLen; t++ {
		m := float32(mask[batchIdx*seqLen+t])
		if m == 0 {
			continue
		}
		tokenOffset := baseOffset + t*int64(hiddenDim)
		for d := 0; d < hiddenDim; d++ {
			pooled[d] += hiddenStates[tokenOffset+int64(d)] * m
		}
	}

	// Divide by mask sum to get mean.
	for d := 0; d < hiddenDim; d++ {
		pooled[d] /= maskSum
	}

	return pooled
}

// l2Normalize normalizes a vector to unit length (L2 norm = 1).
// Modifies the slice in place. Does nothing if the vector is all zeros.
func l2Normalize(v []float32) {
	var sum float64
	for _, val := range v {
		sum += float64(val) * float64(val)
	}
	norm := math.Sqrt(sum)
	if norm == 0 {
		return
	}
	for i := range v {
		v[i] = float32(float64(v[i]) / norm)
	}
}
