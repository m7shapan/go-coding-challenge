package services

import (
	"challenge/repositories"
	"context"
)

type KeyService interface {
	IsValid(context.Context, string) (bool, error)
}

type keyService struct {
	keyRepository repositories.KeyRepository
}

func NewKeyService(r repositories.KeyRepository) keyService {
	return keyService{
		keyRepository: r,
	}
}

func (k keyService) IsValid(ctx context.Context, key string) (bool, error) {
	key, err := k.keyRepository.GetKey(ctx, key)

	if err != nil {
		return false, err
	}

	return true, err
}
