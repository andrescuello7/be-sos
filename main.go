package main

import (
	"go-sms/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.MessageRouter(router)

	log.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}
