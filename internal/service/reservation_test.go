package service

import (
	"sync"
	"testing"
	"time"

	"github.com/qara-qurt/booking_service/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository for testing
type MockReservationRepo struct {
	mock.Mock
}

func (m *MockReservationRepo) GetReservationByRoom(roomID string) ([]model.Reservation, error) {
	args := m.Called(roomID)
	return args.Get(0).([]model.Reservation), args.Error(1)
}

func (m *MockReservationRepo) Create(data *model.ReservationRequest) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestCreateReservation_NoOverlap(t *testing.T) {
	mockRepo := new(MockReservationRepo)
	service := NewReservationService(mockRepo)

	// Setup mock expectations
	mockRepo.On("GetReservationByRoom", "room1").Return([]model.Reservation{}, nil)
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Create reservation
	req := &model.ReservationRequest{
		RoomID:    "room1",
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(2 * time.Hour),
	}
	err := service.Create(req)

	// Assert expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateReservation_Overlap(t *testing.T) {
	mockRepo := new(MockReservationRepo)
	service := NewReservationService(mockRepo)

	// Setup mock expectations
	existingReservations := []model.Reservation{
		{
			RoomID:    "room1",
			StartTime: time.Now().Add(1 * time.Hour),
			EndTime:   time.Now().Add(3 * time.Hour),
		},
	}
	mockRepo.On("GetReservationByRoom", "room1").Return(existingReservations, nil)

	// Create reservation that overlaps
	req := &model.ReservationRequest{
		RoomID:    "room1",
		StartTime: time.Now().Add(2 * time.Hour),
		EndTime:   time.Now().Add(4 * time.Hour),
	}
	err := service.Create(req)

	// Assert expectations
	assert.Equal(t, model.ErrRoomAlreadyReserved, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateReservation_Concurrent(t *testing.T) {
	mockRepo := new(MockReservationRepo)
	service := NewReservationService(mockRepo)

	// Setup mock expectations
	existingReservations := []model.Reservation{
		{
			RoomID:    "room1",
			StartTime: time.Now().Add(1 * time.Hour),
			EndTime:   time.Now().Add(2 * time.Hour),
		},
	}
	mockRepo.On("GetReservationByRoom", "room1").Return(existingReservations, nil)
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Create reservations concurrently
	req1 := &model.ReservationRequest{
		RoomID:    "room1",
		StartTime: time.Now().Add(2 * time.Hour),
		EndTime:   time.Now().Add(3 * time.Hour),
	}

	req2 := &model.ReservationRequest{
		RoomID:    "room1",
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(2 * time.Hour),
	}

	var err1, err2 error
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err1 = service.Create(req1)
	}()

	go func() {
		defer wg.Done()
		err2 = service.Create(req2)
	}()

	wg.Wait()

	// Assert expectations
	assert.NoError(t, err1)
	assert.Equal(t, model.ErrRoomAlreadyReserved, err2)
	mockRepo.AssertExpectations(t)
}
