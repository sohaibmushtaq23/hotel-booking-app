package service

import (
	"context"
	"hotel-booking-backend/internal/models"
	"hotel-booking-backend/internal/repository"
)

type RoomReservationService struct {
	repo *repository.RoomReservationRepository
}

func NewRoomReservationService(repo *repository.RoomReservationRepository) *RoomReservationService {
	return &RoomReservationService{repo: repo}
}

func (s *RoomReservationService) GetByIDReservation(ctx context.Context, idReservation int) ([]models.RoomReservation, error) {
	return s.repo.GetByIDReservation(ctx, idReservation)
}

func (s *RoomReservationService) GetByIDRoom(ctx context.Context, idRoom int) ([]models.RoomReservation, error) {
	return s.repo.GetByIDRoom(ctx, idRoom)
}

func (s *RoomReservationService) Create(ctx context.Context, rv *models.RoomReservation) error {
	return s.repo.Create(ctx, rv)
}

func (s *RoomReservationService) GetByID(ctx context.Context, id int) (*models.RoomReservation, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *RoomReservationService) Update(ctx context.Context, id int, c *models.RoomReservation) (*models.RoomReservation, error) {
	return s.repo.Update(ctx, id, c)
}

func (s *RoomReservationService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
