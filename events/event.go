package events

import (
	"time"

	"aerowatch.com/api/common"
	"aerowatch.com/api/messages"
)

type Source struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Event struct {
	common.Persisted
	Source    Source    `json:"source"`
	Timestamp time.Time `json:"timestamp"`
	VehicleMessage messages.AdsbVehicleMessage `json:"vehicle_message"`
}
