// порт - интерфейс репозитория

package reader

type Repository interface {
	GetOne(id int) (*Reader, error)
	GetAll() ([]Reader, error)
	Create(name string) (*Reader, error)
	Update(id int, name string) error
	Delete(id int) error
}
