package repository

import (
	"time"

	"aerowatch.com/api/receivers"
	"aerowatch.com/api/receivers/messages"
	"aerowatch.com/api/repository"
)

type EventEntity struct {
	repository.DBEntity
	Source         receivers.Source                      `bson:"source"`
	Timestamp      time.Time                   `bson:"timestamp"`
	VehicleMessage messages.AdsbVehicleMessage `bson:"vehicleMessage"`
}
