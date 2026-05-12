package author

import (
	"go-library/internal/entity"
	"go-library/internal/usecase/book"
)

type authorStorage struct {
}

func NewStorage() book.Storage {
	return &authorStorage{}
}

func (as *authorStorage) GetOne(uuid string) *entity.Author {
	return nil
}
func (as *authorStorage) GetAll(limit int, offset int) []*entity.Author {
	return nil
}
func (as *authorStorage) Create(author *entity.Author) *entity.Author {
	return author
}
func (as *authorStorage) Delete(author *entity.Author) error {
	return nil
}
