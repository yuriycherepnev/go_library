package book

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"go-library/internal/handlers"
	"go-library/internal/usecase/book"
	"net/http"
	"strconv"
	"strings"
)

type handler struct {
	bookService book.Service
}

func NewHandler(service book.Service) handlers.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/book/:id", h.GetBookById)
	router.GET("/books", h.GetAllBooks)
	router.POST("/book", h.CreateBook)
	router.PATCH("/book/:id", h.UpdateBook)
	router.DELETE("/book/:id", h.DeleteBook)
}

func (h *handler) GetBookById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	b, err := h.bookService.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlers.WriteError(w, http.StatusNotFound, "book not found")
			return
		}

		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	handlers.WriteJSON(w, http.StatusOK, b)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books, err := h.bookService.GetAll()
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, books)
}

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request book.CreateBookDTO

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if request.Title == "" {
		handlers.WriteError(w, http.StatusBadRequest, "title is required")
		return
	}

	if request.IdAuthor <= 0 {
		handlers.WriteError(w, http.StatusBadRequest, "id_author is required")
		return
	}

	b, err := h.bookService.Create(request)
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	handlers.WriteJSON(w, http.StatusCreated, b)
}

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var request book.UpdateBookDTO

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	request.Id = id
	request.Title = strings.TrimSpace(request.Title)

	if request.Title == "" {
		handlers.WriteError(w, http.StatusBadRequest, "title is required")
		return
	}

	b, err := h.bookService.Update(request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlers.WriteError(w, http.StatusNotFound, "book not found")
			return
		}

		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	handlers.WriteJSON(w, http.StatusOK, b)
}

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	request := book.DeleteBookDTO{
		Id: id,
	}

	b, err := h.bookService.Delete(request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlers.WriteError(w, http.StatusNotFound, "book not found")
			return
		}

		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	handlers.WriteJSON(w, http.StatusOK, b)
}

func (h *handler) BorrowBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request book.BorrowBookDTO

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	b, err := h.bookService.Borrow(request)
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	handlers.WriteJSON(w, http.StatusOK, b)
}
