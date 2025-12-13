package main

import (
	"client/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Config) TestEndpoint(w http.ResponseWriter, r *http.Request) {

	// test query
	searchQuery := "Interceptor Sewers at Various Locations in the Boroughs of M"
	result, err := c.WDBRepo.GetDocuments(searchQuery)
	if err != nil {
		fmt.Println("unable to get data", err)
		return
	}
	//set test resposne
	jsonResponse := make(map[string][]*models.Doc)
	jsonResponse["message"] = result
	//set response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)

}
