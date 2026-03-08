package repository

import (
	"context"

	"aerowatch.com/api/common"
	"aerowatch.com/api/geolocation"
	"aerowatch.com/api/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "aeros"

type AerosRepository struct {
	*repository.MongoRepository[AeroEntity, *AeroEntity]
}

func NewAerosRepository(db *mongo.Database) (*AerosRepository, error) {
	if db == nil {
		return nil, repository.ErrDbRequired
	}
	mongoRepo, err := repository.NewMongoRepository[AeroEntity, *AeroEntity](db, collectionName)
	if err != nil {
		return nil, err
	}

	return &AerosRepository{
		MongoRepository: mongoRepo,
	}, nil
}

func (r *AerosRepository) FindByIcao(ctx context.Context, icao string) (*AeroEntity, error) {
	filter := map[string]any{
		"icao": icao,
	}
	return r.FindOne(ctx, filter)
}

func (r *AerosRepository) Search(ctx context.Context, boundary  geolocation.BoundingBox, timeWindow common.TimeWindow) (*[]AeroEntity, error) {
	filter := map[string]any{
		"lastPosition": map[string]any{
			"$geoWithin": map[string]any{
				"$box": [][]float64{
					{boundary.SW.Longitude, boundary.SW.Latitude},
					{boundary.NE.Longitude, boundary.NE.Latitude},
				},
			},
		},
		"lastSeen": map[string]any{
			"$gte": timeWindow.Start,
			"$lte": timeWindow.End,
		},
	}
	return r.FindMany(ctx, filter)
}
