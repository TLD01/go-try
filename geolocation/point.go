package geolocation

import "fmt"

type Point struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

func (p *Point) String() string {
	return fmt.Sprintf("%f;%f", p.Latitude, p.Longitude)
}
