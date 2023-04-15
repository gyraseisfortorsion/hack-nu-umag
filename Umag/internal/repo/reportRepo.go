package repo

import (
	"database/sql"
	"time"

	"back/model"
)

type ReportRepoIR interface {
	GetReport(barcode int, fromTime, toTime time.Time) (model.Sale, error)
}

type ReportRepoStr struct {
	db *sql.DB
}

func newReportRepo(db *sql.DB) ReportRepoIR {
	return &ReportRepoStr{
		db: db,
	}
}

func (r *ReportRepoStr) GetReport(barcode int, fromTime, toTime time.Time) (model.Sale, error) {
	var sale model.Sale
	saleRows, err := r.db.Query(`SELECT barcode, SUM(quantity) as quantity, SUM(quantity*price) as revenue FROM sale
	WHERE barcode = ? AND sale_time >= ? AND sale_time <= ?
	GROUP BY barcode`, barcode, fromTime, toTime)
	if err != nil {
		return sale, err
	}
	defer saleRows.Close()

	supplyRows, err := r.db.Query(`SELECT barcode, SUM(quantity) as quantity, SUM(quantity*price) as cost FROM supply
	WHERE barcode = ? AND supply_time >= ? AND supply_time <= ?
	GROUP BY barcode`, barcode, fromTime, toTime)
	if err != nil {
		return sale, err
	}
	defer supplyRows.Close()
	var supply struct {
		Barcode  int64 `json:"barcode"`
		Quantity int   `json:"quantity"`
		Cost     int   `json:"cost"`
	}
	if saleRows.Next() {
		if err := saleRows.Scan(&sale.Barcode, &sale.Quantity, &sale.Revenue); err != nil {
			return sale, err
		}
	}
	if supplyRows.Next() {
		if err := supplyRows.Scan(&supply.Barcode, &supply.Quantity, &supply.Cost); err != nil {
			return sale, err
		}
	}
	netProfit := sale.Revenue - supply.Cost
	result := model.Sale{Barcode: barcode, Quantity: sale.Quantity, Revenue: sale.Revenue, NetProfit: netProfit}
	return result, nil
}
