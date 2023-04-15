package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) SuppliesGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("id") {
		h.SuppliesGetById(w, r)
		return
	}
	var barcode int
	if r.URL.Query().Has("barcode") {
		var err error
		barcode, err = strconv.Atoi(r.URL.Query().Get("barcode"))
		if err != nil {
			http.Error(w, "err", http.StatusNotFound)
			return
		}
	}
	var fromTime time.Time
	if r.URL.Query().Has("fromTime") {
		var err error
		layout := "2006-01-02 15:04:05"
		fromTime, err = time.Parse(layout, r.URL.Query().Get("fromTime"))
		if err != nil {
			fmt.Println("Error parsing time string:", err)
			return
		}
	}
	var toTime time.Time
	if r.URL.Query().Has("toTime") {
		var err error
		layout := "2006-01-02 15:04:05"
		toTime, err = time.Parse(layout, r.URL.Query().Get("toTime"))
		if err != nil {
			fmt.Println("Error parsing time string:", err)
			return
		}
	}
	// log.Println(barcode)
	// log.Println(toTime)
	Items, err := h.services.SupplyServiceIR.Get(barcode, toTime, fromTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(Items); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) SuppliesGetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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
}
