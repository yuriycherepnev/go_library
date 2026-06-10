// это порт

package author

type Repository interface {
	GetOne(uuid string) *Author
	GetAll(limit int, offset int) []*Author
	Create(author *Author) *Author
	Delete(author *Author) error
}
