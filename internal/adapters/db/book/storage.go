package book

import (
	"database/sql"
	book2 "go-library/internal/entity"
	"go-library/internal/usecase/book"
)

type bookStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) book.Storage {
	return &bookStorage{
		db: db,
	}
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
