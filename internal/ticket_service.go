package internal

import "context"

type ServiceTicket interface {

	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalTickets(ctx context.Context) (int, error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	GetTicketsAmountByDestinationCountry(ctx context.Context, country string) (int, error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (float64, error)
}
