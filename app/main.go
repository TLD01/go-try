package main

import (
	"fmt"
	"time"

	"aerowatch.com/api/aeros"
	"aerowatch.com/api/common"
	"aerowatch.com/api/geofence"
	"aerowatch.com/api/geolocation"
	"aerowatch.com/api/jsonutil"
	"aerowatch.com/api/lfv"
)

func main() {
	fmt.Println("app is ready")
	aero := &aeros.Aero{
		Callsign:    "TEST123",
		IcaoAddress: "ABC123",
		Model:       "Test Model",
		LastSeen:    time.Now(),
		LastPosition: geolocation.Point{
			Latitude:  40.7128,
			Longitude: -74.0060,
		},
		Persisted: common.Persisted{
			Id:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	json := `{
		"id": "1",
		"createdAt": "2026-02-22T22:59:56.883387035+01:00",
		"updatedAt": "2026-02-22T22:59:56.883387165+01:00",
		"callsign": "TEST123",
		"icao_address": "ABC123",
		"model": "Test Model",
		"last_seen": "2026-02-22T22:59:56.883386864+01:00",
		"last_position": {
			"latitude": 40.7128,
			"longitude": -74.006
		}
		}`

	var notificationType geofence.NotificationType = geofence.NotificationTypeEmail
	var unknownType geofence.NotificationType = geofence.NotificationType("XYZ")

	fmt.Printf("Notification Type: %s\n", notificationType)
	fmt.Printf("Unknown Notification Type: %s\n", unknownType)

	var aero2 aeros.Aero
	err := jsonutil.JsonDeserialize(json, &aero2)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}

	for _, layer := range lfv.All() {
		fmt.Printf("Map Layer: %s, Description: %s\n", layer.Name(), layer.Description)
	}

	fmt.Printf("Deserialized Aero: %+v\n", aero2)
	fmt.Println(aero.Serialize())
}
