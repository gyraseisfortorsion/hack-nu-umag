package service

import (
	"time"

	"back/internal/repo"
	"back/model"
)

type SalesServiceIR interface {
	GetById(int) ([]model.Item, error)
	Get(int, time.Time, time.Time) ([]model.Item, error)
	CreateSales(model.Item) (int, error)
	UpdateSales(id int, item model.Item) error
	DeleteSales(id int) error
}

type SalesServiceStr struct {
	repo repo.SalesRepoIR
}

func NewServiceSales(repo repo.SalesRepoIR) SalesServiceIR {
	return &SalesServiceStr{
		repo: repo,
	}
}

func (s *SalesServiceStr) GetById(id int) ([]model.Item, error) {
	return s.repo.GetById(id)
}

func (s *SalesServiceStr) Get(barcode int, fromTime time.Time, toTime time.Time) ([]model.Item, error) {
	if barcode == 0 && fromTime.IsZero() && !toTime.IsZero() {
		return s.repo.GetByToTime(toTime)
	}
	if barcode == 0 && !fromTime.IsZero() && toTime.IsZero() {
		return s.repo.GetByFromTime(fromTime)
	}
	if barcode != 0 && fromTime.IsZero() && toTime.IsZero() {
		return s.repo.GetByBarcode(barcode)
	}
	if barcode == 0 && !fromTime.IsZero() && !toTime.IsZero() {
		return s.repo.GetByFromToTime(fromTime, toTime)
	}
	if barcode != 0 && fromTime.IsZero() && !toTime.IsZero() {
		return s.repo.GetByBarcodeAndtoTime(barcode, toTime)
	}
	if barcode != 0 && !fromTime.IsZero() && toTime.IsZero() {
		return s.repo.GetByBarcodeAndFromTime(barcode, fromTime)
	}
	return s.repo.Get(barcode, fromTime, toTime)
}

func (s *SalesServiceStr) CreateSales(item model.Item) (int, error) {
	return s.repo.CreateSales(item)
}

func (s *SalesServiceStr) UpdateSales(id int, item model.Item) error {
	return s.repo.UpdateSales(id, item)
}

func (s *SalesServiceStr) DeleteSales(id int) error {
	return s.repo.DeleteSales(id)
}
