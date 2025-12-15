package utils

import (
	"math"
	"strings"
)

func SimpleEmbedding(text string) []float64 {
	// if you have an Open AI Key, you can use openAI Embedding model to make it more robust.
	// Since this is simple RAG Ai Agent, I will keep it as simple as possible
	// TODO: implement Python micro-service to   the embeding
	text = strings.ToLower(text)
	vector := make([]float64, 32)

	for i, c := range text {
		vector[i%32] += float64(c)
	}

	var norm float64
	for _, v := range vector {
		norm += v * v
	}
	norm = math.Sqrt(norm)

	for i := range vector {
		vector[i] /= norm
	}

	return vector
}
