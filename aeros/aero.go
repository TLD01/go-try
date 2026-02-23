package aeros

import (
	"time"

	"aerowatch.com/api/common"
	"aerowatch.com/api/geolocation"
	"aerowatch.com/api/jsonutil"
)

type Aero struct {
	common.Persisted
	Callsign     string            `json:"callsign"`
	IcaoAddress  string            `json:"icao_address"`
	Model        string            `json:"model"`
	LastSeen     time.Time         `json:"last_seen"`
	LastPosition geolocation.Point `json:"last_position"`
}

func (a *Aero) Serialize() string {
	str, _ := jsonutil.JsonSerialize(a)
	return str
}

func (a *Aero) Epoch() int64 {
	return a.LastSeen.Unix()
}
