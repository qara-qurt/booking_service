package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qara-qurt/booking_service/model"
)

func (h *Handler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation model.ReservationRequest

	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reservation.StartTime.After(reservation.EndTime) {
		http.Error(w, "start time connot be adter end", http.StatusBadRequest)
		return
	}

	err := h.service.Reservation.Create(&reservation)
	if err != nil {

		if errors.Is(err, model.ErrRoomAlreadyReserved) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetReservationByRoom(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomID")
	if roomID == "" {
		http.Error(w, "roomID is required", http.StatusBadRequest)
		return
	}

	reservations, err := h.service.Reservation.GetReservationByRoom(roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}
