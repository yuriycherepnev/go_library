package reader

import (
	"go-library/internal/domain"
	readerDomain "go-library/internal/domain/reader"
)

type service struct {
	readerStorage readerDomain.Repository
}

func NewService(storage readerDomain.Repository) Service {
	return &service{readerStorage: storage}
}

func (s *service) GetById(id int) (*readerDomain.Reader, error) {
	author, err := s.readerStorage.GetOne(id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *service) GetAll() ([]readerDomain.Reader, error) {
	authors, err := s.readerStorage.GetAll()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (s *service) Create(request CreateReaderDTO) (*readerDomain.Reader, error) {
	author, err := s.readerStorage.Create(request.Name)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *service) Update(request UpdateReaderDTO) (*readerDomain.Reader, error) {
	a, err := s.readerStorage.GetOne(request.Id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, domain.ErrReaderNotFound
	}
	err = s.readerStorage.Update(request.Id, request.Name)
	if err != nil {
		return nil, err
	}
	return s.readerStorage.GetOne(request.Id)
}

func (s *service) Delete(request DeleteReaderDTO) (*readerDomain.Reader, error) {
	a, err := s.readerStorage.GetOne(request.Id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, domain.ErrReaderNotFound
	}
	err = s.readerStorage.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return a, nil
}
