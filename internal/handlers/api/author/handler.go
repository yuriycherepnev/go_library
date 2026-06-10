package author

import (
	"github.com/julienschmidt/httprouter"
	"go-library/internal/handlers"
	"go-library/internal/usecase/author"
	"net/http"
)

const (
	authorURL  = "/author"
	authorsURL = "/authors"
)

type handler struct {
	bookService author.Service
}

func NewHandler(service author.Service) handlers.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(authorsURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte("books"))
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
