package controllers

import (
	"encoding/json"
	"go-sms/db"
	"go-sms/models"
	"net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message []models.Message
	getMessage := db.DB.Find(&message)

	err := getMessage.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error getting message"))
	}

	json.NewEncoder(w).Encode(&message)
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message models.Message
	json.NewDecoder(r.Body).Decode(&message)
	createMessage := db.DB.Create(&message)

	err := createMessage.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating message"))
	}

	json.NewEncoder(w).Encode(&message)
}
