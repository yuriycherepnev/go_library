package author

import (
	"database/sql"
	domain "go-library/internal/domain/author"
)

type authorStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) domain.Repository {
	return &authorStorage{
		db: db,
	}
}

func (as *authorStorage) GetOne(uuid string) *domain.Author {
	return nil
}
func (as *authorStorage) GetAll(limit int, offset int) []*domain.Author {
	return nil
}
func (as *authorStorage) Create(author *domain.Author) *domain.Author {
	return author
}
func (as *authorStorage) Delete(author *domain.Author) error {
	return nil
}
