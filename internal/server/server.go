package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Start(addr string) error {
	//database connection
	db, err := sql.Open("sqlite3", "./db/app")
	if err != nil {
		return fmt.Errorf("database not loaded: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("database not responding: %w", err)
	}
	log.Println("Database connected successfully!")

	//starting server
	log.Printf("Server starting on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, Router()); err != nil {
		return fmt.Errorf("server stopped with error: %w", err)
	}

	return nil
}
