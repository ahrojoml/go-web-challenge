package handler_test

import (
	"fmt"
	"go-web-challenge/internal"
	"go-web-challenge/internal/service"
	"testing"
)

func TestGetTicketsByCountry(t *testing.T) {
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

		})
	}
}
