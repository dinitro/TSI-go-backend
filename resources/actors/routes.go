package actors

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListActors)
	router.Get("/{id}", GetActorByID)
	router.Get("/search", SearchActors)
	router.Get("/{id}/films", ListActorFilms)
	router.Post("/", CreateActor)
	router.Delete("/{id}", DeleteActor)
	router.Patch("/{id}", UpdateActorByID)

	return router
}
