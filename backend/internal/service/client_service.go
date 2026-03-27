package service

import (
	"context"
	"errors"

	"clientmanager/internal/models"
	"clientmanager/internal/repository"
)

type ClientService struct {
	repo *repository.ClientRepository
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) GetAll(ctx context.Context) ([]models.Client, error) {
	return s.repo.GetAll(ctx)
}

func (s *ClientService) Create(ctx context.Context, c *models.Client) error {
	if c.ClientCode == "" {
		return errors.New("client code is required")
	}
	if c.CompanyName == "" {
		return errors.New("company name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}

	return s.repo.Create(ctx, c)
}

func (s *ClientService) GetByID(ctx context.Context, id int) (*models.Client, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ClientService) Update(ctx context.Context, id int, c *models.Client) (*models.Client, error) {
	if c.CompanyName == "" {
		return nil, errors.New("company name is required")
	}

	return s.repo.Update(ctx, id, c)
}

func (s *ClientService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
