package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *RagConfig) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Get("/test", a.TestEndpoint)

	return mux
}
