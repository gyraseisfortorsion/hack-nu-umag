package service

import "back/internal/repo"

type Service struct {
	SalesServiceIR
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		SalesServiceIR: NewServiceSales(repo.SalesRepoIR),
	}
}
