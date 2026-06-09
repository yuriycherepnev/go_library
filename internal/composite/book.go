package composite

import (
	"go-library/internal/adapters"
	"go-library/internal/adapters/api/book"
	bookStorage "go-library/internal/adapters/db/book"
	book2 "go-library/internal/usecase/book"
)

type BookComposite struct {
	Storage book2.Storage
	Service book2.Service
	Handler adapters.Handler
}

func NewBookComposite(composite *MysqlComposite) (*BookComposite, error) {
	storage := bookStorage.NewStorage(composite.db)
	service := book2.NewService(storage)
	handler := book.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil

}
