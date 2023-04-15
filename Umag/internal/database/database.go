package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ConfigDb struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

func NewConfDb() *ConfigDb {
	var AppConfig ConfigDb
	raw, err := ioutil.ReadFile("configDB.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(raw, &AppConfig)
	return &AppConfig
}

func (c *ConfigDb) InitDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/umag_hacknu", c.Username, c.Password, c.Host, c.Port)
	db, err := sql.Open(c.Driver, dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func (c *ConfigDb) CreateTables(db *sql.DB) {
	log.Println("CreateTables")
	file, err := ioutil.ReadFile("./migrations/create_tables.sql")
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := db.Exec(string(file)); err != nil {
		log.Fatal(err.Error())
	}
}

func (c *ConfigDb) InsertDataSales(db *sql.DB) {
	file, err := ioutil.ReadFile("./migrations/insert_sales.sql")
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := db.Exec(string(file)); err != nil {
		log.Fatal(err.Error())
	}
}

func (c *ConfigDb) InsertDataSupplies(db *sql.DB) {
	file, err := ioutil.ReadFile("./migrations/insert_supplies.sql")
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := db.Exec(string(file)); err != nil {
		log.Fatal(err.Error())
	}
}
