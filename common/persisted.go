package common

import "time"
type Persisted struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`	
}