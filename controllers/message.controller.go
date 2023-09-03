package controllers

import (
	"encoding/json"
	"fmt"
	"go-sms/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func PostMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var messageModel models.Message
	json.NewDecoder(r.Body).Decode(&messageModel)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
	}

	url := os.Getenv("URL_PATH")
	method := "POST"
	message := messageModel.Message
	peticion := "To=whatsapp%3A%2B" + os.Getenv("TO_PHONE") + "&From=whatsapp%3A%2B" + os.Getenv("FROM_PHONE") + "&Body=" + message

	payload := strings.NewReader(peticion)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic QUNlNjg3OGNmZWI0OGVjNDUxZDYzZDg2MjI0YmY3ZTE5YzoxOTZkOGE1ZDM0MGYyMjVjNTFmNGUyYjZkYjc5NzE3Nw==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	ioutil.ReadAll(res.Body)

	json.NewEncoder(w).Encode(&messageModel)
}
