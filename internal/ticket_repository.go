package internal

import "context"

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// GetAll returns all the tickets
	GetAll(ctx context.Context) (map[int]TicketAttributes, error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	GetTicketsByDestinationCountry(ctx context.Context, country string) (map[int]TicketAttributes, error)
}
