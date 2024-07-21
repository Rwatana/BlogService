package main

import (
	"Logging/db"
	"Logging/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func sendResultToServer(url string, result map[string]string) error {
	jsonData, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("received non-OK HTTP status: %s, body: %s", resp.Status, body)
	}
	return nil
}

func main() {
	resp, err := http.Get("http://log-intermediate-srv:4006/messages")
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Server returned non-200 status: %d\nBody: %s\n", resp.StatusCode, string(bodyBytes))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	var messages []Message
	err = json.Unmarshal(body, &messages)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	dbConn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	result := make(map[string]string)
	success := true
	for _, msg := range messages {
		err = logger.SendLogToDB(dbConn, msg.LogLevel, msg.CurrentService, msg.SourceService, msg.TypeOfRequest, msg.Content)
		if err != nil {
			log.Printf("Failed to insert log: %v", err)
			success = false
			break
		}
	}

	if success {
		result["status"] = "success"
		result["message"] = "All logs inserted successfully."
	} else {
		result["status"] = "error"
		result["message"] = "Failed to insert some logs into the database."
	}
	err = sendResultToServer("http://log-intermediate-srv:4006/results", result)
	if err != nil {
		log.Fatalf("Failed to send result to server: %v", err)
	}
	defer dbConn.Close()

	fmt.Println("Result sent to server successfully.")
}
