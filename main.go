package main

import (
	"go-sms/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
	}

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.MessageRouter(router)

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}
