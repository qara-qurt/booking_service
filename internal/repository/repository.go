package repository

import (
	"github.com/qara-qurt/booking_service/config"
	"github.com/qara-qurt/booking_service/internal/repository/postgres"
	"github.com/qara-qurt/booking_service/model"
)

// IReservationRepo represents the reservation repository interface
type IReservationRepo interface {
	Create(data *model.ReservationRequest) error
	GetReservationByRoom(roomID string) ([]model.Reservation, error)
}

// Repository represents the repository
type Repository struct {
	Reservation IReservationRepo
}

func New(config *config.Config) (*Repository, error) {
	// Create a new PostgreSQL connection
	postgresDB, err := postgres.NewPostgres(config)
	if err != nil {
		return nil, err
	}

	// Return the repository and give each repository the PostgreSQL connection
	return &Repository{
		Reservation: postgres.NewReservationRepo(postgresDB.DB),
	}, nil
}
