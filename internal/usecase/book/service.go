// Service - useCase

package book

import (
	"go-library/internal/domain"
	authorDomain "go-library/internal/domain/author"
	bookDomain "go-library/internal/domain/book"
)

type service struct {
	bookStorage   bookDomain.Repository
	authorStorage authorDomain.Repository
}

func NewService(
	bookStorage bookDomain.Repository,
	authorStorage authorDomain.Repository,
) Service {
	return &service{
		bookStorage:   bookStorage,
		authorStorage: authorStorage,
	}
}

func (s *service) GetById(id int) (*bookDomain.Book, error) {
	book, err := s.bookStorage.GetOne(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *service) GetAll() ([]bookDomain.Book, error) {
	books, err := s.bookStorage.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *service) Create(request CreateBookDTO) (*bookDomain.Book, error) {
	a, err := s.authorStorage.GetOne(request.IdAuthor)
	if err != nil {
		return nil, err
	}

	if a == nil {
		return nil, domain.ErrAuthorNotFound
	}

	b, err := s.bookStorage.Create(request.Title, request.IdAuthor)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *service) Update(request UpdateBookDTO) (*bookDomain.Book, error) {
	if request.IdAuthor != nil {
		a, err := s.authorStorage.GetOne(*request.IdAuthor)
		if err != nil {
			return nil, err
		}

		if a == nil {
			return nil, domain.ErrAuthorNotFound
		}
	}

	b, err := s.bookStorage.GetOne(request.Id)
	if err != nil {
		return nil, err
	}

	if b == nil {
		return nil, domain.ErrBookNotFound
	}

	err = s.bookStorage.Update(request.Id, request.Title, request.IdAuthor)
	if err != nil {
		return nil, err
	}

	return s.bookStorage.GetOne(request.Id)
}

func (s *service) Delete(request DeleteBookDTO) (*bookDomain.Book, error) {
	b, err := s.bookStorage.GetOne(request.Id)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, domain.ErrBookNotFound
	}

	err = s.bookStorage.Delete(request.Id)
	if err != nil {
		return nil, err
	}

	return b, nil
}
