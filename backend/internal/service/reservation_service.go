package service

import (
	"context"
	"hotel-booking-backend/internal/models"
	"hotel-booking-backend/internal/repository"
)

type ReservationService struct {
	repo     *repository.ReservationRepository
	roomRepo *repository.RoomRepository
}

func NewReservationService(repo *repository.ReservationRepository, roomRepo *repository.RoomRepository) *ReservationService {
	return &ReservationService{repo: repo, roomRepo: roomRepo}
}

func (s *ReservationService) GetAll(ctx context.Context) ([]models.Reservation, error) {
	return s.repo.GetAll(ctx)
}

func (s *ReservationService) Create(ctx context.Context, rv *models.Reservation) (int, error) {
	roomID, err := s.repo.Create(ctx, rv)
	if err != nil {
		return 0, err
	}

	if err := s.roomRepo.UpdateRoomStatus(ctx, roomID); err != nil {
		return 0, err
	}
	return roomID, nil
}

func (s *ReservationService) GetByID(ctx context.Context, id int) (*models.Reservation, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ReservationService) GetByIDClient(ctx context.Context, idClient int) ([]models.Reservation, error) {
	return s.repo.GetByIDClient(ctx, idClient)
}

func (s *ReservationService) GetByIDRoom(ctx context.Context, idRoom int) ([]models.Reservation, error) {
	return s.repo.GetByIDRoom(ctx, idRoom)
}

func (s *ReservationService) Update(ctx context.Context, id int, c *models.Reservation) (*models.Reservation, error) {
	old, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	updated, err := s.repo.Update(ctx, id, c)
	if err != nil {
		return nil, err
	}
	// Update both old and new rooms
	if old.IDRoom != updated.IDRoom {
		err = s.roomRepo.UpdateRoomStatus(ctx, old.IDRoom)
		if err != nil {
			return nil, err
		}

	}
	err = s.roomRepo.UpdateRoomStatus(ctx, updated.IDRoom)

	if err != nil {
		return nil, err
	}
	return updated, nil

}

func (s *ReservationService) Delete(ctx context.Context, id int) (int, error) {
	roomID, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	if err := s.roomRepo.UpdateRoomStatus(ctx, roomID); err != nil {
		return 0, err
	}
	return roomID, nil
}

func (s *ReservationService) GetAllWithDetails(ctx context.Context) ([]models.BookingDetails, error) {
	return s.repo.GetAllWithDetails(ctx)
}
