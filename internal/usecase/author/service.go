package author

import (
	"database/sql"
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
	author, err := s.storage.Create(request.Name)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *service) Update(request UpdateAuthorDTO) (*domain.Author, error) {
	err := s.storage.Update(request.Id, request.Name)
	if err != nil {
		return nil, err
	}
	return s.storage.GetOne(request.Id)
}

func (s *service) Delete(request DeleteAuthorDTO) (*domain.Author, error) {
	a, err := s.storage.GetOne(request.Id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, sql.ErrNoRows
	}
	err = s.storage.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return a, nil
}
