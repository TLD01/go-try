package aeros

import (
	"aerowatch.com/api/geolocation"
	"aerowatch.com/api/repository"
	"time"
)

type AeroEntity struct {
	repository.DBEntity
	Callsign     string            `bson:"callsign"`
	IcaoAddress  string            `bson:"icaoAddress"`
	Model        string            `bson:"model"`
	LastSeen     time.Time         `bson:"lastSeen"`
	LastPosition geolocation.Point `bson:"lastPosition"`
}
