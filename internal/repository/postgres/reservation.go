package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qara-qurt/booking_service/model"
)

type reservationRepo struct {
	db *pgxpool.Pool
}

// NewReservationRepo creates a new reservationRepo instance
func NewReservationRepo(db *pgxpool.Pool) *reservationRepo {
	return &reservationRepo{
		db: db,
	}
}

// Create creates a new reservation
func (r *reservationRepo) Create(data *model.ReservationRequest) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO reservations (room_id, start_time, end_time)
		VALUES ($1, $2, $3)
	`, data.RoomID, data.StartTime, data.EndTime)
	if err != nil {
		return err
	}

	return nil
}

// GetReservationByRoom retrieves all reservations for a given room
func (r *reservationRepo) GetReservationByRoom(roomID string) ([]model.Reservation, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT id, room_id, start_time, end_time
		FROM reservations
		WHERE room_id = $1
	`, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the rows into a slice of reservations
	var reservations []model.Reservation
	for rows.Next() {
		var reservation model.Reservation
		err = rows.Scan(&reservation.ID, &reservation.RoomID, &reservation.StartTime, &reservation.EndTime)
		if err != nil {
			return nil, err
		}

		// Append the reservation to the slice
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
