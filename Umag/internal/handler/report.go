package handler

import (
	"fmt"
	"log"
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
		fmt.Println("Error parsing time string:", err)
		return
	}
	toTime, err := time.Parse(layout, r.URL.Query().Get("toTime"))
	if err != nil {
		fmt.Println("Error parsing time string:", err)
		return
	}
	log.Println(barcode, fromTime, toTime)
}
