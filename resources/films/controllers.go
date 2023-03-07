package films

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	db "tsi.co/go-api/database"
	e "tsi.co/go-api/error"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	db.DB.Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func SearchFilms(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	// Get the query to search.
	q := r.URL.Query().Get("fq")

	// Find and render.
	db.DB.Where("title LIKE ?", "%"+q+"%").Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func GetFilmByID(w http.ResponseWriter, r *http.Request) {
	filmID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}
	// Get the film from the database
	var film Film
	result := db.DB.First(&film, filmID)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}

	// Render the film as a response
	render.Render(w, r, NewFilmResponse(&film))
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	film := data.Film
	db.DB.Create(film)
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	// Get the film ID from the URL parameter.
	filmID := chi.URLParam(r, "id")

	var film Film
	db.DB.First(&film, filmID)
	db.DB.Delete(&film)
}

func UpdateFilmByID(w http.ResponseWriter, r *http.Request) {
	// Get the film ID from the URL parameter.
	filmID := chi.URLParam(r, "id")

	var film Film
	db.DB.First(&film, filmID)
	json.NewDecoder(r.Body).Decode(&film)

	// Uppercase
	film.Title = strings.ToUpper(film.Title)

	db.DB.Save(&film)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)
}

func ListFilmsByRating(w http.ResponseWriter, r *http.Request) {
	rt := chi.URLParam(r, "r")

	// Get the film from the database
	var film []*Film
	db.DB.Where("rating = ?", rt).Find(&film)

	// Render the film as a response
	render.RenderList(w, r, NewFilmListResponse(film))
}

func ListFilmsByLength(w http.ResponseWriter, r *http.Request) {
	// Get the params from the URL
	params := r.URL.Query()

	// Get the min and max length. If one of them is empty, set default value.
	minLen := params.Get("min")
	if minLen == "" {
		minLen = "0"
	}
	maxLen := params.Get("max")
	if maxLen < minLen || minLen == "" {
		maxLen = "1000"
	}

	// Get the films from the database.
	var film []*Film
	db.DB.Where("length >= " + minLen + " AND length <= " + maxLen).Find(&film)

	// Render the film as a response.
	render.RenderList(w, r, NewFilmListResponse(film))
}

// Routed. To test.
func SearchFilmsDescription(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	// Get the query to search.
	q := r.URL.Query().Get("desc")

	// Find and render.
	db.DB.Where("description LIKE ?", "%"+q+"%").Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

// Routed. To test.
func UpdateActorByID(w http.ResponseWriter, r *http.Request) {
	// Get the actor ID from the URL parameter.
	filmID := chi.URLParam(r, "id")

	var film Film
	db.DB.First(&film, filmID)
	json.NewDecoder(r.Body).Decode(&film)

	// Uppercase the first name and last name
	film.Title = strings.ToUpper(film.Title)

	db.DB.Save(&film)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)

}
