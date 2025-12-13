package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *Config) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Get("/test", a.TestEndpoint)

	return mux
}
