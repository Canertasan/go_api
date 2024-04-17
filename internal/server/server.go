package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Start(router *chi.Mux) {
	fmt.Println("Starting server on port :1234, you can reach it in here http://localhost:1234/")
	http.ListenAndServe(":1234", router)
}
