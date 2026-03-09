package aeros

import (
	"context"
	"strconv"
	"time"

	aero_repository "aerowatch.com/api/aeros/repository"
	"aerowatch.com/api/common"
	"aerowatch.com/api/geolocation"
	"aerowatch.com/api/receivers/messages"
	repository "aerowatch.com/api/repository"
)

type AeroService struct {
	Repository *aero_repository.AerosRepository
}

func NewAeroService(aeroRepo *aero_repository.AerosRepository) (*AeroService, error) {
	if aeroRepo == nil {
		return nil, repository.ErrRepoRequired
	}
	return &AeroService{
		Repository: aeroRepo,
	}, nil
}

func (as *AeroService) SaveLatestEvent(ctx context.Context, vehicleMessage *messages.AdsbVehicleMessage, timestamp time.Time) (*Aero, error) {
	aeroEntity, err := as.Repository.FindByIcao(ctx, strconv.Itoa(vehicleMessage.IcaoAddress))
	if err != nil {
		return nil, err
	}
	if aeroEntity == nil {
		return nil, repository.ErrNotFound
	}

	updates := map[string]any{
		"lastSeen":     timestamp,
		"lastPosition": vehicleMessage.Position(),
		"lastMessage":  vehicleMessage,
	}
	aeroEntity, err = as.Repository.Patch(ctx, aeroEntity.ID(), updates)
	if err != nil {
		return nil, err
	}
	return toAero(aeroEntity), nil
	
}

func (as *AeroService) Get(ctx context.Context, icao string) (*Aero, error) {
	aeroEntity, err := as.Repository.FindByIcao(ctx, icao)
	if err != nil {
		return nil, err
	}	
	return toAero(aeroEntity), nil
}


func (as *AeroService) Search(ctx context.Context, boundary geolocation.BoundingBox, timeWindow common.TimeWindow) (*[]Aero, error) {
	aeroEntities, err := as.Repository.Search(ctx, boundary, timeWindow)
	if err != nil {
		return nil, err
	}
	result := make([]Aero, len(*aeroEntities))
	for i, entity := range *aeroEntities {
		result[i] = *toAero(&entity)
	}
	return &result, nil
}


func create(aero *Aero) *aero_repository.AeroEntity {
	return &aero_repository.AeroEntity{
		DBEntity: repository.Create(aero.Persisted),
		Callsign:     aero.Callsign,
		IcaoAddress:  aero.IcaoAddress,
		Model:        aero.Model,
		LastSeen:     aero.LastSeen,
		LastPosition: aero.LastPosition,
		LastMessage:  aero.LastMessage,
	}
}

func toAero(a *aero_repository.AeroEntity) *Aero {
	return &Aero{
		Persisted: a.DBEntity.ToPersisted(),
		Callsign:     a.Callsign,
		IcaoAddress:  a.IcaoAddress,
		Model:        a.Model,
		LastSeen:     a.LastSeen,
		LastPosition: a.LastPosition,
		LastMessage:  a.LastMessage,
	}
}
