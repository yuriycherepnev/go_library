package book

import (
	"ca-template/internal/domain/author"
	"context"
)

// Service - useCase

type service struct {
	storage       Storage
	authorService author.Service
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, request CreateBookDTO) *Book {
	authorModel := s.authorService.GetByUUID(ctx, request.AuthorUUID)
	if authorModel {
		return nil
	}

	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit int, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
