package repo

import (
	"database/sql"
	"log"
	"time"

	"back/model"
)

type SuppliesRepoIR interface {
	GetById(int) (model.ItemSupply, error)
	Get(int, time.Time, time.Time) ([]model.ItemSupply, error)
	GetByToTime(time.Time) ([]model.ItemSupply, error)
	GetByFromTime(time.Time) ([]model.ItemSupply, error)
	GetByBarcode(int) ([]model.ItemSupply, error)
	GetByFromToTime(time.Time, time.Time) ([]model.ItemSupply, error)
	GetByBarcodeAndtoTime(int, time.Time) ([]model.ItemSupply, error)
	GetByBarcodeAndFromTime(barcode int, fromTime time.Time) ([]model.ItemSupply, error)
	CreateSupply(item model.ItemSupply) (int, error)
	UpdateSupply(id int, item model.ItemSupply) error
	DeleteSupply(id int) error
}

type SuppliesRepoStr struct {
	db *sql.DB
}

func newSuppliesRepo(db *sql.DB) SuppliesRepoIR {
	return &SuppliesRepoStr{
		db: db,
	}
}

func (s *SuppliesRepoStr) GetById(id int) (model.ItemSupply, error) {
	var item model.ItemSupply
	query := `SELECT 
		id,
		barcode,
		quantity,
		price,
		supply_time
	FROM
		supply
	WHERE
		id = ?`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return item, err
	}
	defer stmt.Close()

	var saleTimeStr string
	if err := stmt.QueryRow(id).Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
		log.Println(err.Error())
		return item, err
	}
	item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
	if err != nil {
		log.Println(err.Error())
		return item, err

	}

	return item, nil
}

func (s *SuppliesRepoStr) Get(barcode int, fromTime time.Time, toTime time.Time) ([]model.ItemSupply, error) {
	rows, err := s.db.Query("SELECT * FROM supply WHERE barcode = ? AND supply_time BETWEEN ? AND ?", barcode, fromTime, toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.ItemSupply
	for rows.Next() {
		var item model.ItemSupply
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
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

func (s *SuppliesRepoStr) GetByToTime(toTime time.Time) ([]model.ItemSupply, error) {
	rows, err := s.db.Query("SELECT * FROM supply WHERE supply_time < ? ", toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.ItemSupply
	for rows.Next() {
		var item model.ItemSupply
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
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

func (s *SuppliesRepoStr) GetByFromTime(fromTime time.Time) ([]model.ItemSupply, error) {
	rows, err := s.db.Query("SELECT * FROM supply WHERE supply_time > ? ", fromTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.ItemSupply
	for rows.Next() {
		var item model.ItemSupply
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
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

func (s *SuppliesRepoStr) GetByBarcode(Barcode int) ([]model.ItemSupply, error) {
	rows, err := s.db.Query("SELECT * FROM supply WHERE barcode = ? ", Barcode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.ItemSupply
	for rows.Next() {
		var item model.ItemSupply
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
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

func (s *SuppliesRepoStr) GetByFromToTime(fromTime time.Time, toTime time.Time) ([]model.ItemSupply, error) {
	rows, err := s.db.Query("SELECT * FROM supply WHERE supply_time BETWEEN ? AND ? ", fromTime, toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.ItemSupply
	for rows.Next() {
		var item model.ItemSupply
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
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

func (s *SuppliesRepoStr) GetByBarcodeAndtoTime(barcode int, toTime time.Time) ([]model.ItemSupply, error) {
	rows, err := s.db.Query("SELECT * FROM supply WHERE barcode = ? and supply_time > ?  ", barcode, toTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.ItemSupply
	for rows.Next() {
		var item model.ItemSupply
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
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

func (s *SuppliesRepoStr) GetByBarcodeAndFromTime(barcode int, fromTime time.Time) ([]model.ItemSupply, error) {
	rows, err := s.db.Query("SELECT * FROM supply WHERE barcode = ? and supply_time < ?  ", barcode, fromTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.ItemSupply
	for rows.Next() {
		var item model.ItemSupply
		var saleTimeStr string
		if err := rows.Scan(&item.ID, &item.Barcode, &item.Quantity, &item.Price, &saleTimeStr); err != nil {
			return nil, err
		}
		item.SupplyTime, err = time.Parse("2006-01-02 15:04:05", saleTimeStr)
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

func (s *SuppliesRepoStr) CreateSupply(item model.ItemSupply) (int, error) {
	stmt, err := s.db.Prepare("INSERT INTO supply (Barcode, Price, Quantity, supply_time) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(item.Barcode, item.Price, item.Quantity, item.SupplyTime)
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastID), nil
}

func (s *SuppliesRepoStr) UpdateSupply(id int, item model.ItemSupply) error {
	stmt, err := s.db.Prepare("UPDATE supply SET barcode=?, price=?, quantity=?, supply_time=? WHERE id=?")
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
	_, err = stmt.Exec(item.Barcode, item.Price, item.Quantity, item.SupplyTime, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SuppliesRepoStr) DeleteSupply(id int) error {
	stmt, err := s.db.Prepare("Delete FROM supply WHERE id=?")
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
