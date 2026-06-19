package reader

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go-library/internal/handlers"
	"go-library/internal/usecase/reader"
	"net/http"
	"strconv"
	"strings"
)

type handler struct {
	readerService reader.Service
}

func NewHandler(service reader.Service) handlers.Handler {
	return &handler{readerService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/reader/:id", h.GetReaderByID)
	router.GET("/readers", h.GetReader)
	router.POST("/reader", h.CreateReader)
	router.PATCH("/reader/:id", h.UpdateReader)
	router.DELETE("/reader/:id", h.DeleteReader)
}

func (h *handler) GetReader(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	readers, err := h.readerService.GetAll()
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, readers)
}

func (h *handler) GetReaderByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	a, err := h.readerService.GetById(id)
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, a)
}

func (h *handler) CreateReader(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var request reader.CreateReaderDTO
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	a, err := h.readerService.Create(request)
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusCreated, a)
}

func (h *handler) UpdateReader(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var request reader.UpdateReaderDTO
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	request.Id = id
	request.Name = strings.TrimSpace(request.Name)
	a, err := h.readerService.Update(request)
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, a)
}

func (h *handler) DeleteReader(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}
	request := reader.DeleteReaderDTO{Id: id}
	a, err := h.readerService.Delete(request)
	if err != nil {
		handlers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	handlers.WriteJSON(w, http.StatusOK, a)
}
