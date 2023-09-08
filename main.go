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
	//environment variables -> .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading environment", err)
	}

	//get connection with db
	var DSN = os.Getenv("URL_CONNECT_DB")
	db.PostgresConnect(DSN)
	db.DB.AutoMigrate(models.Message{})

	//port is the default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//middleware with corst
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permite todas las solicitudes de cualquier origen
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	//routes http
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.MessageRouter(router)

	handler := c.Handler(router)

	//server running
	log.Println("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, handler)
}
