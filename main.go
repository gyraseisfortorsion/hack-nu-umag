package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Define a struct for our Supply model
type Supply struct {
	ID       int       `json:"id"`
	Quantity int       `json:"quantity"`
	Price    float64   `json:"price"`
	Time     time.Time `json:"time"`
}

// Define global DB variable
var db *sql.DB

// Function to initialize the database connection
func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:password5@tcp(127.0.0.1:3306)/umag_hacknu")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// Function to retrieve all supplies from the database
func getAllSupplies(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM supplies")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	supplies := make([]Supply, 0)
	for rows.Next() {
		supply := Supply{}
		err := rows.Scan(&supply.ID, &supply.Quantity, &supply.Price, &supply.Time)
		if err != nil {
			log.Fatal(err)
		}
		supplies = append(supplies, supply)
	}

	json.NewEncoder(w).Encode(supplies)
}

// Function to retrieve a single supply by ID from the database
func getSupply(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT * FROM supplies WHERE id = ?", id)

	supply := Supply{}
	err = row.Scan(&supply.ID, &supply.Quantity, &supply.Price, &supply.Time)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(supply)
}

// Function to create a new supply in the database
func createSupply(w http.ResponseWriter, r *http.Request) {
	var supply Supply
	json.NewDecoder(r.Body).Decode(&supply)

	result, err := db.Exec("INSERT INTO supplies (quantity, price, time) VALUES (?, ?, ?)", supply.Quantity, supply.Price, supply.Time)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	supply.ID = int(lastInsertID)

	json.NewEncoder(w).Encode(supply)
}

// Function to update an existing supply in the database
func updateSupply(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	var supply Supply
	json.NewDecoder(r.Body).Decode(&supply)

	_, err = db.Exec("UPDATE supplies SET quantity=?, price=?, time=? WHERE id=?", supply.Quantity, supply.Price, supply.Time, id)
	if err != nil {
		log.Fatal(err)
	}

	supply.ID = id

	json.NewEncoder(w).Encode(supply)

}
func deleteSupply(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)
	w.WriteHeader(http.StatusNoContent)
}
func main() {
	// Initialize the database connection
	initDB()
	defer db.Close()
	router := mux.NewRouter()

	// Define our API endpoints
	router.HandleFunc("/supplies", getAllSupplies).Methods("GET")
	router.HandleFunc("/supplies/{id}", getSupply).Methods("GET")
	router.HandleFunc("/supplies", createSupply).Methods("POST")
	router.HandleFunc("/supplies/{id}", updateSupply).Methods("PUT")
	router.HandleFunc("/supplies/{id}", deleteSupply).Methods("DELETE")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
