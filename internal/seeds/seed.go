package seeds

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

func Run() {
	repositories.HousesRepo.Create(models.House{ID: "1", Name: "House 1", Rooms: 3, Square: 50})
}
