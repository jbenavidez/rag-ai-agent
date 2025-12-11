package models

type Document struct {
	Text          string
	EmbeddingText string
	ProjectName   string
	Description   string
	Distance      float64
}
