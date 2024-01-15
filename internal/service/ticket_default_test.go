package service_test

import (
	"context"
	"fmt"
	"go-web-challenge/internal"
	"go-web-challenge/internal/repository"
	"go-web-challenge/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		ctx := context.Background()
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalTickets(ctx)

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

func TestGetTotalTicketsByDestinationCountry(t *testing.T) {
	testCases := []struct {
		name        string
		tickets     map[int]internal.TicketAttributes
		country     string
		expect      int
		expectError error
	}{
		{
			name:        "error retrieving tickets",
			country:     "",
			expectError: service.NewInvalidCountryError("no country provided"),
		}, {
			name:    "success",
			country: "Chile",
			tickets: map[int]internal.TicketAttributes{
				1: {
					Name:    "test",
					Email:   "test@example.com",
					Country: "Chile",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "test2",
					Email:   "test2@example.com",
					Country: "Argentina",
					Hour:    "10:00",
					Price:   100,
				},
			},
			expect: 1,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprint(idx, tC.name), func(t *testing.T) {
			rp := repository.NewRepositoryTicketMock()

			rp.FuncGetTicketsByDestinationCountry = func(country string) (map[int]internal.TicketAttributes, error) {
				tickets := make(map[int]internal.TicketAttributes)
				for k, v := range tC.tickets {
					if v.Country == country {
						tickets[k] = v
					}
				}
				return tickets, nil
			}

			sv := service.NewServiceTicketDefault(rp)

			total, err := sv.GetTicketsAmountByDestinationCountry(context.Background(), tC.country)

			if tC.expectError != nil {
				require.ErrorAs(t, err, &tC.expectError)
			} else {
				require.Nil(t, err)
				require.Equal(t, tC.expect, total)
			}
		})
	}

}

func TestGetPercentageTicketsByDestinationCountry(t *testing.T) {
	testCases := []struct {
		name        string
		tickets     map[int]internal.TicketAttributes
		country     string
		expect      float64
		expectError error
	}{
		{
			name:        "error retrieving tickets",
			country:     "",
			expectError: service.NewInvalidCountryError("no country provided"),
		}, {
			name:    "success",
			country: "Chile",
			tickets: map[int]internal.TicketAttributes{
				1: {
					Name:    "test",
					Email:   "test@example.com",
					Country: "Chile",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "test2",
					Email:   "test2@example.com",
					Country: "Argentina",
					Hour:    "10:00",
					Price:   100,
				},
			},
			expect: 1.0 / 2.0,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprint(idx, tC.name), func(t *testing.T) {
			rp := repository.NewRepositoryTicketMock()
			rp.FuncGet = func() (map[int]internal.TicketAttributes, error) {
				tickets := tC.tickets
				return tickets, nil
			}

			rp.FuncGetTicketsByDestinationCountry = func(country string) (map[int]internal.TicketAttributes, error) {
				tickets := make(map[int]internal.TicketAttributes)
				for k, v := range tC.tickets {
					if v.Country == country {
						tickets[k] = v
					}
				}
				return tickets, nil
			}

			sv := service.NewServiceTicketDefault(rp)

			total, err := sv.GetPercentageTicketsByDestinationCountry(context.Background(), tC.country)

			if tC.expectError != nil {
				require.ErrorAs(t, err, &tC.expectError)
			} else {
				require.Nil(t, err)
				require.Equal(t, tC.expect, total)
			}
		})
	}

}
