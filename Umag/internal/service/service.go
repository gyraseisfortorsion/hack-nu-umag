package service

import "back/internal/repo"

type Service struct {
	SalesServiceIR
	SupplyServiceIR
	ReportServiceIR
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		SalesServiceIR:  NewServiceSales(repo.SalesRepoIR),
		SupplyServiceIR: NewServiceSupply(repo.SuppliesRepoIR),
		ReportServiceIR: NewServiceReport(repo.ReportRepoIR),
	}
}
