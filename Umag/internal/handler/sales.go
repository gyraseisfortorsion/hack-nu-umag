package handler

import (
	"net/http"
)

func (h *Handler) sales(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.SalesGet(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.SalesPost(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.SalesPut(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		h.SalesDelete(w, r)
		return
	}
}
