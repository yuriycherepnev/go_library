package book

import (
	book2 "go-library/internal/entity"
	"go-library/internal/usecase/book"
)

type bookStorage struct {
}

func NewStorage() book.Storage {
	return &bookStorage{}
}

func (bs *bookStorage) GetOne(uuid string) *book2.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit int, offset int) []*book2.Book {
	return nil
}
func (bs *bookStorage) Create(book *book2.Book) *book2.Book {
	return nil
}
func (bs *bookStorage) Delete(book *book2.Book) error {
	return nil
}
