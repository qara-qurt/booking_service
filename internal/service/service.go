package service

import (
	"github.com/qara-qurt/booking_service/internal/repository"
	"github.com/qara-qurt/booking_service/model"
)

// IReservation represents the reservation service interface
type IReservation interface {
	Create(data *model.ReservationRequest) error
	GetReservationByRoom(roomID string) ([]model.Reservation, error)
}

type Service struct {
	Reservation IReservation
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Reservation: NewReservationService(repo.Reservation),
	}
}
