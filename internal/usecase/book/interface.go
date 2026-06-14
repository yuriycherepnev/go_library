package book

// interface бизнес логики лежит там где он используется - рядом с handler
import (
	"go-library/internal/domain/book"
)

type Service interface {
	GetById(id int) (*book.Book, error)
	GetAll() ([]book.Book, error)
	Create(request CreateBookDTO) (*book.Book, error)
	Update(request UpdateBookDTO) (*book.Book, error)
	Delete(request DeleteBookDTO) (*book.Book, error)
}
