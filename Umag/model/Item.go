package model

import "time"

type Item struct {
	ID       int       `json:"id"`
	Barcode  int       `json:"barcode"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
	SaleTime time.Time `json:"saleTime"`
}
