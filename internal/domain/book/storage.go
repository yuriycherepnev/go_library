package book

// Storage - repository

type Storage interface {
	GetOne(uuid string) *Book
	GetAll(limit int, offset int) []*Book
	Create(book *Book) *Book
	Delete(book *Book) error
}
