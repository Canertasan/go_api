package main

import (
	"fmt"

	"go-api/internal/repositories"
	"go-api/internal/server"
)

func main() {
	// Create a new repository
	repo := repositories.NewHousesRepository()

	// houses = append(houses, models.House{ID: "1", Name: "House 1", Rooms: 3, Square: 50})
	seed(repo)

	fmt.Println("Setting up routes...")
	router := server.SetupRoutes(repo)

	server.Start(router)
}
