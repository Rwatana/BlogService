package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	LogLevel       string    `json:"log_level"`
	Date           time.Time `json:"date"`
	CurrentService string    `json:"current_service"`
	SourceService  string    `json:"source_service"`
	TypeOfRequest  string    `json:"type_of_request"`
	Content        string    `json:"content"`
}

var (
	messageStore []Message
	mu           sync.Mutex
)

func addData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var msg Message

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

		mu.Lock()
		messageStore = append(messageStore, msg)
		mu.Unlock()

		fmt.Fprintf(w, "Received: %+v\n", msg)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mu.Lock()
		messages := messageStore
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(messages); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var result map[string]string
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			log.Printf("Failed to decode request body: %v", err)
			http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
			return
		}

		log.Printf("Received result: %v", result)

		if len(result) == 0 {
			log.Println("Received empty result.")
			http.Error(w, "Empty result received", http.StatusBadRequest)
			return
		}

		status, ok := result["status"]
		if !ok {
			log.Println("Status not found in the result.")
			http.Error(w, "Status not found in the result.", http.StatusBadRequest)
			return
		}

		if status == "success" {
			mu.Lock()
			messageStore = nil
			mu.Unlock()
			log.Println("Messages cleared successfully.")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Messages cleared successfully."))
		} else {
			log.Printf("Failed to insert log into database, status: %s", status)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to insert log into database."))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/data", addData)

	http.HandleFunc("/messages", getMessages)

	http.HandleFunc("/results", resultsHandler)

	log.Println("Server started on :4006")

	log.Fatal(http.ListenAndServe(":4006", nil))
}
