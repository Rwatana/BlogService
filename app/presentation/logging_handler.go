package presentation

import (
    "encoding/json"
    "io"
    "log"
    "github.com/Rwatana/BlogService/app/domain"
)

type LoggingHandler struct {
    repo domain.Repository
}

func NewLoggingHandler(repo domain.Repository) *LoggingHandler {
    return &LoggingHandler{repo: repo}
}

func ParseLogMessage(r io.Reader) ([]*domain.Message, error) {
    var rawMessages []struct {
        LogLevel       string `json:"log_level"`
        Date           string `json:"date"`
        CurrentService string `json:"current_service"`
        SourceService  string `json:"source_service"`
        TypeOfRequest  string `json:"type_of_request"`
        Content        string `json:"content"`
    }

    decoder := json.NewDecoder(r)
    if err := decoder.Decode(&rawMessages); err != nil {
        log.Printf("Error occurred while decoding JSON: %v", err)
        return nil, err
    }

    var messages []*domain.Message
    for _, raw := range rawMessages {
        msg, err := domain.NewMessage(
            raw.LogLevel,
            raw.Date,
            raw.CurrentService,
            raw.SourceService,
            raw.TypeOfRequest,
            raw.Content,
        )
        if err != nil {
            log.Printf("Error occurred while creating message: %v", err)
            continue
        }
        messages = append(messages, msg)
    }

    return messages, nil
}

func (h *LoggingHandler) ParseAndStoreLogMessages(r io.Reader) error {
    messages, err := ParseLogMessage(r)
    if err != nil {
        return err
    }

    for _, msg := range messages {
        if err := h.repo.Save(msg); err != nil {
            log.Printf("Failed to save message: %v", err)
        }
    }

    return nil
}
