package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
    dbUser := "root"
    dbPassword := "lOjit212"
    dbHost := "127.0.0.1"
    dbPort := "3306"
    dbName := "test_db"

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, fmt.Errorf("error connecting to database: %v", err)
    }

    fmt.Println("Successfully connected to the database")
    return db, nil
}
