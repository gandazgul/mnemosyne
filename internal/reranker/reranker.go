package reranker

import (
	"fmt"
	"path/filepath"

	"github.com/daulet/tokenizers"
	ort "github.com/yalue/onnxruntime_go"

	"github.com/gandazgul/mnemosyne/internal/config"
)

// Reranker scores query-document relevance.
type Reranker interface {
	// Score takes a query and a list of documents, and returns a relevance score
	// for each document. Higher score = more relevant.
	Score(query string, documents []string) ([]float32, error)

	// Close releases all resources (ONNX session, tokenizer).
	Close() error
}

// ONNXCrossEncoder implements Reranker using an ONNX cross-encoder model.
type ONNXCrossEncoder struct {
	session      *ort.DynamicAdvancedSession
	tokenizer    *tokenizers.Tokenizer
	maxSeqLength int
}

// NewONNXCrossEncoder creates a new ONNX-based reranker.
// embedding.InitONNXRuntime must be called before this!
// (Wait, `InitONNXRuntime` is in the `embedding` package, we should just assume it's initialized globally).
func NewONNXCrossEncoder(cfg config.RerankerConfig) (*ONNXCrossEncoder, error) {
	if !cfg.Enabled {
		return nil, fmt.Errorf("reranker is disabled in config")
	}

	tokenizerPath := filepath.Join(cfg.ModelPath, "tokenizer.json")
	tk, err := tokenizers.FromFile(tokenizerPath)
	if err != nil {
		return nil, fmt.Errorf("load tokenizer from %s: %w", tokenizerPath, err)
	}

	modelFile := filepath.Join(cfg.ModelPath, "onnx", "model.onnx")

	// Cross-encoders usually take input_ids, attention_mask, token_type_ids
	inputNames := []string{"input_ids", "attention_mask", "token_type_ids"}
	// Typical output for cross-encoders
	outputNames := []string{"logits"}

	session, err := ort.NewDynamicAdvancedSession(
		modelFile,
		inputNames,
		outputNames,
		nil,
	)
	if err != nil {
		tk.Close()
		return nil, fmt.Errorf("create ONNX session for %s: %w", modelFile, err)
	}

	return &ONNXCrossEncoder{
		session:      session,
		tokenizer:    tk,
		maxSeqLength: cfg.MaxSeqLength,
	}, nil
}

// Score scores a single query against a batch of documents.
func (e *ONNXCrossEncoder) Score(query string, documents []string) ([]float32, error) {
	if len(documents) == 0 {
		return nil, nil
	}

	opts := []tokenizers.EncodeOption{
		tokenizers.WithReturnAttentionMask(),
		tokenizers.WithReturnTypeIDs(),
	}

	// 1. Encode query once
	encQ := e.tokenizer.EncodeWithOptions(query, true, opts...)
	if encQ.IDs == nil {
		return nil, fmt.Errorf("failed to tokenize query")
	}

	batchSize := int64(len(documents))

	// Pre-encode all documents to find the max length
	docIDs := make([][]uint32, len(documents))
	maxLen := 0

	for i, doc := range documents {
		encD := e.tokenizer.EncodeWithOptions(doc, true, opts...)
		if encD.IDs == nil {
			return nil, fmt.Errorf("failed to tokenize document %d", i)
		}

		// The encoded document includes [CLS] ... [SEP].
		// For a pair, we append B's tokens excluding the [CLS] token (index 0).
		var bIDs []uint32
		if len(encD.IDs) > 1 {
			bIDs = encD.IDs[1:]
		}

		docIDs[i] = bIDs

		pairLen := len(encQ.IDs) + len(bIDs)
		if pairLen > maxLen {
			maxLen = pairLen
		}
	}

	// Truncate maxLen if it exceeds maxSeqLength
	if maxLen > e.maxSeqLength {
		maxLen = e.maxSeqLength
	}

	seqLen := int64(maxLen)

	// 2. Build flat tensors
	flatIDs := make([]int64, batchSize*seqLen)
	flatMask := make([]int64, batchSize*seqLen)
	flatTokenTypes := make([]int64, batchSize*seqLen)

	for i := range documents {
		bIDs := docIDs[i]

		// Concatenate: A_IDs + B_IDs
		pairIDs := make([]uint32, 0, len(encQ.IDs)+len(bIDs))
		pairIDs = append(pairIDs, encQ.IDs...)
		pairIDs = append(pairIDs, bIDs...)

		// Truncate pair if needed
		if len(pairIDs) > int(seqLen) {
			pairIDs = pairIDs[:seqLen]
		}

		// Calculate lengths
		actualLen := len(pairIDs)
		lenA := len(encQ.IDs)

		baseOffset := int64(i) * seqLen
		for j := 0; j < actualLen; j++ {
			flatIDs[baseOffset+int64(j)] = int64(pairIDs[j])
			flatMask[baseOffset+int64(j)] = 1

			// token_type_ids: 0 for A (query), 1 for B (document)
			if j >= lenA {
				flatTokenTypes[baseOffset+int64(j)] = 1
			} else {
				flatTokenTypes[baseOffset+int64(j)] = 0
			}
		}
		// padding is 0 for all three tensors (already 0-initialized)
	}

	// 3. Create tensors
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

	typeTensor, err := ort.NewTensor(inputShape, flatTokenTypes)
	if err != nil {
		return nil, fmt.Errorf("create token_type_ids tensor: %w", err)
	}
	defer typeTensor.Destroy()

	// 4. Run inference
	inputs := []ort.Value{idsTensor, maskTensor, typeTensor}
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

	// 5. Extract output data
	outputTensor, ok := outputs[0].(*ort.Tensor[float32])
	if !ok {
		return nil, fmt.Errorf("unexpected output tensor type")
	}
	outputData := outputTensor.GetData()

	// Shape is usually [batch_size, 1] for regression tasks like ms-marco-MiniLM-L-6-v2
	// But it could be [batch_size, num_classes]. We'll assume the score is the first float
	// per batch item, which works for single-logit regression outputs.
	scores := make([]float32, batchSize)

	outputShape := outputTensor.GetShape()
	if len(outputShape) < 2 {
		return nil, fmt.Errorf("unexpected output shape: %v", outputShape)
	}

	logitsDim := int(outputShape[1])

	for i := int64(0); i < batchSize; i++ {
		scores[i] = outputData[i*int64(logitsDim)]
	}

	return scores, nil
}

// Close releases resources.
func (e *ONNXCrossEncoder) Close() error {
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
