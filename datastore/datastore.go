package datastore

import "go-api/models"

type Datastore interface {
	UpdateLocation(userID string, loc *models.Location) error
	GetLocation(userID string, max int) ([]*models.Location, error)
	Delete(userID string) error
}

type inMemoryStore struct {
	locations map[string][]*models.Location
}

func (in *inMemoryStore) UpdateLocation(userID string, loc *models.Location) error {
	if locs, ok := in.locations[userID]; ok {
		locs = append(locs, loc)
	} else {
		in.locations[userID] = []*models.Location{loc}
	}
	return nil
}

func (in *inMemoryStore) GetLocation(userID string, max int) ([]*models.Location, error) {
	if locs, ok := in.locations[userID]; ok {
		if len(locs) < max {
			return locs, nil
		}
		return locs[:max], nil
	}

	return nil, nil
}

func (in *inMemoryStore) Delete(userID string) error {
	if _, ok := in.locations[userID]; ok {
		delete(in.locations, userID)
	}
	return nil
}

func NewInMemeryStore() Datastore {
	return &inMemoryStore{
		locations: make(map[string][]*models.Location),
	}
}
