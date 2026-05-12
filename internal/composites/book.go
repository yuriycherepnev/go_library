package composites

import (
	"ca-template/internal/adapters"
	"ca-template/internal/adapters/api/book"
	bookStorage "ca-template/internal/adapters/db/book"
	bookService "ca-template/internal/domain/book"
)

type BookComposite struct {
	Storage bookService.Storage
	Service bookService.Service
	Handler adapters.Handler
}

func NewBookComposite() (*BookComposite, error) {
	storage := bookStorage.NewStorage()
	service := bookService.NewService(storage)
	handler := book.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil

}
