package services

import (
	"challenge/models"
	"challenge/repositories"
	"context"
)

type FactService interface {
	GetFacts(context.Context) ([]models.Fact, error)
}

type factService struct {
	factRepository repositories.FactRepository
}

func NewFactService(r repositories.FactRepository) factService {
	return factService{
		factRepository: r,
	}
}

func (f factService) GetFacts(ctx context.Context) ([]models.Fact, error) {
	return f.factRepository.GetFacts(ctx)
}
