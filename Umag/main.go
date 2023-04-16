package main

import (
	"log"
	"net/http"

	"back/internal/database"
	"back/internal/handler"
	"back/internal/repo"
	"back/internal/service"
)

func main() {
	configDb := database.NewConfDb()
	db := configDb.InitDB()
	// configDb.AddCollumns(db)
	// configDb.InsertDataSales(db)
	// configDb.InsertDataSupplies(db)
	defer db.Close()
	log.Println("Successfully Initiated the Data Base")

	repo := repo.NewRepo(db)
	log.Println("Successfully Initiated the Repository")

	serve := service.NewService(repo)
	log.Println("Successfully Initiated the Service")

	handler := handler.NewHandler(serve)
	if err := http.ListenAndServe(":8080", handler.Start()); err != nil {
		panic(err)
	}
}
