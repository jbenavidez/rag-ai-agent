package main

import (
	"fmt"
	"log"
	"net/http"

	"frontend/internal/config"
	"frontend/internal/handlers"
	"frontend/internal/render"
	pb "frontend/proto/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":4000"

var app config.AppConfig

func main() {
	//set grpc client

	conn, err := grpc.Dial("rag-ai-agent-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println("unable to connected to grpc server")
		panic(err)
	}
	client := pb.NewAIAgentServiceClient(conn)
	app.GRPCClient = client

	// //init web
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	log.Printf("Front-end server connected on port %s ", port)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
