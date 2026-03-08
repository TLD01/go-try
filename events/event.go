package events

import (
	"time"

	"aerowatch.com/api/common"
	"aerowatch.com/api/receivers"
	"aerowatch.com/api/receivers/messages"
)

type Event struct {
	common.Persisted
	Source         receivers.Source            `json:"source"`
	Timestamp      time.Time                   `json:"timestamp"`
	VehicleMessage messages.AdsbVehicleMessage `json:"vehicle_message"`
}
