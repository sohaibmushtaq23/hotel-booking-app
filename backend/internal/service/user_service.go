package service

import (
	"context"
	"errors"
	"hotel-booking-backend/internal/models"
	"hotel-booking-backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) Create(ctx context.Context, u *models.User) error {
	if u.UserName == "" {
		return errors.New("user name is required")
	}

	if u.Password == "" {
		return errors.New("a strong password is required")
	}

	if u.UserRole == "" {
		return errors.New("user role is required")
	}

	return s.repo.Create(ctx, u)
}

func (s *UserService) GetByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id int, u *models.User) (*models.User, error) {
	if u.UserName == "" {
		return nil, errors.New("user name is required")
	}

	if u.Password == "" {
		return nil, errors.New("a strong password is required")
	}

	if u.UserRole == "" {
		return nil, errors.New("user role is required")
	}

	return s.repo.Update(ctx, id, u)
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
