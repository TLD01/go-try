package repository

import (
	"context"
	"time"

	"aerowatch.com/api/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "users"

type UserRepository struct {
	*repository.MongoRepository[UserEntity, *UserEntity]
}

func NewUserRepository(db *mongo.Database) (*UserRepository, error) {
	if db == nil {
		return nil, repository.ErrDbRequired
	}
	mongoRepo, err := repository.NewMongoRepository[UserEntity](db, collectionName)
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		MongoRepository: mongoRepo,
	}, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*UserEntity, error) {
	filter := map[string]any{
		"email": email,
	}
	return r.FindOne(ctx, filter)
}

func (r *UserRepository) UpdateLastSignIn(ctx context.Context, id string, lastSignin time.Time) (*UserEntity, error) {
	update := map[string]any{
		"lastSignIn": lastSignin,
	}
	return r.Patch(ctx, r.ToID(id), update)
}
