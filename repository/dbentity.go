package repository

import (
	"time"
	"log/slog"
	"aerowatch.com/api/common"
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
	IDField        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAtField time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAtField time.Time     `json:"updatedAt" bson:"updatedAt"`
}

func Create(p common.Persisted) DBEntity {
	var id bson.ObjectID
	if p.ID != "" {
		if objID, err := bson.ObjectIDFromHex(p.ID); err == nil {
			id = objID
		}else {
			slog.Error("Invalid ID format", "id", p.ID, "error", err)
			panic("Invalid ID format: " + p.ID)
		}
	}

	return DBEntity{
		IDField:        id,
		CreatedAtField: p.CreatedAt,
		UpdatedAtField: p.UpdatedAt,
	}
}

func (e DBEntity) ToPersisted() common.Persisted {
	return common.Persisted{
		ID:        e.ID().Hex(),
		CreatedAt: e.CreatedAt(),
		UpdatedAt: e.UpdatedAt(),
	}
}

func (e DBEntity) ID() bson.ObjectID {
	return e.IDField
}

func (e DBEntity) CreatedAt() time.Time { 
	return e.CreatedAtField
}

func (e DBEntity) UpdatedAt() time.Time {
	return e.UpdatedAtField
}

func (e *DBEntity) SetID(id bson.ObjectID) {
	e.IDField = id
}

func (e *DBEntity) SetCreatedAt(t time.Time) {
	e.CreatedAtField = t
}

func (e *DBEntity) SetUpdatedAt(t time.Time) {
	e.UpdatedAtField = t
}
