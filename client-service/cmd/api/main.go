package main

import (
	pb "client/proto/generated"
	"client/repository"
	dbrepo "client/repository/db_repo"
	"fmt"
	"log"
	"net/http"
)

const (
	gRpcPort = "50001"
	port     = 4000
)

type Config struct {
	DB         repository.DatabaseRepo
	GRPCClient pb.EmbeddingServiceClient
}

func main() {
	fmt.Println("starting url-shortener service...")
	app := Config{}

	conn := app.connectToDB()
	if conn == nil {
		log.Panic("Cannot connect to Postgres!")
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	fmt.Println("client up and running")

	// connect to GRPC
	app.GRPCClient = newGRPCConn()
	fmt.Println("GRPC is conencted")
	// init helper
	NewGrpcHelper(&app)
	// load data
	err := app.LoadData()
	if err != nil {
		fmt.Println("somethings break", err)
	}

	_, err = app.DB.GetEmbeddingDocument("valinor", 3)

	if err != nil {
		fmt.Println("failed_test", err)
		panic(err)
	}
	log.Println("Starting agent on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
