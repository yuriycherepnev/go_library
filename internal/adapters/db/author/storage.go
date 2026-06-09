package author

import (
	"database/sql"
	"go-library/internal/entity"
	"go-library/internal/usecase/author"
)

type authorStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) author.Storage {
	return &authorStorage{
		db: db,
	}
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
