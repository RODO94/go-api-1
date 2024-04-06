package handlers

import (
	"api-1/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Capital function name means it is exported
// and not private

func Handler(r *chi.Mux){
// Global middleware

	r.Use(chimiddle.StripSlashes) 
	r.Route("/account", func(router chi.Router){

		router.Use(middleware.Authorization)
		router.Get("/coins",GetCoinBalance)
	})
}

