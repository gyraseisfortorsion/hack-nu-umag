package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) supplies(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.SuppliesGet(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.SuppliesPost(w, r)
		return
	}

	http.Error(w, "Mehtod Not allowd", http.StatusMethodNotAllowed)
	return
}

func (h *Handler) suppliesId(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.URL.Path[len("/api/supplies/"):])
		if err != nil {
			http.Error(w, "err", http.StatusNotFound)
			return
		}
		Items, err := h.services.SupplyServiceIR.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(Items); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		return
	}
	if r.Method == http.MethodDelete {
		h.SuppliesDelete(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.SuppliesPut(w, r)
		return
	}
	http.Error(w, "Mehtod Not allowd", http.StatusMethodNotAllowed)
	return
}
