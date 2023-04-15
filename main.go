package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "password"
	dbName     = "supply"
	dbHost     = "localhost"
	dbPort     = 3306
)

var db *sql.DB

type Product struct {
	ID        int             `json:"id"`
	Barcode   string          `json:"barcode"`
	Quantity  int             `json:"quantity"`
	Price     sql.NullFloat64 `json:"price"`
	Timestamp time.Time       `json:"timestamp"`
}

func main() {
	// Connect to the database
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err)
		//return
	}
	defer db.Close()

	// Create the router
	r := mux.NewRouter()

	// Add the routes
	r.HandleFunc("/products", getProducts).Methods("GET").Queries("fromTime", "{fromTime}", "toTime", "{toTime}", "barcode", "{barcode}")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	// Enable CORS
	handler := cors.Default().Handler(r)

	// Start the server
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", handler)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	// Get the query parameters
	fromTimeStr := r.FormValue("fromTime")
	toTimeStr := r.FormValue("toTime")
	barcode := r.FormValue("barcode")

	// Parse the timestamps
	fromTime, err := time.Parse(time.RFC3339, fromTimeStr)
	if err != nil {
		http.Error(w, "Invalid fromTime", http.StatusBadRequest)
		return
	}
	toTime, err := time.Parse(time.RFC3339, toTimeStr)
	if err != nil {
		http.Error(w, "Invalid toTime", http.StatusBadRequest)
		return
	}

	// Build the query
	query := "SELECT id, barcode, quantity, price, timestamp FROM products WHERE timestamp BETWEEN ? AND ?"
	args := []interface{}{fromTime, toTime}
	if barcode != "" {
		query += " AND barcode = ?"
		args = append(args, barcode)
	}

	// Execute the query
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Build the
	// products slice
	products := []Product{}
	// Iterate over the rows
	for rows.Next() {
		// Create a new Product
		var product Product
		err := rows.Scan(&product.ID, &product.Barcode, &product.Quantity, &product.Price, &product.Timestamp)
		if err != nil {
			panic(err)
		}

		// Append the Product to the slice
		products = append(products, product)
	}

	// Encode the products slice as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Product
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Set the timestamp to the current time
	product.Timestamp = time.Now()

	// Insert the Product into the database
	result, err := db.Exec("INSERT INTO products (barcode, quantity, price, timestamp) VALUES (?, ?, ?, ?)", product.Barcode, product.Quantity, product.Price, product.Timestamp)
	if err != nil {
		panic(err)
	}

	// Get the ID of the new Product
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Set the ID of the Product
	product.ID = int(id)

	// Encode the Product as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

func getProduct(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	// Convert the ID parameter to an int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Query the database for the Product with the given ID
	row := db.QueryRow("SELECT id, barcode, quantity, price, timestamp FROM products WHERE id = ?", id)

	// Create a new Product
	var product Product
	err = row.Scan(&product.ID, &product.Barcode, &product.Quantity, &product.Price, &product.Timestamp)
	if err == sql.ErrNoRows {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		panic(err)
	}

	// Encode the Product as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	// Convert the ID parameter to an int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Decode the request body into a Product
	var product Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure that the ID of the request body matches the ID in the URL
	if product.ID != id {
		http.Error(w, "ID in request body does not match ID in URL", http.StatusBadRequest)
		return
	}

	// Update the Product in the database
	result, err := db.Exec("UPDATE products SET barcode = ?, quantity = ?, price = ?, timestamp = ? WHERE id = ?", product.Barcode, product.Quantity, product.Price, product.Timestamp, id)
	if err != nil {
		panic(err)
	}
	// Get the number of rows affected by the update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	// If no rows were affected, return a 404 Not Found error
	if rowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Encode the Product as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	// Convert the ID parameter to an int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Delete the Product from the database
	result, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	// Get the number of rows affected by the delete
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	// If no rows were affected, return a 404 Not Found error
	if rowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Return a 204 No Content status code
	w.WriteHeader(http.StatusNoContent)
}

// func main() {
// // Open a connection to the MySQL database
// db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
// if err != nil {
// panic(err)
// }
// // Close the database connection when the program exits
// defer db.Close()

// // Create a new Router
// r := mux.NewRouter()

// // Define the routes
// r.HandleFunc("/products", getProducts).Methods("GET").Queries("fromTime", "{fromTime}", "toTime", "{toTime}", "barcode", "{barcode}")
// r.HandleFunc("/products", createProduct).Methods("POST")
// r.HandleFunc("/products/{id}", getProduct).Methods("GET")
// r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
// r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

// // Start the HTTP server
// http.ListenAndServe(":8080", r)
// }
