package setup

func OnnxRuntimeLibNameForTest() string {
	return onnxRuntimeLibName()
}

func AllModelsForTest() []EmbedModel {
	return []EmbedModel{EmbeddingModel, RerankerModel}
}
