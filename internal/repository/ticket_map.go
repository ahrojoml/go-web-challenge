package repository

import (
	"context"
	"go-web-challenge/internal"
	"go-web-challenge/internal/loader"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(dbFile string, lastId int) (*RepositoryTicketMap, error) {
	dbLoader := loader.NewLoaderTicketCSV(dbFile)
	dbMap, err := dbLoader.Load()
	if err != nil {
		return nil, err
	}

	return &RepositoryTicketMap{
		dbMap:  dbMap,
		lastId: lastId,
	}, nil
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	dbMap map[int]internal.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) GetAll(ctx context.Context) (map[int]internal.TicketAttributes, error) {
	// create a copy of the map
	t := make(map[int]internal.TicketAttributes, len(r.dbMap))
	for k, v := range r.dbMap {
		t[k] = v
	}

	return t, nil
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error) {
	// create a copy of the map
	t := make(map[int]internal.TicketAttributes)
	for k, v := range r.dbMap {
		if v.Country == country {
			t[k] = v
		}
	}

	return t, nil
}
