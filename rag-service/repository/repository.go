package repository

import (
	"client/models"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

type DatabaseRepo interface {
	Connection() *weaviate.Client
	GetTotalDocs() (int, error)
	InsertDocument(doc *models.Doc) error
	InsertDocuments(docs []*models.Doc) error
	GetDocuments(q string) ([]*models.Doc, error)
}
