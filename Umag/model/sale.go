package model

type Sale struct {
	Barcode   int `json:"barcode"`
	Quantity  int `json:"quantity"`
	Revenue   int `json:"revenue"`
	NetProfit int `json:"netProfit"`
}
