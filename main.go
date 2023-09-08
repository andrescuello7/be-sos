package main

import (
	"go-sms/db"
	"go-sms/models"
	"go-sms/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
	}

	var DSN = os.Getenv("URL_CONNECT_DB")
	db.PostgresConnect(DSN)
	db.DB.AutoMigrate(models.Message{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.MessageRouter(router)

	log.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}
