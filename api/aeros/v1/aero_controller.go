package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"aerowatch.com/api/aeros"
	"aerowatch.com/api/repository"
)

type AeroController struct {
	service *aeros.AeroService
}

func NewAeroController(service *aeros.AeroService) *AeroController {
	return &AeroController{service: service}
}

// RegisterRoutes attaches all routes to the given mux.
// Base path: /api/v1/aeros
func (c *AeroController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/aeros/{icao}", c.getByIcao)
}

// GET /api/v1/aeros/{icao}
func (c *AeroController) getByIcao(w http.ResponseWriter, r *http.Request) {
	icao := r.PathValue("icao")

	aero, err := c.service.Get(r.Context(), icao)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "aero not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aero)
}
