package composite

import (
	authorAdapter "go-library/internal/adapters/db/author"
	bookAdapter "go-library/internal/adapters/db/book"
	bookDomain "go-library/internal/domain/book"
	handler "go-library/internal/handlers"
	bookHandler "go-library/internal/handlers/api/book"
	bookUseCase "go-library/internal/usecase/book"
)

type BookComposite struct {
	Storage bookDomain.Repository
	Service bookUseCase.Service
	Handler handler.Handler
}

func NewBookComposite(composite *MysqlComposite) (*BookComposite, error) {
	bookStorage := bookAdapter.NewStorage(composite.db)
	authorStorage := authorAdapter.NewStorage(composite.db)

	service := bookUseCase.NewService(bookStorage, authorStorage)
	apiHandler := bookHandler.NewHandler(service)

	return &BookComposite{
		Storage: bookStorage,
		Service: service,
		Handler: apiHandler,
	}, nil
}
