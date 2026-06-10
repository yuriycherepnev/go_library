package author

import (
	"go-library/internal/domain/author"
)

type Service interface {
	GetById(id int) (*author.Author, error)
	GetAll() ([]*author.Author, error)
	Create(request CreateAuthorDTO) (*author.Author, error)
	Update(request UpdateAuthorDTO) (*author.Author, error)
	Delete(request DeleteAuthorDto) (*author.Author, error)
}
