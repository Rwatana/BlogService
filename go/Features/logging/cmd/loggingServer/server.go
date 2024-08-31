package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"example.com/m/db"
	"example.com/m/logger"
	"github.com/streadway/amqp"
)

type Message struct {
	LogLevel       string    `json:"log_level"`
	Date           time.Time `json:"date"`
	CurrentService string    `json:"current_service"`
	SourceService  string    `json:"source_service"`
	TypeOfRequest  string    `json:"type_of_request"`
	Content        string    `json:"content"`
}

// TODO https://github.com/Rwatana/BlogService/issues/21
// - [Log] Unlock the Message Store and Send to the Database Promptly
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

func processLogs() {
	mu.Lock()
	defer mu.Unlock()

	if len(messageStore) == 0 {
		log.Println("No messages to process.")
		return
	}

	dbConn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer dbConn.Close()

	result := make(map[string]string)
	success := true

	for _, msg := range messageStore {
		err = logger.SendLogToDB(dbConn, msg.LogLevel, msg.CurrentService, msg.SourceService, msg.TypeOfRequest, msg.Content)
		if err != nil {
			log.Printf("Failed to insert log: %v", err)
			success = false
			break
		}
		// remove message from store
		messageStore = messageStore[1:]
	}

	if success {
		result["status"] = "success"
		result["message"] = "All logs inserted successfully."
		messageStore = []Message{}
	} else {
		result["status"] = "error"
		result["message"] = "Failed to insert some logs into the database."
	}

	fmt.Println("Result sent to server successfully.")
}


func consumeMessages() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"logs", // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		var logData Message
		err := json.Unmarshal(msg.Body, &logData)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		log.Printf("Received log from RabbitMQ: %+v", logData)

		mu.Lock()
		messageStore = append(messageStore, logData)
		mu.Unlock()
	}
}

func main() {
	// Start the server to handle requests
	go func() {
		http.HandleFunc("/data", addData)
		http.HandleFunc("/messages", getMessages)
		http.HandleFunc("/results", resultsHandler)

		log.Println("Server started on :4007")
		log.Fatal(http.ListenAndServe(":4007", nil))
	}()

	// Start RabbitMQ consumer

	// TODO https://github.com/Rwatana/BlogService/issues/20
	// - [Log] Test rabbit mq that send control log traffic
	go consumeMessages()

	// Start a ticker to process logs every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			processLogs()
		}
	}
}
