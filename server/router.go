package server

import (
	"github.com/go-chi/chi/v5"
	"tsi.co/go-api/resources/actors"
	"tsi.co/go-api/resources/films"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Mount("/films", films.Routes())
	router.Mount("/actors", actors.Routes())

	return router
}
