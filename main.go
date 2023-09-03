package main

import (
	"go-sms/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routes.MessageRouter(router)

	log.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", router)
}
