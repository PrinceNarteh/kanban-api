package handlers

import (
	"net/http"

	"github.com/PrinceNarteh/kanban-api/internals/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"version": "1.0.1",
	}

	utils.WriteResponse(w, http.StatusOK, data)
}
