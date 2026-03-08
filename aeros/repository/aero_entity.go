package repository

import (
	"time"

	"aerowatch.com/api/geolocation"
	"aerowatch.com/api/receivers/messages"
	"aerowatch.com/api/repository"
)

type AeroEntity struct {
	repository.DBEntity
	Callsign     string                      `bson:"callsign"`
	IcaoAddress  string                      `bson:"icaoAddress"`
	Model        string                      `bson:"model"`
	LastSeen     time.Time                   `bson:"lastSeen"`
	LastPosition geolocation.Point           `bson:"lastPosition"`
	LastMessage  messages.AdsbVehicleMessage `bson:"lastMessage"`
}

