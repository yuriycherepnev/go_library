package author

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"go-library/internal/handlers"
	"go-library/internal/usecase/author"
	"net/http"
	"strconv"
	"strings"
)

type handler struct {
	authorService author.Service
}

func NewHandler(service author.Service) handlers.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/author/:id", h.GetAuthorByID)
	router.GET("/authors", h.GetAuthors)
	router.POST("/author", h.CreateAuthor)
	router.PUT("/author/:id", h.UpdateAuthor)
	router.DELETE("/author/:id", h.DeleteAuthor)
}

func (h *handler) GetAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authors, err := h.authorService.GetAll()
	if err != nil {
		handlers.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, authors)
}

func (h *handler) GetAuthorByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	a, err := h.authorService.GetById(id)
	if err != nil {
		handlers.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, a)
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var request author.CreateAuthorDTO
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	a, err := h.authorService.Create(request)
	if err != nil {
		handlers.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusCreated, a)
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var request author.UpdateAuthorDTO
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	request.Id = id
	request.Name = strings.TrimSpace(request.Name)
	a, err := h.authorService.Update(request)
	if err != nil {
		handlers.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, a)
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	request := author.DeleteAuthorDTO{Id: id}
	a, err := h.authorService.Delete(request)
	if errors.Is(err, sql.ErrNoRows) {
		handlers.WriteError(w, http.StatusNotFound, "author not found")
		return
	}
	if err != nil {
		handlers.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, a)
}
