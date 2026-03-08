package embedding

import (
	"math"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gandazgul/mnemosyne/internal/config"
)

// testModelsDir returns the path to the test models directory.
// Tests are skipped if models are not available.
func testModelsDir(t *testing.T) string {
	t.Helper()

	// Walk up from the test file to find the project root.
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	modelsDir := filepath.Join(projectRoot, "models")

	// Check if embedding model directory exists with required files.
	embedDir := filepath.Join(modelsDir, "embeddinggemma-300m")
	if _, err := os.Stat(filepath.Join(embedDir, "tokenizer.json")); err != nil {
		t.Skipf("Skipping: embedding model not found at %s (run 'task download-models')", embedDir)
	}
	if _, err := os.Stat(filepath.Join(embedDir, "onnx", "model_quantized.onnx")); err != nil {
		t.Skipf("Skipping: ONNX model not found at %s (run 'task download-models')", embedDir)
	}

	return modelsDir
}

// testOnnxRuntimeLib returns the path to the ONNX Runtime shared library.
// Tests are skipped if the library is not available.
func testOnnxRuntimeLib(t *testing.T) string {
	t.Helper()

	// Check ONNXRUNTIME_SHARED_LIBRARY_PATH env var first.
	if envPath := os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH"); envPath != "" {
		return envPath
	}

	// Walk up from the test file to find the project root.
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	libDir := filepath.Join(projectRoot, "lib")

	// Try platform-specific library names.
	candidates := []string{
		"libonnxruntime.dylib",        // macOS
		"libonnxruntime.so",           // Linux
		"libonnxruntime.1.23.1.dylib", // macOS versioned
	}

	for _, name := range candidates {
		path := filepath.Join(libDir, name)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	t.Skipf("Skipping: ONNX Runtime library not found in %s (run 'task download-onnxruntime')", libDir)
	return ""
}

// testEmbedderConfig returns an EmbeddingConfig for testing.
func testEmbedderConfig(t *testing.T) config.EmbeddingConfig {
	t.Helper()
	modelsDir := testModelsDir(t)

	return config.EmbeddingConfig{
		ModelPath:      filepath.Join(modelsDir, "embeddinggemma-300m"),
		OnnxFile:       "onnx/model_quantized.onnx",
		Dimensions:     768,
		MaxSeqLength:   2048,
		QueryPrefix:    "task: search result | query: ",
		DocumentPrefix: "title: none | text: ",
	}
}

// --- Unit tests for math operations (no model needed) ---

func TestL2Normalize(t *testing.T) {
	v := []float32{3.0, 4.0}
	l2Normalize(v)

	// Expected: [0.6, 0.8] (3/5, 4/5)
	if math.Abs(float64(v[0])-0.6) > 1e-6 {
		t.Errorf("v[0] = %f, want 0.6", v[0])
	}
	if math.Abs(float64(v[1])-0.8) > 1e-6 {
		t.Errorf("v[1] = %f, want 0.8", v[1])
	}

	// Verify norm is ~1.0.
	var norm float64
	for _, val := range v {
		norm += float64(val) * float64(val)
	}
	norm = math.Sqrt(norm)
	if math.Abs(norm-1.0) > 1e-6 {
		t.Errorf("norm = %f, want 1.0", norm)
	}
}

func TestL2Normalize_ZeroVector(t *testing.T) {
	v := []float32{0.0, 0.0, 0.0}
	l2Normalize(v) // should not panic or produce NaN

	for i, val := range v {
		if val != 0 {
			t.Errorf("v[%d] = %f, want 0.0", i, val)
		}
	}
}

func TestMeanPool(t *testing.T) {
	// Batch of 1, seq_len=3, hidden_dim=2
	// Token 0: [1, 2], Token 1: [3, 4], Token 2: [5, 6] (padding)
	hiddenStates := []float32{1, 2, 3, 4, 5, 6}
	mask := []int64{1, 1, 0} // only first 2 tokens are real

	result := meanPool(hiddenStates, mask, 0, 3, 2)

	// Expected: mean of [1,2] and [3,4] = [2, 3]
	expected := []float32{2.0, 3.0}
	for i, val := range result {
		if math.Abs(float64(val)-float64(expected[i])) > 1e-6 {
			t.Errorf("result[%d] = %f, want %f", i, val, expected[i])
		}
	}
}

func TestMeanPool_AllMasked(t *testing.T) {
	hiddenStates := []float32{1, 2, 3, 4}
	mask := []int64{0, 0}

	result := meanPool(hiddenStates, mask, 0, 2, 2)

	// All zeros when everything is masked.
	for i, val := range result {
		if val != 0 {
			t.Errorf("result[%d] = %f, want 0.0", i, val)
		}
	}
}

func TestMeanPool_Batch(t *testing.T) {
	// Batch of 2, seq_len=2, hidden_dim=2
	// Batch 0: [1, 2], [3, 4]  mask [1, 1] -> mean [2, 3]
	// Batch 1: [10, 20], [30, 40]  mask [1, 0] -> mean [10, 20]
	hiddenStates := []float32{
		1, 2, 3, 4, // batch 0
		10, 20, 30, 40, // batch 1
	}
	mask := []int64{
		1, 1, // batch 0
		1, 0, // batch 1
	}

	result0 := meanPool(hiddenStates, mask, 0, 2, 2)
	result1 := meanPool(hiddenStates, mask, 1, 2, 2)

	expected0 := []float32{2.0, 3.0}
	expected1 := []float32{10.0, 20.0}

	for i := range expected0 {
		if math.Abs(float64(result0[i])-float64(expected0[i])) > 1e-6 {
			t.Errorf("batch0[%d] = %f, want %f", i, result0[i], expected0[i])
		}
	}
	for i := range expected1 {
		if math.Abs(float64(result1[i])-float64(expected1[i])) > 1e-6 {
			t.Errorf("batch1[%d] = %f, want %f", i, result1[i], expected1[i])
		}
	}
}

// --- Tokenizer tests (require model files) ---

func TestTokenizer_Encode(t *testing.T) {
	modelsDir := testModelsDir(t)
	modelDir := filepath.Join(modelsDir, "embeddinggemma-300m")

	tok, err := NewTokenizer(modelDir, 2048)
	if err != nil {
		t.Fatalf("NewTokenizer: %v", err)
	}
	defer tok.Close()

	enc, err := tok.Encode("hello world")
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	if enc.Length == 0 {
		t.Fatal("Encode returned 0 tokens")
	}
	if len(enc.InputIDs) != enc.Length {
		t.Errorf("InputIDs length %d != Length %d", len(enc.InputIDs), enc.Length)
	}
	if len(enc.AttentionMask) != enc.Length {
		t.Errorf("AttentionMask length %d != Length %d", len(enc.AttentionMask), enc.Length)
	}

	// All attention mask values should be 1 (no padding for a single encode).
	for i, m := range enc.AttentionMask {
		if m != 1 {
			t.Errorf("AttentionMask[%d] = %d, want 1", i, m)
		}
	}

	t.Logf("Tokenized 'hello world' into %d tokens", enc.Length)
	t.Logf("InputIDs: %v", enc.InputIDs[:min(10, len(enc.InputIDs))])
}

func TestTokenizer_Truncation(t *testing.T) {
	modelsDir := testModelsDir(t)
	modelDir := filepath.Join(modelsDir, "embeddinggemma-300m")

	// Create tokenizer with very short max length.
	tok, err := NewTokenizer(modelDir, 5)
	if err != nil {
		t.Fatalf("NewTokenizer: %v", err)
	}
	defer tok.Close()

	enc, err := tok.Encode("this is a long sentence that should be truncated to only five tokens")
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	if enc.Length > 5 {
		t.Errorf("Length %d exceeds maxSeqLen 5", enc.Length)
	}
}

func TestTokenizer_EncodeBatch(t *testing.T) {
	modelsDir := testModelsDir(t)
	modelDir := filepath.Join(modelsDir, "embeddinggemma-300m")

	tok, err := NewTokenizer(modelDir, 2048)
	if err != nil {
		t.Fatalf("NewTokenizer: %v", err)
	}
	defer tok.Close()

	texts := []string{"short", "this is a longer sentence with more tokens"}
	encoded, maxLen, err := tok.EncodeBatch(texts)
	if err != nil {
		t.Fatalf("EncodeBatch: %v", err)
	}

	if len(encoded) != 2 {
		t.Fatalf("expected 2 encoded texts, got %d", len(encoded))
	}

	// Both should be padded to maxLen.
	for i, enc := range encoded {
		if len(enc.InputIDs) != maxLen {
			t.Errorf("text %d: InputIDs length %d != maxLen %d", i, len(enc.InputIDs), maxLen)
		}
		if len(enc.AttentionMask) != maxLen {
			t.Errorf("text %d: AttentionMask length %d != maxLen %d", i, len(enc.AttentionMask), maxLen)
		}
	}

	// The shorter text should have padding (0s in attention mask).
	if encoded[0].Length >= encoded[1].Length {
		t.Errorf("expected first text to be shorter: %d >= %d", encoded[0].Length, encoded[1].Length)
	}

	// Check that padding positions have mask = 0.
	shortEnc := encoded[0]
	for i := shortEnc.Length; i < maxLen; i++ {
		if shortEnc.AttentionMask[i] != 0 {
			t.Errorf("padding position %d has mask %d, want 0", i, shortEnc.AttentionMask[i])
		}
		if shortEnc.InputIDs[i] != 0 {
			t.Errorf("padding position %d has ID %d, want 0", i, shortEnc.InputIDs[i])
		}
	}

	t.Logf("Batch tokenized: lengths [%d, %d], padded to %d", encoded[0].Length, encoded[1].Length, maxLen)
}

// --- ONNX Embedder integration tests (require model + ONNX Runtime) ---

func TestONNXEmbedder_Embed(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	cfg := testEmbedderConfig(t)

	if err := InitONNXRuntime(libPath); err != nil {
		t.Fatalf("InitONNXRuntime: %v", err)
	}
	// Note: don't destroy environment in tests, as it's shared globally.

	embedder, err := NewONNXEmbedder(cfg)
	if err != nil {
		t.Fatalf("NewONNXEmbedder: %v", err)
	}
	defer embedder.Close()

	embedding, err := embedder.Embed("hello world")
	if err != nil {
		t.Fatalf("Embed: %v", err)
	}

	// Check dimensions.
	if len(embedding) != cfg.Dimensions {
		t.Errorf("embedding dimensions = %d, want %d", len(embedding), cfg.Dimensions)
	}

	// Check that it's L2-normalized (norm should be ~1.0).
	var norm float64
	for _, v := range embedding {
		norm += float64(v) * float64(v)
	}
	norm = math.Sqrt(norm)
	if math.Abs(norm-1.0) > 1e-3 {
		t.Errorf("L2 norm = %f, want ~1.0", norm)
	}

	// Check that values are not all zeros.
	allZero := true
	for _, v := range embedding {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Error("embedding is all zeros")
	}

	t.Logf("Embedding dimensions: %d, L2 norm: %.6f", len(embedding), norm)
	t.Logf("First 5 values: %v", embedding[:min(5, len(embedding))])
}

func TestONNXEmbedder_Deterministic(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	cfg := testEmbedderConfig(t)

	if err := InitONNXRuntime(libPath); err != nil {
		t.Fatalf("InitONNXRuntime: %v", err)
	}

	embedder, err := NewONNXEmbedder(cfg)
	if err != nil {
		t.Fatalf("NewONNXEmbedder: %v", err)
	}
	defer embedder.Close()

	// Same input should produce the same output.
	emb1, err := embedder.Embed("deterministic test")
	if err != nil {
		t.Fatalf("Embed 1: %v", err)
	}
	emb2, err := embedder.Embed("deterministic test")
	if err != nil {
		t.Fatalf("Embed 2: %v", err)
	}

	if len(emb1) != len(emb2) {
		t.Fatalf("different lengths: %d vs %d", len(emb1), len(emb2))
	}

	for i := range emb1 {
		if emb1[i] != emb2[i] {
			t.Errorf("emb1[%d] = %f != emb2[%d] = %f", i, emb1[i], i, emb2[i])
			break
		}
	}
}

func TestONNXEmbedder_DifferentInputsDifferentOutputs(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	cfg := testEmbedderConfig(t)

	if err := InitONNXRuntime(libPath); err != nil {
		t.Fatalf("InitONNXRuntime: %v", err)
	}

	embedder, err := NewONNXEmbedder(cfg)
	if err != nil {
		t.Fatalf("NewONNXEmbedder: %v", err)
	}
	defer embedder.Close()

	emb1, err := embedder.Embed("the cat sat on the mat")
	if err != nil {
		t.Fatalf("Embed 1: %v", err)
	}
	emb2, err := embedder.Embed("quantum mechanics and string theory")
	if err != nil {
		t.Fatalf("Embed 2: %v", err)
	}

	// Should be different embeddings.
	same := true
	for i := range emb1 {
		if emb1[i] != emb2[i] {
			same = false
			break
		}
	}
	if same {
		t.Error("different inputs produced identical embeddings")
	}

	// Check cosine similarity (should be < 1.0 for different texts).
	var dot float64
	for i := range emb1 {
		dot += float64(emb1[i]) * float64(emb2[i])
	}
	t.Logf("Cosine similarity between unrelated texts: %.6f", dot)
}

func TestONNXEmbedder_QueryDocumentPrefixes(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	cfg := testEmbedderConfig(t)

	if err := InitONNXRuntime(libPath); err != nil {
		t.Fatalf("InitONNXRuntime: %v", err)
	}

	embedder, err := NewONNXEmbedder(cfg)
	if err != nil {
		t.Fatalf("NewONNXEmbedder: %v", err)
	}
	defer embedder.Close()

	text := "golang concurrency patterns"

	queryEmb, err := embedder.EmbedQuery(text)
	if err != nil {
		t.Fatalf("EmbedQuery: %v", err)
	}

	docEmb, err := embedder.EmbedDocument(text)
	if err != nil {
		t.Fatalf("EmbedDocument: %v", err)
	}

	// Query and document embeddings for the same text should be different
	// (due to different prefixes) but somewhat similar.
	different := false
	for i := range queryEmb {
		if queryEmb[i] != docEmb[i] {
			different = true
			break
		}
	}
	if !different {
		t.Error("query and document embeddings are identical (prefixes not applied?)")
	}

	// Cosine similarity should be positive (related content).
	var dot float64
	for i := range queryEmb {
		dot += float64(queryEmb[i]) * float64(docEmb[i])
	}
	t.Logf("Cosine similarity (query vs doc, same text): %.6f", dot)
	if dot < 0 {
		t.Errorf("cosine similarity = %.6f, expected positive", dot)
	}
}

func TestONNXEmbedder_EmbedBatch(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	cfg := testEmbedderConfig(t)

	if err := InitONNXRuntime(libPath); err != nil {
		t.Fatalf("InitONNXRuntime: %v", err)
	}

	embedder, err := NewONNXEmbedder(cfg)
	if err != nil {
		t.Fatalf("NewONNXEmbedder: %v", err)
	}
	defer embedder.Close()

	texts := []string{
		"machine learning algorithms",
		"web development frameworks",
		"database optimization techniques",
	}

	embeddings, err := embedder.EmbedBatch(texts)
	if err != nil {
		t.Fatalf("EmbedBatch: %v", err)
	}

	if len(embeddings) != len(texts) {
		t.Fatalf("expected %d embeddings, got %d", len(texts), len(embeddings))
	}

	for i, emb := range embeddings {
		if len(emb) != cfg.Dimensions {
			t.Errorf("embedding %d: dimensions = %d, want %d", i, len(emb), cfg.Dimensions)
		}

		// Check L2 norm.
		var norm float64
		for _, v := range emb {
			norm += float64(v) * float64(v)
		}
		norm = math.Sqrt(norm)
		if math.Abs(norm-1.0) > 1e-3 {
			t.Errorf("embedding %d: L2 norm = %f, want ~1.0", i, norm)
		}
	}

	// Verify batch results match individual results.
	for i, text := range texts {
		singleEmb, err := embedder.Embed(text)
		if err != nil {
			t.Fatalf("single Embed(%q): %v", text, err)
		}

		// Due to padding differences in batch vs single, results may differ slightly.
		// Check cosine similarity is very high (> 0.99).
		var dot float64
		for j := range singleEmb {
			dot += float64(singleEmb[j]) * float64(embeddings[i][j])
		}
		if dot < 0.99 {
			t.Errorf("text %d: batch vs single cosine similarity = %.6f, want > 0.99", i, dot)
		}
	}
}

func TestONNXEmbedder_MRLTruncation(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	modelsDir := testModelsDir(t)

	if err := InitONNXRuntime(libPath); err != nil {
		t.Fatalf("InitONNXRuntime: %v", err)
	}

	// Test with reduced dimensions (MRL truncation).
	cfg := config.EmbeddingConfig{
		ModelPath:      filepath.Join(modelsDir, "embeddinggemma-300m"),
		OnnxFile:       "onnx/model_quantized.onnx",
		Dimensions:     256, // Truncated from 768
		MaxSeqLength:   2048,
		QueryPrefix:    "task: search result | query: ",
		DocumentPrefix: "title: none | text: ",
	}

	embedder, err := NewONNXEmbedder(cfg)
	if err != nil {
		t.Fatalf("NewONNXEmbedder: %v", err)
	}
	defer embedder.Close()

	embedding, err := embedder.Embed("test MRL truncation")
	if err != nil {
		t.Fatalf("Embed: %v", err)
	}

	if len(embedding) != 256 {
		t.Errorf("embedding dimensions = %d, want 256", len(embedding))
	}

	// Should still be L2-normalized after truncation.
	var norm float64
	for _, v := range embedding {
		norm += float64(v) * float64(v)
	}
	norm = math.Sqrt(norm)
	if math.Abs(norm-1.0) > 1e-3 {
		t.Errorf("L2 norm = %f, want ~1.0", norm)
	}
}

func TestONNXEmbedder_Dimensions(t *testing.T) {
	libPath := testOnnxRuntimeLib(t)
	cfg := testEmbedderConfig(t)

	if err := InitONNXRuntime(libPath); err != nil {
		t.Fatalf("InitONNXRuntime: %v", err)
	}

	embedder, err := NewONNXEmbedder(cfg)
	if err != nil {
		t.Fatalf("NewONNXEmbedder: %v", err)
	}
	defer embedder.Close()

	if embedder.Dimensions() != 768 {
		t.Errorf("Dimensions() = %d, want 768", embedder.Dimensions())
	}
}
