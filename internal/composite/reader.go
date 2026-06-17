package composite

import (
	adapter "go-library/internal/adapters/db/reader"
	domain "go-library/internal/domain/reader"
	handler "go-library/internal/handlers"
	readerHandler "go-library/internal/handlers/api/reader"
	useCase "go-library/internal/usecase/reader"
)

type ReaderComposite struct {
	Storage domain.Repository
	Service useCase.Service
	Handler handler.Handler
}

func NewReaderComposite(composite *MysqlComposite) (*ReaderComposite, error) {
	storage := adapter.NewStorage(composite.db)
	service := useCase.NewService(storage)
	apiHandler := readerHandler.NewHandler(service)

	return &ReaderComposite{
		Storage: storage,
		Service: service,
		Handler: apiHandler,
	}, nil

}
