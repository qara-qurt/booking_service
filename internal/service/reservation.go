package service

import (
	"github.com/qara-qurt/booking_service/internal/repository"
	"github.com/qara-qurt/booking_service/model"
)

type reservationService struct {
	repo repository.IReservationRepo
}

// Create creates a new reservation service
func NewReservationService(repo repository.IReservationRepo) *reservationService {
	return &reservationService{
		repo: repo,
	}
}

// Create creates a new reservation
func (s *reservationService) Create(data *model.ReservationRequest) error {

	// Get all reservations for the room by roomID
	existsReservation, err := s.repo.GetReservationByRoom(data.RoomID)
	if err != nil {
		return err
	}

	// Check if the room is already reserved
	for _, r := range existsReservation {
		if r.StartTime.Before(data.EndTime) && r.EndTime.After(data.StartTime) {
			return model.ErrRoomAlreadyReserved
		}
	}

	// Create the reservation
	return s.repo.Create(data)
}

// GetReservationByRoom retrieves all reservations for a given room
func (s *reservationService) GetReservationByRoom(roomID string) ([]model.Reservation, error) {
	return s.repo.GetReservationByRoom(roomID)
}
