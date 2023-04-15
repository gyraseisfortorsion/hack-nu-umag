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
	h.Router.HandleFunc("/api/sales/", h.salesId)
	h.Router.HandleFunc("/api/supplies", h.supplies)
	h.Router.HandleFunc("/api/supplies/", h.suppliesId)

	h.Router.HandleFunc("/api/reports", h.reports)
	return h.Router
}
