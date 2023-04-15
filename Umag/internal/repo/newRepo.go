package repo

import "database/sql"

type Repo struct {
	SalesRepoIR
	SuppliesRepoIR
	ReportRepoIR
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		SalesRepoIR:    newSalesRepo(db),
		SuppliesRepoIR: newSuppliesRepo(db),
		ReportRepoIR:   newReportRepo(db),
	}
}
