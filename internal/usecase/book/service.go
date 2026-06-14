// Service - useCase

package book

import (
	domain "go-library/internal/domain/book"
)

type service struct {
	storage domain.Repository
}

func NewService(storage domain.Repository) Service {
	return &service{storage: storage}
}

func (s *service) GetById(id int) (*domain.Book, error) {
	book, err := s.storage.GetOne(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *service) GetAll() ([]domain.Book, error) {
	books, err := s.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *service) Create(request CreateBookDTO) (*domain.Book, error) {
	book, err := s.storage.Create(request.Title, request.IdAuthor)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *service) Update(request UpdateBookDTO) (*domain.Book, error) {
	err := s.storage.Update(request.Id, request.Title)
	if err != nil {
		return nil, err
	}

	return s.storage.GetOne(request.Id)
}

func (s *service) Delete(request DeleteBookDTO) (*domain.Book, error) {
	b, err := s.storage.GetOne(request.Id)
	if err != nil {
		return nil, err
	}

	err = s.storage.Delete(request.Id)
	if err != nil {
		return nil, err
	}

	return b, nil
}
