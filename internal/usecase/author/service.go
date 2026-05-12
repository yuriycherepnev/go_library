package author

import (
	"context"
	"go-library/internal/entity"
)

// Service - useCase

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, request CreateAuthorDTO) *entity.Author {
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *entity.Author {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit int, offset int) []*entity.Author {
	return s.storage.GetAll(limit, offset)
}
