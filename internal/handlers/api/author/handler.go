package author

import (
	"github.com/julienschmidt/httprouter"
	"go-library/internal/handlers"
	"go-library/internal/usecase/author"
	"net/http"
	"strconv"
)

type handler struct {
	authorService author.Service
}

func NewHandler(service author.Service) handlers.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/author/:id", h.GetAuthorByID)
	router.GET("/author", h.GetAuthors)
	//router.POST("/author", h.CreateAuthor)
	//router.PUT("/author/:id", h.UpdateAuthor)
	//router.DELETE("/author/:id", h.DeleteAuthor)
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
