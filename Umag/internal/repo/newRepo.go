package repo

import "database/sql"

type Repo struct {
	SalesRepoIR
	SuppliesRepoIR
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		SalesRepoIR:    newSalesRepo(db),
		SuppliesRepoIR: newSuppliesRepo(db),
	}
}
