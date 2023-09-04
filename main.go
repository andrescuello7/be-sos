package main

import (
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

	port := os.Getenv("PORT_ENV")

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.MessageRouter(router)

	log.Println("Server listening on port" + port)
	http.ListenAndServe(":8080", router)
}
