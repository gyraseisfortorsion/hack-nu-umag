package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) reports(w http.ResponseWriter, r *http.Request) {
	barcode, err := strconv.Atoi(r.URL.Query().Get("barcode"))
	if err != nil {
		http.Error(w, "err", http.StatusNotFound)
		return
	}
	layout := "2006-01-02 15:04:05"

	fromTime, err := time.Parse(layout, r.URL.Query().Get("fromTime"))
	if err != nil {
		http.Error(w, "err", http.StatusNotFound)
		return
	}
	toTime, err := time.Parse(layout, r.URL.Query().Get("toTime"))
	if err != nil {
		http.Error(w, "err", http.StatusNotFound)
		return
	}
	resp, err := h.services.ReportServiceIR.GetReport(barcode, fromTime, toTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
