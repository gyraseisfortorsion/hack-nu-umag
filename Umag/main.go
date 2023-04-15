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
	defer db.Close()
	log.Println("Successfully Initiated the Data Base")
	repo := repo.NewRepo(db)
	serve := service.NewService(repo)
	handler := handler.NewHandler(serve)
	if err := http.ListenAndServe(":8080", handler.Start()); err != nil {
		panic(err)
	}
}
