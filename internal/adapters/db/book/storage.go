package book

import (
	"database/sql"
	domain "go-library/internal/domain/book"
)

type bookStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) domain.Repository {
	return &bookStorage{
		db: db,
	}
}

func (bs *bookStorage) GetOne(uuid string) *domain.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit int, offset int) []*domain.Book {
	return nil
}
func (bs *bookStorage) Create(book *domain.Book) *domain.Book {
	return nil
}
func (bs *bookStorage) Delete(book *domain.Book) error {
	return nil
}
