package handler

import (
	"github.com/gorilla/mux"
)

type Handler struct {
	storage string
}

func NewHandler(storageName string) *Handler {
	return &Handler{storage: storageName}
}

func (h *Handler) InitRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/{queue}", h.GetHandler).Methods("GET")
	r.HandleFunc("/{queue}", h.PutHandler).Methods("PUT")
	return r
}
