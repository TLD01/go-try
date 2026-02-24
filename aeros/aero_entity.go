package aeros

import (
	"time"
	"aerowatch.com/api/repository"
	"aerowatch.com/api/geolocation"
)

type AeroEntity struct {
	repository.DBEntity
	Callsign     string            `json:"callsign"`
	IcaoAddress  string            `json:"icao_address"`
	Model        string            `json:"model"`
	LastSeen     time.Time         `json:"last_seen"`
	LastPosition geolocation.Point `json:"last_position"`
}