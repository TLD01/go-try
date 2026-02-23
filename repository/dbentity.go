package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DBEntity struct {
	Id        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}
