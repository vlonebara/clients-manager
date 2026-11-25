package server

import (
	"clients-manager/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func Router() http.Handler {

	router := chi.NewRouter()
	router.Get("/test", handlers.TestHandler)

	return router
}
