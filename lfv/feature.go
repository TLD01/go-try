package lfv

import (
	"aerowatch.com/api/common"
	"github.com/paulmach/orb/geojson"
)

type Feature struct {
	common.Persisted
	Type       string            `json:"type" bson:"type"`
	FeatureID  string            `json:"featureId" bson:"featureId"`
	Geometry   *geojson.Geometry `json:"geometry,omitempty" bson:"geometry,omitempty"`
	Properties map[string]any    `json:"properties,omitempty" bson:"properties,omitempty"`
}

type FeatureCollection struct {
	Type           string     `json:"type" bson:"type"`
	Features       []*Feature `json:"features,omitempty" bson:"features,omitempty"`
	TotalFeatures  int        `json:"totalFeatures" bson:"totalFeatures"`
	NumberMatched  int        `json:"numberMatched" bson:"numberMatched"`
	NumberReturned int        `json:"numberReturned" bson:"numberReturned"`
	Timestamp      string     `json:"timestamp" bson:"timestamp"`
	Crs            *Crs       `json:"crs,omitempty" bson:"crs,omitempty"`
}
