package service

import (
	"errors"

	"booking-api/internal/models"
	"booking-api/internal/repository"
)

type RoomService struct {
	repo *repository.RoomRepository
}

func NewRoomService(repo *repository.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) CreateRoom(room *models.Room) error {
	if room.Class == "" {
		return errors.New("class is required")
	}

	if room.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	err := s.repo.Create(room)
	if err != nil {
		return err
	}

	return nil
}

func (s *RoomService) GetAllRooms() ([]models.Room, error) {
	return s.repo.GetAll()
}

func (s *RoomService) GetRoomByID(id uint) (*models.Room, error) {
	return s.repo.GetByID(id)
}

func (s *RoomService) DeleteRoom(id uint) error {
	return s.repo.Delete(id)
}
