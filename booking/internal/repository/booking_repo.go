package repository

import (
	"errors"

	"booking-api/internal/models"

	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Create(booking *models.Booking) error {
	var allBookings []models.Booking
	err := r.db.Where("room_id = ?", booking.RoomID).Find(&allBookings).Error
	if err != nil {
		return err
	}

	for _, b := range allBookings {

		if booking.StartTime.Before(b.EndTime) && booking.EndTime.After(b.StartTime) {
			return errors.New("room is busy")
		}
	}

	err = r.db.Create(booking).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) GetByID(id uint) (*models.Booking, error) {
	var booking models.Booking

	err := r.db.First(&booking, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}

	return &booking, nil
}

func (r *BookingRepository) GetByUser(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking

	err := r.db.Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}

	return bookings, nil
}
