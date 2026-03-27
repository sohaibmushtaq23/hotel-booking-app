package service

import (
	"context"
	"errors"

	"clientmanager/internal/models"
	"clientmanager/internal/repository"
)

type ClientContactService struct {
	repo *repository.ClientContactRepository
}

func NewClientContactService(repo *repository.ClientContactRepository) *ClientContactService {
	return &ClientContactService{repo: repo}
}

func (s *ClientContactService) GetByClientID(ctx context.Context, idClient int) ([]models.ClientContact, error) {
	return s.repo.GetByClientID(ctx, idClient)
}

func (s *ClientContactService) GetByID(ctx context.Context, id int) (*models.ClientContact, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ClientContactService) Create(ctx context.Context, contact *models.ClientContact) error {
	if contact.FirstName == "" {
		return errors.New("first name is required")
	}
	if contact.LastName == "" {
		return errors.New("last name is required")
	}

	return s.repo.Create(ctx, contact)
}

func (s *ClientContactService) Update(ctx context.Context, id int, contact *models.ClientContact) error {
	if contact.FirstName == "" {
		return errors.New("first name is required")
	}
	if contact.LastName == "" {
		return errors.New("last name is required")
	}

	return s.repo.Update(ctx, id, contact)
}

func (s *ClientContactService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
