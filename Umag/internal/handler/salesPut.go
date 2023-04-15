package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"back/model"
)

func (h *Handler) SalesPut(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/sales/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var item model.Item
	if err := json.Unmarshal(body, &item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.services.SalesServiceIR.UpdateSales(id, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
