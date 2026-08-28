package server

import (
	"log"
	"net/http"

	"clients-manager/internal/handlers"

	"github.com/go-chi/chi"
)

func Start(adr string) {
	router := chi.NewRouter()

	router.Get("/test", handlers.TestHandler)

	err := http.ListenAndServe(adr, router)

	if err != nil {
		log.Println("Error server starting: ", err.Error())
		return
	}
}
