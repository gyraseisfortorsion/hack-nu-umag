package handler

import "net/http"

func (h *Handler) supplies(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.SuppliesGet(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.SuppliesPost(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.SuppliesPut(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		h.SuppliesDelete(w, r)
		return
	}
}
