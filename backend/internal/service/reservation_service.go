package service

import (
	"context"
	"hotel-booking-backend/internal/models"
	"hotel-booking-backend/internal/repository"
)

type ReservationService struct {
	repo *repository.ReservationRepository
}

func NewReservationService(repo *repository.ReservationRepository) *ReservationService {
	return &ReservationService{repo: repo}
}

func (s *ReservationService) GetAll(ctx context.Context) ([]models.Reservation, error) {
	return s.repo.GetAll(ctx)
}

func (s *ReservationService) Create(ctx context.Context, rv *models.Reservation) error {
	return s.repo.Create(ctx, rv)
}

func (s *ReservationService) GetByID(ctx context.Context, id int) (*models.Reservation, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ReservationService) Update(ctx context.Context, id int, c *models.Reservation) (*models.Reservation, error) {
	return s.repo.Update(ctx, id, c)
}

func (s *ReservationService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
