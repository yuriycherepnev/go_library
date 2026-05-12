// storage - repository

package book

import "go-library/internal/entity"

type Storage interface {
	GetOne(uuid string) *entity.Book
	GetAll(limit int, offset int) []*entity.Book
	Create(book *entity.Book) *entity.Book
	Delete(book *entity.Book) error
}
