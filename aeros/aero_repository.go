package aeros

import (
	"context"

	"aerowatch.com/api/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "aeros"

type AerosRepository struct {
	repository.MongoRepository
}

func NewAerosRepository(db *mongo.Database) (*AerosRepository, error) {
	if db == nil {
		return nil, repository.ErrDbRequired
	}
	mongoRepo, err := repository.NewMongoRepository(db, collectionName)
	if err != nil {
		return nil, err
	}
	return &AerosRepository{
		MongoRepository: *mongoRepo,
	}, nil
}


func (r *AerosRepository) Patch2(ctx context.Context, aero *Aero) (*Aero, error) {
	return r.Save(ctx, aero)
}

