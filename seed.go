package main

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

func seed(repo *repositories.HousesRepository) {
	// repositories.NewHousesRepository()
	// houses = append(houses, models.House{ID: "1", Name: "House 1", Rooms: 3, Square: 50})
	repo.Create(models.House{ID: "1", Name: "House 1", Rooms: 3, Square: 50})
}
