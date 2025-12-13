package models

type Document struct {
	Text          string
	EmbeddingText []float32
	ProjectName   string
	Description   string
	Distance      float64
}

type Doc struct {
	ID          string
	Text        string
	ProjectName string
	Description string
}
