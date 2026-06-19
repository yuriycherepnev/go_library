// порт - интерфейс репозитория

package book

type Repository interface {
	GetOne(id int) (*Book, error)
	GetAll() ([]Book, error)
	Create(title string, idAuthor int) (*Book, error)
	Update(id int, title string, idAuthor *int) error
	Delete(id int) error
	Borrow(bookID int, readerID int) error
	Return(bookID int) error
}
