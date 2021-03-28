package handler

import (
	"github.com/alipniczkij/web-broker/pkg/repository"
	"github.com/gorilla/mux"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) InitRoutes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/{queue}", h.GetHandler).Methods("GET")
	r.HandleFunc("/{queue}", h.PutHandler).Methods("PUT")
	return r
}
