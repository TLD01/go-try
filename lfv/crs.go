package lfv

import "fmt"

type Crs struct {
	// CRS code, e.g. "EPSG:4326"
	Type string `json:"type" bson:"type"`
	Properties map[string]string `json:"properties,omitempty" bson:"properties,omitempty"`
}

func NewCrs(crsType string) (*Crs, error) {
	if crsType == "" {	
		return nil, fmt.Errorf("crsType cannot be empty")
	}
    return &Crs{
        Type:       crsType,
		Properties: make(map[string]string),
    }, nil
}