package repository

import (
	"github.com/qara-qurt/booking_service/config"
	"github.com/qara-qurt/booking_service/internal/repository/postgres"
	"github.com/qara-qurt/booking_service/model"
)

type IReservationRepo interface {
	Create(data *model.ReservationRequest) error
	GetReservationByRoom(roomID string) ([]model.Reservation, error)
}

type Repository struct {
	Reservation IReservationRepo
}

func New(config *config.Config) (*Repository, error) {

	postgresDB, err := postgres.NewPostgres(config)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Reservation: postgres.NewReservationRepo(postgresDB.DB),
	}, nil
}
