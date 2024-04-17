package repositories

import "go-api/internal/models"

type HousesRepository struct {
	Houses []models.House
}

func NewHousesRepository() *HousesRepository {
	return &HousesRepository{}
}

func (r *HousesRepository) GetByID(id string) *models.House {
	for _, item := range r.Houses {
		if item.ID == id {
			return &item
		}
	}
	return &models.House{}
}

func (r *HousesRepository) GetAll() []models.House {
	return r.Houses
}

func (r *HousesRepository) Create(house models.House) {
	r.Houses = append(r.Houses, house)
}

func (r *HousesRepository) Delete(id string) {
	for index, item := range r.Houses {
		if item.ID == id {
			r.Houses = append(r.Houses[:index], r.Houses[index+1:]...)
			break
		}
	}
}

func (r *HousesRepository) Update(id string, house models.House) {
	for index, item := range r.Houses {
		if item.ID == id {
			r.Houses[index] = house
			break
		}
	}
}

// func init() {
// 	houses = append(houses, models.House{ID: "1", Name: "House 1", Rooms: 3, Square: 50})
// }
