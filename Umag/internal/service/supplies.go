package service

import (
	"time"

	"back/internal/repo"
	"back/model"
)

type SupplyServiceIR interface {
	GetById(int) (model.ItemSupply, error)
	Get(int, time.Time, time.Time) ([]model.ItemSupply, error)
	CreateSupply(model.ItemSupply) (int, error)
	UpdateSupply(id int, item model.ItemSupply) error
	DeleteSupply(id int) error
}

type SupplyServiceStr struct {
	repo repo.SuppliesRepoIR
}

func NewServiceSupply(repo repo.SuppliesRepoIR) SupplyServiceIR {
	return &SupplyServiceStr{
		repo: repo,
	}
}

func (s *SupplyServiceStr) GetById(id int) (model.ItemSupply, error) {
	return s.repo.GetById(id)
}

func (s *SupplyServiceStr) Get(barcode int, fromTime time.Time, toTime time.Time) ([]model.ItemSupply, error) {
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

func (s *SupplyServiceStr) CreateSupply(item model.ItemSupply) (int, error) {
	return s.repo.CreateSupply(item)
}

func (s *SupplyServiceStr) UpdateSupply(id int, item model.ItemSupply) error {
	return s.repo.UpdateSupply(id, item)
}

func (s *SupplyServiceStr) DeleteSupply(id int) error {
	return s.repo.DeleteSupply(id)
}
