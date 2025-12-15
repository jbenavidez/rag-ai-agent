package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
)

func (app *RagConfig) ConnectWeaviateDB() (*weaviate.Client, error) {
	ctx := context.Background()
	client := weaviate.New(weaviate.Config{
		Scheme: "http",
		Host:   "weaviate:8080",
	})

	// Wait until Weaviate is ready
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		_, err := client.Schema().Getter().Do(ctx)
		if err == nil {
			//break since Weaviate is ready
			break
		}
		fmt.Println("Waiting for Weaviate to be ready...")
		time.Sleep(2 * time.Second)
	}

	// Check one last time
	_, err := client.Schema().Getter().Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("Weaviate not ready after retries: %w", err)
	}

	// check is class exist
	className := "Document"
	exists, err := client.Schema().ClassExistenceChecker().
		WithClassName(className).
		Do(context.Background())

	if err != nil {
		log.Fatalf("Failed to check class existence: %v", err)
	}
	if !exists {
		fmt.Println("creating class for documents exist")
		// create class
		documentClass := &models.Class{
			Class:       className,
			Description: "Collection of documents",
			Vectorizer:  "text2vec-openai",
			Properties: []*models.Property{
				{Name: "body", DataType: schema.DataTypeText.PropString()},
			},
		}

		if err := client.Schema().ClassCreator().WithClass(documentClass).Do(ctx); err != nil {
			return nil, err
		}
		fmt.Println("class for documents created")
	}

	return client, nil
}
