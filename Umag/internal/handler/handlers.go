package handler

import (
	"net/http"

	"back/internal/service"
)

type Handler struct {
	services *service.Service
	Router   *http.ServeMux
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		services: serv,
		Router:   http.NewServeMux(),
	}
}

func (h *Handler) Start() http.Handler {
	h.Router.HandleFunc("/api/sales", h.sales)

	return h.Router
}
