package repository

import "go-api/internal/models"

type HousesRepository struct {
	Houses []models.House
}

func NewHousesRepository() *HousesRepository {
	return &HousesRepository{}
}

// func init() {
// 	houses = append(houses, models.House{ID: "1", Name: "House 1", Rooms: 3, Square: 50})
// }
