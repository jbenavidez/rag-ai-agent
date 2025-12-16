package main

import (
	"frontend/internal/config"
	"frontend/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	// display chat-room
	mux.Get("/", handlers.Repo.ChatRoom)
	//TODO: make it ws after integrations
	mux.Post("/answer-user-question", handlers.Repo.AnswerUserQuestion)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
