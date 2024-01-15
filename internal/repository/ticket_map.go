package repository

import (
	"context"
	"go-web-challenge/internal"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(db map[int]internal.TicketAttributes, lastId int) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		dbMap:  db,
		lastId: lastId,
	}
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
	tickets := make(map[int]internal.TicketAttributes, len(r.dbMap))
	for k, v := range r.dbMap {
		tickets[k] = v
	}

	return tickets, nil
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error) {
	// create a copy of the map
	tickets := make(map[int]internal.TicketAttributes)
	for k, v := range r.dbMap {
		if v.Country == country {
			tickets[k] = v
		}
	}

	return tickets, nil
}
