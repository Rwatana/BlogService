package logger

import (
	"database/sql"
	"fmt"
	"time"
)


const (
    DB_NAME = "test_db3"
)

func SendLogToDB(db *sql.DB, log_level, currentService, sourceService, api, errorMessage string) error {
	currentTime := time.Now()
	query := "INSERT INTO test_log3 (log_level, date, current_service, source_service, type_of_request, content) VALUES (?, ?, ?, ?, ?, ?)"
	_ , err := db.Exec(query, log_level, currentTime, currentService, sourceService, api, errorMessage)
	if err != nil {
		return fmt.Errorf("error inserting log into database: %v", err)
	}

	fmt.Printf("Inserted log successfully")
	return nil
}
