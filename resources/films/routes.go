package films

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilms)
	router.Get("/{id}", GetFilmByID)
	router.Get("/rating/{r}", ListFilmsByRating)
	router.Get("/length", ListFilmsByLength)
	router.Get("/search", SearchFilms)
	router.Get("/search/description", SearchFilmsDescription)
	router.Post("/", CreateFilm)
	router.Delete("/{id}", DeleteFilm)

	return router
}
