package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error)  {
	db, err := sql.Open("mysql", "game:password@tcp(localhost:3306)/pong")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL database")

	// Set maximum number of open connections
	db.SetMaxOpenConns(10)

	// Set maximum number of idle connections
	db.SetMaxIdleConns(5)

	// Return the connection to the pool after 2 minutes of inactivity
	db.SetConnMaxLifetime(120 * time.Second)


	return db, nil
}
