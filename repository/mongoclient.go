package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	ErrMongoClientNotConnected = errors.New("mongo client is not connected")
	ErrMongoDatabaseNotSet     = errors.New("mongo database is not set")
	ErrMongoCollectionName     = errors.New("collection name is required")
	ErrMongoClientNotInit      = errors.New("mongo client singleton is not initialized")
)

var (
	mongoClientOnce     sync.Once
	mongoClientInstance *MongoClient
	mongoClientInitErr  error
)

type MongoConnection struct {
	URI         string
	Database    string
	PingTimeout time.Duration
}

func NewMongoConnection(uri, database string) (MongoConnection, error) {
	if uri == "" {
		return MongoConnection{}, errors.New("mongo uri is required")
	}

	if database == "" {
		return MongoConnection{}, errors.New("mongo database is required")
	}

	return MongoConnection{
		URI:         uri,
		Database:    database,
		PingTimeout: 10 * time.Second,
	}, nil
}

type MongoClient struct {
	connection MongoConnection
	client     *mongo.Client
	database   *mongo.Database
}

func NewMongoClient(connection MongoConnection) *MongoClient {
	return &MongoClient{connection: connection}
}

func InitMongoClient(ctx context.Context, uri, database string) (*MongoClient, error) {
	mongoClientOnce.Do(func() {
		connection, err := NewMongoConnection(uri, database)
		if err != nil {
			mongoClientInitErr = err
			return
		}

		client := NewMongoClient(connection)
		if err := client.Setup(ctx); err != nil {
			mongoClientInitErr = err
			return
		}

		mongoClientInstance = client
	})

	if mongoClientInitErr != nil {
		return nil, mongoClientInitErr
	}

	if mongoClientInstance == nil {
		return nil, ErrMongoClientNotInit
	}

	return mongoClientInstance, nil
}

func GetMongoClient() (*MongoClient, error) {
	if mongoClientInstance == nil {
		return nil, ErrMongoClientNotInit
	}

	return mongoClientInstance, nil
}

func (mc *MongoClient) Setup(ctx context.Context) error {
	client, err := mongo.Connect(options.Client().ApplyURI(mc.connection.URI))
	if err != nil {
		return fmt.Errorf("connect mongo client: %w", err)
	}

	pingCtx := ctx
	var cancel context.CancelFunc
	if mc.connection.PingTimeout > 0 {
		pingCtx, cancel = context.WithTimeout(ctx, mc.connection.PingTimeout)
		defer cancel()
	}

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		_ = client.Disconnect(context.Background())
		return fmt.Errorf("ping mongo deployment: %w", err)
	}

	mc.client = client
	mc.database = client.Database(mc.connection.Database)
	return nil
}

func (mc *MongoClient) Database() (*mongo.Database, error) {
	if mc.client == nil {
		return nil, ErrMongoClientNotConnected
	}

	if mc.database == nil {
		return nil, ErrMongoDatabaseNotSet
	}

	return mc.database, nil
}

func (mc *MongoClient) Collection(name string) (*mongo.Collection, error) {
	if name == "" {
		return nil, ErrMongoCollectionName
	}

	db, err := mc.Database()
	if err != nil {
		return nil, err
	}

	return db.Collection(name), nil
}

func (mc *MongoClient) Disconnect(ctx context.Context) error {
	if mc.client == nil {
		return nil
	}

	err := mc.client.Disconnect(ctx)
	mc.client = nil
	mc.database = nil
	return err
}
