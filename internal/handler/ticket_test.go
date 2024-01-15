package handler_test

import (
	"context"
	"fmt"
	"go-web-challenge/internal"
	"go-web-challenge/internal/handler"
	"go-web-challenge/internal/repository"
	"go-web-challenge/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetTicketsByCountry(t *testing.T) {
	testCases := []struct {
		name         string
		tickets      map[int]internal.TicketAttributes
		country      string
		expectBody   string
		expectedCode int
	}{
		{
			name:         "missing parameter", // I dont know how it would be possible to get here, but just in case
			country:      "",
			expectedCode: http.StatusBadRequest,
			expectBody:   "must provide a country",
		},
		{
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
			expectBody:   "1",
			expectedCode: http.StatusOK,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprint(idx, tC.name), func(t *testing.T) {
			repo := repository.NewRepositoryTicketMap(tC.tickets, 2)
			sv := service.NewServiceTicketDefault(repo)
			hd := handler.NewTicketHandler(sv)

			req := httptest.NewRequest("GET", "/tickets/country/Chile", nil)
			routeContext := chi.NewRouteContext()
			routeContext.URLParams.Add("dest", tC.country)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
			res := httptest.NewRecorder()
			hd.GetTicketsByCountry()(res, req)

			require.Equal(t, tC.expectedCode, res.Code)
			if tC.expectedCode < 300 {
				require.Equal(t, tC.expectBody, res.Body.String())
			}
		})
	}
}

func TestGetPercentageTicketsByCountry(t *testing.T) {
	testCases := []struct {
		name         string
		tickets      map[int]internal.TicketAttributes
		country      string
		expectBody   string
		expectedCode int
	}{
		{
			name:         "missing parameter", // I dont know how it would be possible to get here, but just in case
			country:      "",
			expectedCode: http.StatusBadRequest,
			expectBody:   "must provide a country",
		},
		{
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
			expectBody:   "0.5",
			expectedCode: http.StatusOK,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprint(idx, tC.name), func(t *testing.T) {
			repo := repository.NewRepositoryTicketMap(tC.tickets, 2)
			sv := service.NewServiceTicketDefault(repo)
			hd := handler.NewTicketHandler(sv)

			req := httptest.NewRequest("GET", "/tickets/country/percentage/Chile", nil)
			routeContext := chi.NewRouteContext()
			routeContext.URLParams.Add("dest", tC.country)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
			res := httptest.NewRecorder()
			hd.GetPercentageTicketsByCountry()(res, req)

			require.Equal(t, tC.expectedCode, res.Code)
			if tC.expectedCode < 300 {
				require.Equal(t, tC.expectBody, res.Body.String())
			}
		})
	}
}
