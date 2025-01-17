package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Gouravjoshi986/Golang/ApiUsingGo/internal/middleware"
)

func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes) // to ignore slaches at the end in url
	r.Route("/account",func(router chi.Router){
		// middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins",GetCoinBalance)
	})
}