package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"aerowatch.com/api/common"
	"aerowatch.com/api/events"
	"aerowatch.com/api/repository"
)

type EventController struct {
	service *events.EventsService
}

func NewEventController(service *events.EventsService) *EventController {
	return &EventController{service: service}
}

// RegisterRoutes attaches all routes to the given mux.
// Base path: /api/v1/events
func (c *EventController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/events", c.search)
}

// GET /api/v1/events?icao={icao}&from={RFC3339}&to={RFC3339}
func (c *EventController) search(w http.ResponseWriter, r *http.Request) {
	icao := r.URL.Query().Get("icao")
	if icao == "" {
		http.Error(w, "icao query param is required", http.StatusBadRequest)
		return
	}

	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		http.Error(w, "from must be a valid RFC3339 timestamp", http.StatusBadRequest)
		return
	}

	to, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		http.Error(w, "to must be a valid RFC3339 timestamp", http.StatusBadRequest)
		return
	}

	result, err := c.service.Search(r.Context(), icao, common.TimeWindow{Start: from, End: to})
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
