package server

import (
	"database/sql"
	"log"
	"net/http"

	"clients-manager/internal/handlers"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func Start(addr string) {
	db, err := sql.Open("sqlite3", "./db/app")
	if err != nil {
		log.Fatal("Database not loaded: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database not responding: ", err)
	}

	log.Println("Database connected successfully!")

	router := chi.NewRouter()
	router.Get("/test", handlers.TestHandler)

	log.Printf("Server starting on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("Server stopped with error: ", err)
	}
}
