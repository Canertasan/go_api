package server

import (
	"fmt"
	"go-api/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)

	housesController := controllers.NewHousesController()

	r.Get("/house", housesController.GetAll)         // Retrieves all houses
	r.Get("/house/{id}", housesController.Get)       // Retrieves a house by id
	r.Post("/house", housesController.Create)        // Creates a new house
	r.Delete("/house/{id}", housesController.Delete) // Deletes a house by id
	r.Put("/house/{id}", housesController.Update)    // Updates a house by id

	printRoutes(r)

	return r
}

func printRoutes(r chi.Router) {
	// Iterate through all routes
	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	})
}
