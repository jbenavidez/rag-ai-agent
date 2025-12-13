package dbrepo

import (
	"client/models"
	"context"
	"fmt"
	"time"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
	mmodels "github.com/weaviate/weaviate/entities/models"
)

type WeaviateDBRepo struct {
	DB *weaviate.Client
}

const timeout = time.Second * 3
const className = "Document"

func (m *WeaviateDBRepo) Connection() *weaviate.Client {
	return m.DB
}

func (m *WeaviateDBRepo) GetTotalDocs() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := m.DB.Data().ObjectsGetter().
		WithClassName(className).
		WithLimit(10000).
		Do(ctx)
	if err != nil {
		return 0, fmt.Errorf("error fetching documents: %w", err)
	}

	return len(res), nil
}

func (m *WeaviateDBRepo) InsertDocument(doc *models.Doc) error {
	fmt.Println("the document_id:", doc.ID)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// insert doc
	_, err := m.DB.Data().Creator().
		WithClassName(className).
		WithProperties(map[string]interface{}{
			"text": doc.Text,
		}).
		Do(ctx)
	if err != nil {
		fmt.Printf("unable_to_insert %v", err)
		return err
	}
	fmt.Printf("document was saved")
	return nil
}

func (m *WeaviateDBRepo) InsertDocuments(docs []*models.Doc) error {
	batchSize := 50
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	batcher := m.DB.Batch().ObjectsBatcher()
	count := 0

	for i, doc := range docs {
		// Create a object for eachdoc
		obj := &mmodels.Object{
			Class: className,
			Properties: map[string]interface{}{
				"text": doc.Text,
			},
		}

		// append to batch
		batcher = batcher.WithObject(obj)
		count++

		// Insert batch
		if count >= batchSize || i == len(docs)-1 {
			_, err := batcher.Do(ctx)
			if err != nil {
				fmt.Printf("Error inserting batch: %v\n", err)
				return err
			}

			fmt.Printf("Inserted %d documents\n", count)

			// Reset batch and counter
			batcher = m.DB.Batch().ObjectsBatcher()
			count = 0

			// sleep to prevent api lock
			time.Sleep(100 * time.Millisecond)
		}
	}

	return nil
}
