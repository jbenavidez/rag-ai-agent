package main

import (
	"client/repository"
	dbrepo "client/repository/db_repo"
	"fmt"
	"log"
	"net/http"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

const (
	port = 4000
)

type Config struct {
	Llm            llms.Model
	WeaviateClient *weaviate.Client
	WDBRepo        repository.DatabaseRepo
}

func main() {
	app := Config{}
	fmt.Println("************* connecting to Weaviate *************")
	client, err := app.ConnectWeaviateDB()
	if err != nil {
		fmt.Printf("unable to connected %v", err)
		panic(err)
	}
	app.WDBRepo = &dbrepo.WeaviateDBRepo{DB: client}
	app.WeaviateClient = client
	fmt.Println("************* Loading Data *************")
	err = app.LoadDataSet()
	if err != nil {
		fmt.Println("somethings break", err)
	}
	fmt.Println("*************  Init Ollama *************")
	// Initialize Ollama LLMs
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		fmt.Println("failed to Initialize Ollama: ", err)
		panic(err)
	}
	fmt.Println("*************  Ollama Connected *************")
	app.Llm = llm
	log.Println("Starting agent on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
