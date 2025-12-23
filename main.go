package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RouteResponse struct {
	Message string `json:"message"`
}

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(w)
		if err := enc.Encode(RouteResponse{Message: "Hello, World!!!"}); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":4000", router))
}
