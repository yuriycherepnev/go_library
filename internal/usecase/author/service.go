package author

import (
	"context"
	domain "go-library/internal/domain/author"
)

type service struct {
	storage domain.Repository
}

func NewService(storage domain.Repository) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, request CreateAuthorDTO) *domain.Author {
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *domain.Author {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit int, offset int) []domain.Author {
	return s.storage.GetAll()
}
