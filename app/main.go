package main

import (
	"fmt"
	"time"

	"aerowatch.com/api/aeros"
	"aerowatch.com/api/common"
	"aerowatch.com/api/geofence"
	"aerowatch.com/api/geolocation"
	"aerowatch.com/api/jsonutil"
	"aerowatch.com/api/lfv/altitude_unit"
	"aerowatch.com/api/lfv/reference_system"
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

	var altitudeUnit1 altitude_unit.AltitudeUnit
	altitudeUnit1 = altitude_unit.GND_SFC
	var altitudeUnit2 altitude_unit.AltitudeUnit = altitude_unit.GND_SFC
	var altitudeUnit3 *altitude_unit.AltitudeUnit = &altitude_unit.GND_SFC

	fmt.Println(altitudeUnit1.String())
	fmt.Println(altitudeUnit2.String())

	fmt.Printf("%p\n", &altitudeUnit1)
	fmt.Printf("%p\n", &altitudeUnit2)
	fmt.Printf("%p\n", altitudeUnit3)

	jsonBytes, err := altitudeUnit1.MarshalJSON()
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Printf("%s\n", string(jsonBytes))

	fmt.Println(&altitudeUnit1 == &altitudeUnit2)
	fmt.Println(altitudeUnit1.Equal(altitudeUnit2))

	fmt.Printf("Deserialized Aero: %+v\n", aero2)
	fmt.Println(aero.Serialize())

	 allReferencePoints := reference_system.All()
	 for _, rp := range allReferencePoints {
		 fmt.Printf("Reference Point: %s - %s\n", rp.Name(), rp.Description)
	 }
}
