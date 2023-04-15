package handler

import (
	"net/http"
	"strconv"
)

func (h *Handler) SalesDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/sales/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.services.SalesServiceIR.DeleteSales(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
