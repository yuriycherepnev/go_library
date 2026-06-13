// порт - интерфейс репозитория

package author

type Repository interface {
	GetOne(id int) (*Author, error)
	GetAll() ([]Author, error)
	Create(name string) (*Author, error)
	Update(id int, name string) error
	Delete(id int) error
}
