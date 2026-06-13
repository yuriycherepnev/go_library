// adapter - реализация репозитория

package book

import (
	"database/sql"
	"go-library/internal/domain/book"
)

type bookStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) book.Repository {
	return &bookStorage{
		db: db,
	}
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
