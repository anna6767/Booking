package service

import (
	"errors"
	"time"

	"booking-api/internal/models"
	"booking-api/internal/repository"
)

type BookingService struct {
	repo     *repository.BookingRepository
	roomRepo *repository.RoomRepository
}

func NewBookingService(
	repo *repository.BookingRepository,
	roomRepo *repository.RoomRepository,
) *BookingService {
	return &BookingService{
		repo:     repo,
		roomRepo: roomRepo,
	}
}

func (s *BookingService) BookRoom(req models.BookingRequest) (*models.Booking, error) {
	if !req.StartTime.Before(req.EndTime) {
		return nil, errors.New("start_time must be before end_time")
	}

	if req.StartTime.Before(time.Now()) {
		return nil, errors.New("cannot book in the past")
	}

	_, err := s.roomRepo.GetByID(req.RoomID)
	if err != nil {
		return nil, errors.New("room not found")
	}

	booking := models.Booking{
		UserID:    req.UserID,
		RoomID:    req.RoomID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	err = s.repo.Create(&booking)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (s *BookingService) GetBookingByID(id uint) (*models.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *BookingService) GetBookingsByUser(userID uint) ([]models.Booking, error) {
	return s.repo.GetByUser(userID)
}
