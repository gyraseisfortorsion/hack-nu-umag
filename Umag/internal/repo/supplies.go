package repo

import (
	"database/sql"
	"time"

	"back/model"
)

type SuppliesRepoIR interface {
	GetById(int) ([]model.Item, error)
	Get(int, time.Time, time.Time) ([]model.Item, error)
	GetByToTime(time.Time) ([]model.Item, error)
	GetByFromTime(time.Time) ([]model.Item, error)
	GetByBarcode(int) ([]model.Item, error)
	GetByFromToTime(time.Time, time.Time) ([]model.Item, error)
	GetByBarcodeAndtoTime(int, time.Time) ([]model.Item, error)
	GetByBarcodeAndFromTime(barcode int, fromTime time.Time) ([]model.Item, error)
	CreateSales(item model.Item) (int, error)
	UpdateSales(id int, item model.Item) error
	DeleteSales(id int) error
}

type SuppliesRepoStr struct {
	db *sql.DB
}

func newSuppliesRepo(db *sql.DB) SuppliesRepoIR {
	return &SuppliesRepoStr{
		db: db,
	}
}

func (s *SuppliesRepoStr) GetById(id int) ([]model.Item, error) {
	var items []model.Item
	query := `SELECT 
		id,
		barcode,
		quantity,
		price,
		sale_time
	FROM
		sale
	WHERE
		id = ?`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *SuppliesRepoStr) Get(barcode int, fromTime time.Time, toTime time.Time) ([]model.Item, error) {
	rows, err := s.db.Query("SELECT * FROM sale WHERE barcode = ? AND sale_time BETWEEN ? AND ?", barcode, fromTime, toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SuppliesRepoStr) GetByToTime(toTime time.Time) ([]model.Item, error) {
	rows, err := s.db.Query("SELECT * FROM sale WHERE sale_time < ? ", toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SuppliesRepoStr) GetByFromTime(fromTime time.Time) ([]model.Item, error) {
	rows, err := s.db.Query("SELECT * FROM sale WHERE sale_time > ? ", fromTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SuppliesRepoStr) GetByBarcode(Barcode int) ([]model.Item, error) {
	rows, err := s.db.Query("SELECT * FROM sale WHERE barcode = ? ", Barcode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.Item
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SuppliesRepoStr) GetByFromToTime(fromTime time.Time, toTime time.Time) ([]model.Item, error) {
	rows, err := s.db.Query("SELECT * FROM sale WHERE sale_time BETWEEN ? AND ? ", fromTime, toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SuppliesRepoStr) GetByBarcodeAndtoTime(barcode int, toTime time.Time) ([]model.Item, error) {
	rows, err := s.db.Query("SELECT * FROM sale WHERE barcode = ? and sale_time > ?  ", barcode, toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SuppliesRepoStr) GetByBarcodeAndFromTime(barcode int, fromTime time.Time) ([]model.Item, error) {
	rows, err := s.db.Query("SELECT * FROM sale WHERE barcode = ? and sale_time < ?  ", barcode, fromTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SaleTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SuppliesRepoStr) CreateSales(item model.Item) (int, error) {
	stmt, err := s.db.Prepare("INSERT INTO sale (Barcode, Price, Quantity, SaleTime) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(item.Barcode, item.Price, item.Quantity, item.SaleTime)
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastID), nil
}

func (s *SuppliesRepoStr) UpdateSales(id int, item model.Item) error {
	stmt, err := s.db.Prepare("UPDATE sale SET barcode=?, price=?, quantity=?, sale_time=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	/*
			    {
		      "barcode": 12334565,
		      "price": 123,
		      "quantity": 2,
		      "saleTime": "2022-12-28 11:00:02"
		    }
	*/
	_, err = stmt.Exec(item.Barcode, item.Price, item.Quantity, item.SaleTime, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SuppliesRepoStr) DeleteSales(id int) error {
	stmt, err := s.db.Prepare("Delete FROM sale WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
