package repo

import "database/sql"

type Repo struct {
	SalesRepoIR
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		SalesRepoIR: newSalesRepo(db),
	}
}
