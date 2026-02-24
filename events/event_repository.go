package events

import (
	"aerowatch.com/api/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "aero_event_log"


type EventRepository struct {
	*repository.MongoRepository[*EventEntity]
}

func NewEventRepository(db *mongo.Database) (*EventRepository, error) {
	if db == nil {
		return nil, repository.ErrDbRequired
	}
	mongoRepo, err := repository.NewMongoRepository[*EventEntity](db, collectionName)
	if err != nil {
		return nil, err
	}
	return &EventRepository{
		MongoRepository: mongoRepo,
	}, nil
}