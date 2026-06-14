package author

import (
	"database/sql"
	authorDomain "go-library/internal/domain/author"
)

type service struct {
	authorStorage authorDomain.Repository
}

func NewService(storage authorDomain.Repository) Service {
	return &service{authorStorage: storage}
}

func (s *service) GetById(id int) (*authorDomain.Author, error) {
	author, err := s.authorStorage.GetOne(id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *service) GetAll() ([]authorDomain.Author, error) {
	authors, err := s.authorStorage.GetAll()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (s *service) Create(request CreateAuthorDTO) (*authorDomain.Author, error) {
	author, err := s.authorStorage.Create(request.Name)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *service) Update(request UpdateAuthorDTO) (*authorDomain.Author, error) {
	err := s.authorStorage.Update(request.Id, request.Name)
	if err != nil {
		return nil, err
	}
	return s.authorStorage.GetOne(request.Id)
}

func (s *service) Delete(request DeleteAuthorDTO) (*authorDomain.Author, error) {
	a, err := s.authorStorage.GetOne(request.Id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, sql.ErrNoRows
	}
	err = s.authorStorage.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return a, nil
}
