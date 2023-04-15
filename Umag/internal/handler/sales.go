package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	return
}

func (h *Handler) salesId(w http.ResponseWriter, r *http.Request) {
	log.Println("asd")
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.URL.Path[len("/api/sales/"):])
		if err != nil {
			http.Error(w, "err", http.StatusNotFound)
			return
		}
		Items, err := h.services.SalesServiceIR.GetById(id)
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
		h.SalesDelete(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.SalesPut(w, r)
		return
	}
	http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	return
}
