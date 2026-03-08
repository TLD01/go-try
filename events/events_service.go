package events

import (
	"context"
	"strconv"

	"aerowatch.com/api/aeros"
	"aerowatch.com/api/common"
	events_repository "aerowatch.com/api/events/repository"
	repository "aerowatch.com/api/repository"
)

type EventsService struct {
	Repository  *events_repository.EventRepository
	AeroService *aeros.AeroService
}

func NewEventsService(eventsRepository *events_repository.EventRepository, aeroService *aeros.AeroService) (*EventsService, error) {
	if eventsRepository == nil {
		return nil, repository.ErrRepoRequired

	}
	
	return &EventsService{
		Repository:  eventsRepository,
		AeroService: aeroService,
	}, nil
}

func (e *EventsService) Search(ctx context.Context, icaoAddress string, timeWindow common.TimeWindow) (*[]Event, error) {
	icaoAddressInt, err := strconv.Atoi(icaoAddress)	
	if err != nil {
		return nil, err
	}
	eventEntities, err := e.Repository.Search(ctx, icaoAddressInt, timeWindow)
	if err != nil {
		return nil, err
	}
	result := make([]Event, len(*eventEntities))
	for i, event := range *eventEntities {
		result[i] = *toEvent(&event)
	}
	return &result, nil
}


func toEvent(eventEntity *events_repository.EventEntity) *Event {
	return &Event{
		Persisted: eventEntity.DBEntity.ToPersisted(),
		Source : eventEntity.Source,
		Timestamp: eventEntity.Timestamp,
		VehicleMessage: eventEntity.VehicleMessage,
	}
}
