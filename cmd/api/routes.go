package main

import (
	"github.com/PrinceNarteh/kanban-api/internals/handlers"
	"github.com/go-chi/chi/v5"
)

func (app *application) initRoutes(r *chi.Mux) {
	r.Get("/", handlers.HealthCheck)
}
