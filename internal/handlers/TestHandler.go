package handlers

import "net/http"

func TestHandler(w http.ResponseWriter, r *http.Request) {
	response := "Hello World!"

	w.Header().Set("Content-Type", "text=plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}
