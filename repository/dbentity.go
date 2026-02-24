package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Entity interface {
	ID() bson.ObjectID
	CreatedAt() time.Time
	UpdatedAt() time.Time

	// Setters
	SetID(id bson.ObjectID)
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
}

type DBEntity struct {
	Id        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	createdAt time.Time     `json:"createdAt" bson:"createdAt"`
	updatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

func (e *DBEntity) ID() bson.ObjectID {
	return e.Id
}

func (e *DBEntity) CreatedAt() time.Time {
	return e.createdAt
}

func (e *DBEntity) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *DBEntity) SetID(id bson.ObjectID) {
	e.Id = id
}

func (e *DBEntity) SetCreatedAt(t time.Time) {
	e.createdAt = t
}

func (e *DBEntity) SetUpdatedAt(t time.Time) {
	e.updatedAt = t
}
