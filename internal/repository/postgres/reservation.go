package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qara-qurt/booking_service/model"
)

type reservationRepo struct {
	db *pgxpool.Pool
}

func NewReservationRepo(db *pgxpool.Pool) *reservationRepo {
	return &reservationRepo{
		db: db,
	}
}

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

	var reservations []model.Reservation
	for rows.Next() {
		var reservation model.Reservation
		err = rows.Scan(&reservation.ID, &reservation.RoomID, &reservation.StartTime, &reservation.EndTime)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
