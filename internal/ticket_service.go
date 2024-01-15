package internal

type TicketService interface {
	GetTotalTickets() (total int, err error)
	GetTicketsAmountByDestinationCountry() (amount int, err error)
}
