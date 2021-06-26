package services

import (
	"challenge/models"
	"challenge/repositories"
	"context"
)

type FactService interface {
	GetFacts(context.Context, *models.Filters) ([]models.Fact, int64, error)
}

type factService struct {
	factRepository repositories.FactRepository
}

func NewFactService(r repositories.FactRepository) factService {
	return factService{
		factRepository: r,
	}
}

func (f factService) GetFacts(ctx context.Context, filter *models.Filters) ([]models.Fact, int64, error) {
	return f.factRepository.GetFacts(ctx, filter)
}
