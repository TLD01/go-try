package events

import (
	"time"

	"aerowatch.com/api/messages"
	"aerowatch.com/api/repository"
)

type EventEntity struct {
	repository.DBEntity
	Source         Source                      `bson:"source"`
	Timestamp      time.Time                   `bson:"timestamp"`
	VehicleMessage messages.AdsbVehicleMessage `bson:"vehicleMessage"`
}
