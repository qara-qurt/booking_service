package service

import (
	"github.com/qara-qurt/booking_service/internal/repository"
	"github.com/qara-qurt/booking_service/model"
)

type reservationService struct {
	repo repository.IReservationRepo
}

func NewReservationService(repo repository.IReservationRepo) *reservationService {
	return &reservationService{
		repo: repo,
	}
}

func (s *reservationService) Create(data *model.ReservationRequest) error {
	existsReservation, err := s.repo.GetReservationByRoom(data.RoomID)
	if err != nil {
		return err
	}

	for _, r := range existsReservation {
		if r.StartTime.Before(data.EndTime) && r.EndTime.After(data.StartTime) {
			return model.ErrRoomAlreadyReserved
		}
	}

	return s.repo.Create(data)
}

func (s *reservationService) GetReservationByRoom(roomID string) ([]model.Reservation, error) {
	return s.repo.GetReservationByRoom(roomID)
}
