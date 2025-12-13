package repository

import (
	"client/models"
	"database/sql"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetTotalDocuments() (int, error)
	InsertDocument(documents []models.Document) error
	GetEmbeddingDocument(queryText []float32, topK int, keyword string) ([]string, error)
}

type WeaviateDatabaseRepo interface {
	Connection() *weaviate.Client
	GetTotalDocs() (int, error)
	InsertDocument(doc *models.Doc) error
	InsertDocuments(docs []*models.Doc) error
}
