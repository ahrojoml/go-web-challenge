package handler

import (
	"context"
	"go-web-challenge/internal"
	"go-web-challenge/platform/web/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	sv internal.ServiceTicket
}

func NewTicketHandler(sv internal.ServiceTicket) *TicketHandler {
	return &TicketHandler{
		sv: sv,
	}
}

func (h *TicketHandler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Text(w, http.StatusOK, "ok")
	}
}

func (h *TicketHandler) GetTicketsByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "dest")
		if country == "" {
			response.Error(w, http.StatusBadRequest, "must provide a country")
			return
		}

		ctx := context.Background()

		tickets, err := h.sv.GetTicketsAmountByDestinationCountry(ctx, country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "could not get tickets")
			return
		}

		response.JSON(w, http.StatusOK, tickets)
	}
}

func (h *TicketHandler) GetPercentageTicketsByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "dest")
		if country == "" {
			response.Error(w, http.StatusBadRequest, "must provide a country")
			return
		}

		ctx := context.Background()

		tickets, err := h.sv.GetPercentageTicketsByDestinationCountry(ctx, country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "could not get tickets")
			return
		}

		response.JSON(w, http.StatusOK, tickets)
	}
}
