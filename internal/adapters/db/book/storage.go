package book

import "ca-template/internal/domain/book"

type bookStorage struct {
}

func NewStorage() book.Storage {
	return &bookStorage{}
}

func (bs *bookStorage) GetOne(uuid string) *book.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit int, offset int) []*book.Book {
	return nil
}
func (bs *bookStorage) Create(book *book.Book) *book.Book {
	return nil
}
func (bs *bookStorage) Delete(book *book.Book) error {
	return nil
}
