// это порт
// repository

package book

type Repository interface {
	GetOne(uuid string) *Book
	GetAll(limit int, offset int) []*Book
	Create(book *Book) *Book
	Delete(book *Book) error
}
