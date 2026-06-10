package author

import (
	domain "go-library/internal/domain/author"
)

type service struct {
	storage domain.Repository
}

func NewService(storage domain.Repository) Service {
	return &service{storage: storage}
}

func (s *service) GetById(id int) (*domain.Author, error) {
	author, err := s.storage.GetOne(id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *service) GetAll() ([]domain.Author, error) {
	authors, err := s.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (s *service) Create(request CreateAuthorDTO) (*domain.Author, error) {
	return nil, nil
}

func (s *service) Update(request UpdateAuthorDTO) (*domain.Author, error) {
	return nil, nil
}

func (s *service) Delete(request DeleteAuthorDto) (*domain.Author, error) {
	return nil, nil
}
