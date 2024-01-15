package service

import (
	"context"
	"go-web-challenge/internal"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTickets(ctx context.Context) (int, error) {
	tickets, err := s.rp.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

// GetTotalTicketsByDestinationCountry returns the total number of tickets by destination country
func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(ctx context.Context, country string) (int, error) {
	if country == "" {
		return 0, NewInvalidCountryError("no country provided")
	}

	tickets, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (float64, error) {
	if country == "" {
		return 0, NewInvalidCountryError("no country provided")
	}

	allTickets, err := s.rp.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	ticketsByCountry, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		return 0, err
	}

	percentage := float64(len(ticketsByCountry)) / float64(len(allTickets))

	return percentage, nil
}
