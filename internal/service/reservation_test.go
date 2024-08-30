package service

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qara-qurt/booking_service/internal/repository/postgres"
	"github.com/qara-qurt/booking_service/model"
	"github.com/stretchr/testify/assert"
)

// setupDB establishes a connection to the PostgreSQL database and returns a pool
func setupDB() (*pgxpool.Pool, func()) {
	connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return db, func() {
		db.Close()
	}
}

// clearReservations clears reservations in the test database
func clearReservations(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), "DELETE FROM reservations")
	return err
}

// TestCreateReservation tests the Create method of the ReservationService
func TestCreateReservation(t *testing.T) {
	// Setup the database connection and defer the teardown
	db, teardown := setupDB()
	defer teardown()

	// Create a new reservation layers
	repo := postgres.NewReservationRepo(db)
	service := NewReservationService(repo)

	// Clear any existing reservations
	clearReservations(db)

	req := &model.ReservationRequest{
		RoomID:    "room1",
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(2 * time.Hour),
	}

	// Create a new reservation
	err := service.Create(req)
	assert.NoError(t, err)

	// Verify that the reservation was created successfully
	reservations, err := repo.GetReservationByRoom("room1")
	assert.NoError(t, err)
	assert.Len(t, reservations, 1)
	assert.Equal(t, req.RoomID, reservations[0].RoomID)

	// Clear reservations after the test
	clearReservations(db)
}

// TestCreateReservationWithOverlap tests creating a reservation that overlaps with an existing one
func TestCreateReservationWithOverlap(t *testing.T) {
	// Setup the database connection and defer the teardown
	db, teardown := setupDB()
	defer teardown()

	// Create all layers
	repo := postgres.NewReservationRepo(db)
	service := NewReservationService(repo)
	clearReservations(db)

	// 2024-08-12 13:00:00 UTC
	date := time.Date(2024, 8, 12, 13, 0, 0, 0, time.UTC)

	req1 := &model.ReservationRequest{
		RoomID:    "room1",
		StartTime: date,
		EndTime:   date.Add(2 * time.Hour),
	}
	err := service.Create(req1)
	assert.NoError(t, err)

	// Create a second reservation that overlaps with the first one
	req2 := &model.ReservationRequest{
		RoomID:    "room1",
		StartTime: date,
		EndTime:   date.Add(1 * time.Hour),
	}

	err = service.Create(req2)
	assert.Error(t, err)
	assert.Equal(t, model.ErrRoomAlreadyReserved, err)

	// Verify that only the first reservation exists
	reservations, err := repo.GetReservationByRoom("room1")
	assert.NoError(t, err)
	assert.Len(t, reservations, 1)

	// Clear reservations after the test
	clearReservations(db)
}

// TestConcurrentReservations tests creating reservations concurrently
func TestConcurrentReservations(t *testing.T) {
	//Setup the database connection and defer the teardown
	db, teardown := setupDB()
	defer teardown()

	repo := postgres.NewReservationRepo(db)
	service := NewReservationService(repo)

	// Clear any existing reservations to ensure a clean state
	err := clearReservations(db)
	assert.NoError(t, err)

	wg := sync.WaitGroup{}
	// 2024-08-12 13:00:00 UTC
	date := time.Date(2024, 8, 12, 13, 0, 0, 0, time.UTC)

	err = service.Create(&model.ReservationRequest{
		RoomID:    "room1",
		StartTime: date.Add(1 * time.Hour),
		EndTime:   date.Add(2 * time.Hour),
	})
	assert.NoError(t, err)

	createReservation := func() {
		req := &model.ReservationRequest{
			RoomID:    "room1",
			StartTime: date.Add(1 * time.Hour),
			EndTime:   date.Add(2 * time.Hour),
		}
		err := service.Create(req)
		assert.Error(t, err)
	}

	// Launch multiple goroutines to create reservations concurrently
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			createReservation()
		}()
	}
	wg.Wait()

	// Verify that only one reservation was successfully created
	reservations, err := repo.GetReservationByRoom("room1")
	assert.NoError(t, err)
	assert.Len(t, reservations, 1)
	clearReservations(db)
}
