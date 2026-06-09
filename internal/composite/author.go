package composite

import (
	"go-library/internal/adapters"
	"go-library/internal/adapters/api/author"
	authorStorage "go-library/internal/adapters/db/author"
	author2 "go-library/internal/usecase/author"
	"go-library/internal/usecase/book"
)

type AuthorComposite struct {
	Storage book.Storage
	Service author2.Service
	Handler adapters.Handler
}

func NewAuthorComposite() (*AuthorComposite, error) {
	storage := authorStorage.NewStorage()
	service := author2.NewService(storage)
	handler := author.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil

}
