package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"aerowatch.com/api/aeros"
	aero_repository "aerowatch.com/api/aeros/repository"
	v1aeros "aerowatch.com/api/api/aeros/v1"
	v1events "aerowatch.com/api/api/events/v1"
	"aerowatch.com/api/api/filters"
	"aerowatch.com/api/config/logging"
	"aerowatch.com/api/events"
	events_repository "aerowatch.com/api/events/repository"
	"aerowatch.com/api/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			provideMongoConnection,
			provideMongoDatabase,
			aero_repository.NewAerosRepository,
			events_repository.NewEventRepository,
			aeros.NewAeroService,
			events.NewEventsService,
			v1aeros.NewAeroController,
			v1events.NewEventController,
		),
		fx.Invoke(startHTTPServer),
	).Run()
}

func provideMongoConnection() (repository.MongoConnection, error) {
	uri := os.Getenv("MONGO_URI")
	database := os.Getenv("MONGO_DATABASE")
	return repository.NewMongoConnection(uri, database)
}

func provideMongoDatabase(lc fx.Lifecycle, conn repository.MongoConnection) (*mongo.Database, error) {
	client := repository.NewMongoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Setup(ctx); err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			// TODO: add MongoClient.Disconnect when available
			return nil
		},
	})

	return client.Database()
}

func startHTTPServer(lc fx.Lifecycle, aeroCtrl *v1aeros.AeroController, eventCtrl *v1events.EventController) {
	logger := logging.GetLogger("http")
	mux := http.NewServeMux()
	aeroCtrl.RegisterRoutes(mux)
	eventCtrl.RegisterRoutes(mux)
	handler := filters.CorrelationIDMiddleware(filters.LoggingMiddleware(mux))
	server := &http.Server{Addr: ":8080", Handler: handler}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("HTTP server starting", "addr", server.Addr)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("HTTP server error", "err", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("HTTP server shutting down")
			return server.Shutdown(ctx)
		},
	})
}
