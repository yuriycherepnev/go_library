package composite

import (
	adapter "go-library/internal/adapters/db/book"
	domain "go-library/internal/domain/book"
	handler "go-library/internal/handlers"
	bookHandler "go-library/internal/handlers/api/book"
	useCase "go-library/internal/usecase/book"
)

type BookComposite struct {
	Storage domain.Repository
	Service useCase.Service
	Handler handler.Handler
}

func NewBookComposite(composite *MysqlComposite) (*BookComposite, error) {
	storage := adapter.NewStorage(composite.db)
	service := useCase.NewService(storage)
	apiHandler := bookHandler.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: apiHandler,
	}, nil

}
