package routes

import (
	"go-sms/controllers"

	"github.com/gorilla/mux"
)

func MessageRouter(routes *mux.Router) {
	routes.HandleFunc("/message", controllers.GetMessage).Methods("GET")
	routes.HandleFunc("/message", controllers.PostMessage).Methods("POST")
}
