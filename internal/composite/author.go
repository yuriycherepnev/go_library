package composite

import (
	"go-library/internal/adapters"
	"go-library/internal/adapters/api/author"
	authorStorage "go-library/internal/adapters/db/author"
	author2 "go-library/internal/usecase/author"
)

type AuthorComposite struct {
	Storage author2.Storage
	Service author2.Service
	Handler adapters.Handler
}

func NewAuthorComposite(composite *MysqlComposite) (*AuthorComposite, error) {
	storage := authorStorage.NewStorage(composite.db)
	service := author2.NewService(storage)
	handler := author.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil

}
