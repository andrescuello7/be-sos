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
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permite todas las solicitudes de cualquier origen
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.MessageRouter(router)

	handler := c.Handler(router)

	log.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, handler)
}
