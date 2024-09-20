// TODO: Create a presentation layer
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LogHandler struct{}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

func (h *LogHandler) InsertLogMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req InsertLogRequest

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received Message: %+v\n", req)

		msg := Message{
			LogLevel:       req.LogLevel,
			CurrentService: req.CurrentService,
			SourceService:  req.SourceService,
			TypeOfRequest:  req.TypeOfRequest,
			Content:        req.Content,
		}

		// TODO: introduce an infrastructure layer
		mu.Lock()
		messageStore = append(messageStore, msg)
		mu.Unlock()

		// TODO: Consolidate response handling into a common function.
		fmt.Fprintf(w, "Received: %+v\n", msg)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
