package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qara-qurt/booking_service/model"
)

// CreateReservation creates a new reservation
func (h *Handler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation model.ReservationRequest

	// Decode the request body into the reservation struct
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the reservation start_time should be before end_time
	if reservation.StartTime.After(reservation.EndTime) {
		http.Error(w, "start time connot be adter end", http.StatusBadRequest)
		return
	}

	// Create the reservation service business logic
	err := h.service.Reservation.Create(&reservation)
	if err != nil {

		// Check if the error is a room already reserved error
		if errors.Is(err, model.ErrRoomAlreadyReserved) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ok
	w.WriteHeader(http.StatusCreated)
}

// GetReservationByRoom retrieves all reservations for a given room
func (h *Handler) GetReservationByRoom(w http.ResponseWriter, r *http.Request) {

	// Get the roomID from the URL
	roomID := chi.URLParam(r, "roomID")
	if roomID == "" {
		http.Error(w, "roomID is required", http.StatusBadRequest)
		return
	}

	// Get the reservations for the room
	reservations, err := h.service.Reservation.GetReservationByRoom(roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ok
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}
