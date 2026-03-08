package repository

import (
	"context"

	"aerowatch.com/api/common"
	"aerowatch.com/api/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "aero_event_log"

type EventRepository struct {
	*repository.MongoRepository[EventEntity, *EventEntity]
}

func NewEventRepository(db *mongo.Database) (*EventRepository, error) {
	if db == nil {
		return nil, repository.ErrDbRequired
	}
	mongoRepo, err := repository.NewMongoRepository[EventEntity](db, collectionName)
	if err != nil {
		return nil, err
	}
	return &EventRepository{
		MongoRepository: mongoRepo,
	}, nil
}

func (r *EventRepository) Search(ctx context.Context, icaoAddress int, timeWindow common.TimeWindow) (*[]EventEntity, error) {
	filter := map[string]any{
		"vehicleMessage.icaoAddress": icaoAddress,
		"timestamp": map[string]any{
			"$gte": timeWindow.Start,
			"$lte": timeWindow.End,
		},
	}
	return r.FindMany(ctx, filter)
}
