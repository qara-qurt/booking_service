package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/qara-qurt/booking_service/internal/service"
)

type Handler struct {
	router  *chi.Mux
	service *service.Service
}

func New(service *service.Service, router *chi.Mux) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

func (h *Handler) RegisterRoutes() {
	h.router.Post("/reservation", h.CreateReservation)
	h.router.Get("/reservation/{roomID}", h.GetReservationByRoom)
}
