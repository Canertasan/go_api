package main

import (
	"fmt"
	"net/http"

	"go-api/internal/server"
)

func main() {
	// houses = append(houses, models.House{ID: "1", Name: "House 1", Rooms: 3, Square: 50})

	fmt.Println("Setting up routes...")
	router := server.SetupRoutes()

	fmt.Println("Starting server on port :1234, you can reach it in here http://localhost:1234/")
	http.ListenAndServe(":1234", router)
}
