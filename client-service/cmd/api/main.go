package main

import (
	"client/repository"
	dbrepo "client/repository/db_repo"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type Config struct {
	DB repository.DatabaseRepo
}

func main() {
	fmt.Println("starting url-shortener service...")
	app := Config{}

	conn := app.connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	fmt.Println("client up and running")

	// text := "Hello Postgres Vector"
	// embedding := simpleEmbedding(text)
	// fmt.Println("the text valinor", embedding)

	err := app.LoadCSVAndInsert()
	if err != nil {
		fmt.Println("sotmhthing break", err)
	}
	log.Println("Starting agent on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
