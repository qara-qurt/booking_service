package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/qara-qurt/booking_service/internal/service"
)

// Handler represents the API handler
type Handler struct {
	router  *chi.Mux
	service *service.Service
}

// New creates a new Handler instance
func New(service *service.Service, router *chi.Mux) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

// RegisterRoutes registers the API routes
func (h *Handler) RegisterRoutes() {
	h.router.Post("/reservation", h.CreateReservation)
	h.router.Get("/reservation/{roomID}", h.GetReservationByRoom)
}
