package messages

import "aerowatch.com/api/geolocation"

type AdsbVehicleMessage struct {
	IcaoAddress  int    `json:"icao_address" bson:"icaoAddress"`
	Lat          int    `json:"lat" bson:"lat"`
	Lon          int    `json:"lon" bson:"lon"`
	AltitudeType byte   `json:"altitude_type" bson:"altitudeType"`
	Altitude     int    `json:"altitude" bson:"altitude"`
	Heading      int    `json:"heading" bson:"heading"`
	HorVelocity  int    `json:"hor_velocity" bson:"horVelocity"`
	VerVelocity  int    `json:"ver_velocity" bson:"verVelocity"`
	Callsign     string `json:"callsign" bson:"callsign"`
	EmitterType  byte   `json:"emitter_type" bson:"emitterType"`
	Tslc         int    `json:"tslc" bson:"tslc"`
	Flags        int    `json:"flags" bson:"flags"`
	Squawk       int    `json:"squawk" bson:"squawk"`
}

func (m *AdsbVehicleMessage) Position() geolocation.Point {

	if m == nil {
		return geolocation.Point{}
	}

	return geolocation.Point{
		Latitude:  float64(m.Lat) / 10_000_000.0,
		Longitude: float64(m.Lon) / 10_000_000.0,
	}
}
