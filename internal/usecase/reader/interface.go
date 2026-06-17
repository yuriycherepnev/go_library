package reader

import (
	"go-library/internal/domain/reader"
)

type Service interface {
	GetById(id int) (*reader.Reader, error)
	GetAll() ([]reader.Reader, error)
	Create(request CreateReaderDTO) (*reader.Reader, error)
	Update(request UpdateReaderDTO) (*reader.Reader, error)
	Delete(request DeleteReaderDTO) (*reader.Reader, error)
}
