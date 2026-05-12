package book

import (
	"context"
	bookEntity "go-library/internal/entity"
	"go-library/internal/usecase/author"
)

// Service - useCase

type service struct {
	storage       Storage
	authorService author.Service
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, request CreateBookDTO) *bookEntity.Book {
	authorModel := s.authorService.GetByUUID(ctx, request.AuthorUUID)
	if authorModel {
		return nil
	}

	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *bookEntity.Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit int, offset int) []*bookEntity.Book {
	return s.storage.GetAll(limit, offset)
}
