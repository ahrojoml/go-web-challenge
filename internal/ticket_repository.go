package internal

import "context"

type TicketRepository interface {
	GetAll(ctx context.Context) (map[int]TicketAttributes, error)
	GetTicketsByDestinationCountry(ctx context.Context, country string) (map[int]TicketAttributes, error)
}
