package composite

import (
	authorAdapter "go-library/internal/adapters/db/author"
	bookAdapter "go-library/internal/adapters/db/book"
	readerAdapter "go-library/internal/adapters/db/reader"
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
	authorStorage := authorAdapter.NewStorage(composite.db)
	bookStorage := bookAdapter.NewStorage(composite.db)
	readerStorage := readerAdapter.NewStorage(composite.db)

	service := bookUseCase.NewService(authorStorage, bookStorage, readerStorage)
	apiHandler := bookHandler.NewHandler(service)

	return &BookComposite{
		Storage: bookStorage,
		Service: service,
		Handler: apiHandler,
	}, nil
}
