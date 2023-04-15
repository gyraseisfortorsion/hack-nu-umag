package service

import (
	"time"

	"back/internal/repo"
	"back/model"
)

type ReportServiceIR interface {
	GetReport(barcode int, fromTime, toTime time.Time) (model.Sale, error)
}

type ReportServiceStr struct {
	repo repo.ReportRepoIR
}

func NewServiceReport(repo repo.ReportRepoIR) ReportServiceIR {
	return &ReportServiceStr{
		repo,
	}
}

func (r *ReportServiceStr) GetReport(barcode int, fromTime, toTime time.Time) (model.Sale, error) {
	return r.repo.GetReport(barcode, fromTime, toTime)
}
