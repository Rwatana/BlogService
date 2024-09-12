package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/m/models"
)

func LogStore(w http.ResponseWriter, r *http.Request, messageStore *[]models.Message) {
	if r.Method == "POST" {
		var msg models.Message

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received Message: %+v\n", msg)

		messageStore = append(messageStore, msg)

		fmt.Fprintf(w, "Received: %+v\n", msg)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func LogSend(w http.ResponseWriter, r *http.Request, messageStore *[]models.Message) {
	if r.Method == "GET" {
		messages := messageStore

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(messages); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
