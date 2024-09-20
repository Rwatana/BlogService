// TODO: Create a presentation layer
package main

import "time"

type InsertLogRequest struct {
	LogLevel       string `json:"log_level"`
	Date           string `json:"date"`
	CurrentService string `json:"current_service"`
	SourceService  string `json:"source_service"`
	TypeOfRequest  string `json:"type_of_request"`
	Content        string `json:"content"`
}

// TODO: introduce dtos and domains
type Message struct {
	LogLevel       string    `json:"log_level"`
	Date           time.Time `json:"date"`
	CurrentService string    `json:"current_service"`
	SourceService  string    `json:"source_service"`
	TypeOfRequest  string    `json:"type_of_request"`
	Content        string    `json:"content"`
}
