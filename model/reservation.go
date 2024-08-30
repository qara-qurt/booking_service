package model

import "time"

// ReservationRequest represents a reservation request
type ReservationRequest struct {
	RoomID    string    `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// Reservation represents a reservation
type Reservation struct {
	ID        int       `json:"id"`
	RoomID    string    `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
