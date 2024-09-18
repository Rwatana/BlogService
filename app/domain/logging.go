package logging

import (
	"encoding/json"
	"io"
	"log"
)

type Message struct {
	LogLevel       string `json:"log_level"`
	Date           string `json:"date"`
	CurrentService string `json:"current_service"`
	SourceService  string `json:"source_service"`
	TypeOfRequest  string `json:"type_of_request"`
	Content        string `json:"content"`
}

var messageStore []Message

// ParseLogMessage parses JSON log messages from the provided io.Reader and appends them to messageStore
func ParseLogMessage(r io.Reader) []Message {
	var result []Message
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&result); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		return nil
	}

	for _, logData := range result {
		log.Printf("Received log from RabbitMQ: %+v", logData)
		messageStore = append(messageStore, logData)
	}
	return result
}
