package book

import (
	"context"
	domain "go-library/internal/domain/book"
)

// Service - useCase

type service struct {
	storage domain.Repository
}

func NewService(storage domain.Repository) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, request CreateBookDTO) *domain.Book {
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *domain.Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit int, offset int) []*domain.Book {
	return s.storage.GetAll(limit, offset)
}
