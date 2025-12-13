package main

import (
	"client/models"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func (c *Config) LoadDataSet() error {
	totalDocs, err := c.WDBRepo.GetTotalDocs()
	if err != nil {
		return err
	}
	if totalDocs > 0 {
		fmt.Printf("the total docs %v", totalDocs)
		return nil
	}
	fmt.Println("******** Getting  data from csv ********")
	docs, err := c.GetData()
	if err != nil {
		return err
	}
	fmt.Printf("******** Total rows to insert %v ********", len(docs))

	err = c.WDBRepo.InsertDocuments(docs)
	if err != nil {
		return err
	}
	fmt.Println("******** Data load completed ********")
	return nil

}

func (c *Config) GetData() ([]*models.Doc, error) {
	f, err := os.Open("./cmd/data/data.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var docs []*models.Doc
	// skip ehader
	for _, row := range rows[1:] {
		projectcName := row[2]
		description := row[3]
		text := strings.TrimSpace(fmt.Sprintf("%s — %s", projectcName, description))
		d := models.Doc{
			ID:          row[1],
			Text:        text,
			ProjectName: projectcName,
			Description: description,
		}
		docs = append(docs, &d)
	}

	return docs, nil
}
