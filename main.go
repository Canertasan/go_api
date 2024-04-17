package main

import (
	"fmt"

	"go-api/internal/seeds"
	"go-api/internal/server"
)

func main() {
	seeds.Run()

	fmt.Println("Setting up routes...")
	router := server.SetupRoutes()

	server.Start(router)
}
