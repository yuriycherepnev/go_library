package composite

import (
	adapter "go-library/internal/adapters/db/author"
	domain "go-library/internal/domain/author"
	handler "go-library/internal/handlers"
	authorHandler "go-library/internal/handlers/api/author"
	useCase "go-library/internal/usecase/author"
)

type AuthorComposite struct {
	Storage domain.Repository
	Service useCase.Service
	Handler handler.Handler
}

func NewAuthorComposite(composite *MysqlComposite) (*AuthorComposite, error) {
	storage := adapter.NewStorage(composite.db)
	service := useCase.NewService(storage)
	apiHandler := authorHandler.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: apiHandler,
	}, nil

}
