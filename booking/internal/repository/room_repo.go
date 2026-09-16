package repository

import (
	"booking-api/internal/models"
	"errors"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}
func (r *RoomRepository) Create(room *models.Room) error {
	err := r.db.Create(room).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *RoomRepository) GetAll() ([]models.Room, error) {
	var rooms []models.Room

	err := r.db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
func (r *RoomRepository) GetByID(id uint) (*models.Room, error) {
	var room models.Room

	err := r.db.First(&room, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}
func (r *RoomRepository) Delete(id uint) error {
	var room models.Room
	err := r.db.First(&room, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("room not found")
		}
		return err
	}
	err = r.db.Delete(&room).Error
	if err != nil {
		return err
	}
	return nil
}
