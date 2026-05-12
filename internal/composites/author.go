package composites

import (
	"ca-template/internal/adapters"
	"ca-template/internal/adapters/api/author"
	authorStorage "ca-template/internal/adapters/db/author"
	authorService "ca-template/internal/domain/author"
)

type AuthorComposite struct {
	Storage authorService.Storage
	Service authorService.Service
	Handler adapters.Handler
}

func NewAuthorComposite() (*AuthorComposite, error) {
	storage := authorStorage.NewStorage()
	service := authorService.NewService(storage)
	handler := author.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil

}
