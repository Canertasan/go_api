package controllers

import (
	"encoding/json"
	"go-api/internal/models"
	"go-api/internal/repository"
	"net/http"

	"github.com/go-chi/chi"
)

type housesController struct {
	repo *repository.HousesRepository
}

func NewHousesController(repo *repository.HousesRepository) *housesController {
	return &housesController{repo: repo}
}

func (h *housesController) Get(w http.ResponseWriter, req *http.Request) {
	idParam := chi.URLParam(req, "id")
	for _, item := range h.repo.Houses {
		if item.ID == idParam {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.House{})
}

func (h *housesController) GetAll(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(h.repo.Houses)
}

func (h *housesController) Create(w http.ResponseWriter, req *http.Request) {
	idParam := chi.URLParam(req, "id")
	var house models.House

	err := json.NewDecoder(req.Body).Decode(&house)
	if err != nil {
		// Write an error message and stop the handler
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	house.ID = idParam
	h.repo.Houses = append(h.repo.Houses, house)

	err = json.NewEncoder(w).Encode(h.repo.Houses)
	if err != nil {
		// Write an error message and stop the handler
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *housesController) Delete(w http.ResponseWriter, req *http.Request) {
	idParam := chi.URLParam(req, "id")
	for index, item := range h.repo.Houses {
		if item.ID == idParam {
			h.repo.Houses = append(h.repo.Houses[:index], h.repo.Houses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(h.repo.Houses)
}

func (h *housesController) Update(w http.ResponseWriter, req *http.Request) {
	idParam := chi.URLParam(req, "id")
	var updatedHouse models.House

	// Handle errors occurring during decoding the request body
	err := json.NewDecoder(req.Body).Decode(&updatedHouse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Unable to parse request body: " + err.Error())
		return
	}

	for index, item := range h.repo.Houses {
		if item.ID == idParam {
			// Updating the House with the data received in the request body
			h.repo.Houses[index] = updatedHouse

			// Handle errors occurring during encoding the response
			err := json.NewEncoder(w).Encode(updatedHouse)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode("Error encoding response: " + err.Error())
				return
			}
			return
		}
	}

	// If we've reached this line, it means the house was not found.
	w.WriteHeader(http.StatusNotFound)

	// Handle errors occurring during encoding the failure message
	err = json.NewEncoder(w).Encode("House not found.")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error encoding response: " + err.Error())
	}
}
