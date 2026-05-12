package book

import (
	"ca-template/internal/adapters"
	"ca-template/internal/domain/book"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "/book"
	booksURL = "/books"
)

type handler struct {
	bookService book.Service
}

func NewHandler(service book.Service) adapters.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte("books"))
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
