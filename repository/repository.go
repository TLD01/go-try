package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var ErrNotFound = errors.New("entity not found")
var ErrDbRequired = errors.New("mongo database is required")

type Repository[T Entity] interface {
	Save(ctx context.Context, entity *T) (*T, error)
	Find(ctx context.Context, id bson.ObjectID) (*T, error)
	Patch(ctx context.Context, id bson.ObjectID, vals map[string]any) (*T, error)
	Delete(ctx context.Context, id bson.ObjectID) error
}

type MongoRepository[T Entity] struct {
	collection *mongo.Collection
}

func NewMongoRepository[T Entity](db *mongo.Database, collectionName string) (*MongoRepository[T], error) {

	if collectionName == "" {
		return nil, ErrMongoCollectionName
	}

	if db == nil {
		return nil, ErrDbRequired
	}
	collection := db.Collection(collectionName)
	return &MongoRepository[T]{collection: collection}, nil
}

func (r *MongoRepository[T]) Find(ctx context.Context, id bson.ObjectID) (*T, error) {
	var entity T
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

func (r *MongoRepository[T]) Delete(ctx context.Context, id bson.ObjectID) error {
	deleteResult, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *MongoRepository[T]) Patch(ctx context.Context, id bson.ObjectID, vals map[string]any) (*T, error) {
	if _, ok := vals["_id"]; ok {
		return nil, errors.New("_id field cannot be modified")
	}
	if _, ok := vals["createdAt"]; ok {
		return nil, errors.New("createdAt field cannot be modified")
	}

	vals["updatedAt"] = time.Now()
	update := bson.M{"$set": vals}
	var entity T
	err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&entity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

func (r *MongoRepository[T]) Save(ctx context.Context, entity *T) (*T, error) {
	if (*entity).ID().IsZero() {
		// inserted, err := r.insert(ctx, entity)
		_, err := r.insert(ctx, entity)
		if err != nil {
			return nil, err
		}
		// return r.Find(ctx, inserted.InsertedID.(bson.ObjectID))
		return entity, nil
	}
	return r.replace(ctx, entity)
}

func (r *MongoRepository[T]) replace(ctx context.Context, entity *T) (*T, error) {
	(*entity).SetUpdatedAt(time.Now())

	updateResult, err := r.collection.ReplaceOne(ctx, bson.M{"_id": (*entity).ID()}, entity, options.Replace().SetUpsert(false))
	if err != nil {
		return nil, err
	}
	if updateResult.MatchedCount == 0 {
		return nil, ErrNotFound
	}
	return entity, nil
}

func (r *MongoRepository[T]) insert(ctx context.Context, entity *T) (*mongo.InsertOneResult, error) {
	(*entity).SetCreatedAt(time.Now())
	(*entity).SetUpdatedAt(time.Now())
	(*entity).SetID(bson.NewObjectID())

	insertOneResult, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	return insertOneResult, nil
}
