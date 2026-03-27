package service

import (
	"context"
	"errors"
	"hotel-booking-backend/internal/models"
	"hotel-booking-backend/internal/repository"
)

type RoomService struct {
	repo *repository.RoomRepository
}

func NewRoomService(repo *repository.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) GetAll(ctx context.Context) ([]models.Room, error) {
	return s.repo.GetAll(ctx)
}

func (s *RoomService) Create(ctx context.Context, rm *models.Room) error {
	if rm.RoomNo == "" {
		return errors.New("room number is required")
	}

	return s.repo.Create(ctx, rm)
}

func (s *RoomService) GetByID(ctx context.Context, id int) (*models.Room, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *RoomService) Update(ctx context.Context, id int, rm *models.Room) (*models.Room, error) {
	if rm.RoomNo == "" {
		return nil, errors.New("room number is required")
	}

	return s.repo.Update(ctx, id, rm)
}

func (s *RoomService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
