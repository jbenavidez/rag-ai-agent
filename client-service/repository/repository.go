package repository

import (
	"client/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetTotalDocuments() (int, error)
	InsertDocument(documents []models.Document) error
	GetEmbeddingDocument(queryText string, topK int) ([]models.Document, error)
}
