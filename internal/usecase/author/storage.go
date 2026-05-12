// storage - repository

package author

import "go-library/internal/entity"

type Storage interface {
	GetOne(uuid string) *entity.Author
	GetAll(limit int, offset int) []*entity.Author
	Create(author *entity.Author) *entity.Author
	Delete(author *entity.Author) error
}
