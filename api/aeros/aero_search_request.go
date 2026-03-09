package aeros

import (
	"aerowatch.com/api/common"
	"aerowatch.com/api/geolocation"
)


type AeroSearchRequest struct {
	Boundary geolocation.BoundingBox `json:"boundary"`
	TimeWindow common.TimeWindow `json:"timeWindow"`
}