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

		obj := &mmodels.Object{
			Class: className,
			Properties: map[string]interface{}{
				"text":                  doc.Text,
				"dateReported":          doc.DateReported,
				"projectName":           doc.ProjectName,
				"description":           doc.Description,
				"category":              doc.Category,
				"borough":               doc.Borough,
				"managingAgency":        doc.ManagingAgency,
				"clientAgency":          doc.ClientAgency,
				"currentPhase":          doc.CurrentPhase,
				"designStart":           doc.DesignStart,
				"budgetForecast":        doc.BudgetForecast,
				"latestBudgetChanges":   doc.LatestBudgetChanges,
				"totalBudgetChanges":    doc.TotalBudgetChanges,
				"forecastCompletion":    doc.ForecastCompletion,
				"latestScheduleChanges": doc.LatestScheduleChanges,
				"totalScheduleChanges":  doc.TotalScheduleChanges,
			},
		}

		// Add object to batch
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
			// sleep couple of sec
			time.Sleep(100 * time.Millisecond)
		}
	}

	return nil
}

func (m *WeaviateDBRepo) GetDocuments(q string) ([]*models.Doc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Raw GraphQL query
	query := fmt.Sprintf(`
    {
      Get {
        Document(
          nearText: { concepts: ["%s"] }
          limit: 5
        ) {
          text
          dateReported
          projectName
          description
          category
          borough
          managingAgency
          clientAgency
          currentPhase
          designStart
          budgetForecast
          latestBudgetChanges
          totalBudgetChanges
          forecastCompletion
          latestScheduleChanges
          totalScheduleChanges
        }
      }
    }
    `, q)

	res, err := m.DB.GraphQL().Raw().WithQuery(query).Do(ctx)
	if err != nil {
		return nil, err
	}

	// get docs from res
	var results []*models.Doc
	if getData, ok := res.Data["Get"].(map[string]interface{}); ok {
		if docs, ok := getData["Document"].([]interface{}); ok {
			for _, doc := range docs {
				if d, ok := doc.(map[string]interface{}); ok {
					results = append(results, &models.Doc{
						Text:                  d["text"].(string),
						DateReported:          d["dateReported"].(string),
						ProjectName:           d["projectName"].(string),
						Description:           d["description"].(string),
						Category:              d["category"].(string),
						Borough:               d["borough"].(string),
						ManagingAgency:        d["managingAgency"].(string),
						ClientAgency:          d["clientAgency"].(string),
						CurrentPhase:          d["currentPhase"].(string),
						DesignStart:           d["designStart"].(string),
						BudgetForecast:        d["budgetForecast"].(string),
						LatestBudgetChanges:   d["latestBudgetChanges"].(string),
						TotalBudgetChanges:    d["totalBudgetChanges"].(string),
						ForecastCompletion:    d["forecastCompletion"].(string),
						LatestScheduleChanges: d["latestScheduleChanges"].(string),
						TotalScheduleChanges:  d["totalScheduleChanges"].(string),
					})
				}
			}
		}
	}

	return results, nil
}
