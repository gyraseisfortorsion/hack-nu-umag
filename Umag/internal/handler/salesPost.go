package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"back/model"
)

func (h *Handler) SalesPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var item model.Item
	err = json.Unmarshal(body, &item)
	id, err := h.services.SalesServiceIR.CreateSales(item)
	resp := struct {
		id int
	}{
		id,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
