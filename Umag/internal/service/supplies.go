package service

import (
	"time"

	"back/internal/repo"
	"back/model"
)

type SupplyServiceIR interface {
	GetById(int) ([]model.Item, error)
	Get(int, time.Time, time.Time) ([]model.Item, error)
	CreateSales(model.Item) (int, error)
	UpdateSales(id int, item model.Item) error
	DeleteSales(id int) error
}

type SupplyServiceStr struct {
	repo repo.SuppliesRepoIR
}

func NewServiceSupply(repo repo.SuppliesRepoIR) SalesServiceIR {
	return &SupplyServiceStr{
		repo: repo,
	}
}

func (s *SupplyServiceStr) GetById(id int) ([]model.Item, error) {
	return s.repo.GetById(id)
}

func (s *SupplyServiceStr) Get(barcode int, fromTime time.Time, toTime time.Time) ([]model.Item, error) {
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

func (s *SupplyServiceStr) CreateSales(item model.Item) (int, error) {
	return s.repo.CreateSales(item)
}

func (s *SupplyServiceStr) UpdateSales(id int, item model.Item) error {
	return s.repo.UpdateSales(id, item)
}

func (s *SupplyServiceStr) DeleteSales(id int) error {
	return s.repo.DeleteSales(id)
}
