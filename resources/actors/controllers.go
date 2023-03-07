package actors

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	db "tsi.co/go-api/database"
	e "tsi.co/go-api/error"
	"tsi.co/go-api/resources/filmactor"
	"tsi.co/go-api/resources/films"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	db.DB.Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func SearchActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	// Get the query to search.
	q := r.URL.Query().Get("aq")

	// Find and render.
	db.DB.Where("first_name LIKE ? OR last_name LIKE ?", "%"+q+"%", "%"+q+"%").Find(&actors)
	// db.DB.Where("first_name = ?", q).Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func GetActorByID(w http.ResponseWriter, r *http.Request) {
	// Get the actor ID from the URL parameter
	actorID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}
	// actorID := chi.URLParam(r, "actorID")

	// Get the actor from the database
	var actor Actor
	result := db.DB.First(&actor, actorID)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}

	// Render the actor as a response
	render.Render(w, r, NewActorResponse(&actor))
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor
	log.Println(actor)
	db.DB.Create(actor)
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {
	// Get the actor ID from the URL parameter.
	actorID := chi.URLParam(r, "id")

	var actor Actor
	db.DB.First(&actor, actorID)
	db.DB.Delete(&actor)
}

func UpdateActorByID(w http.ResponseWriter, r *http.Request) {
	// Get the actor ID from the URL parameter.
	actorID := chi.URLParam(r, "id")

	var actor Actor
	db.DB.First(&actor, actorID)
	json.NewDecoder(r.Body).Decode(&actor)

	// Uppercase the first name and last name
	actor.FirstName = strings.ToUpper(actor.FirstName)
	actor.LastName = strings.ToUpper(actor.LastName)

	db.DB.Save(&actor)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actor)

}

func ListActorFilms(w http.ResponseWriter, r *http.Request) {
	actorID := chi.URLParam(r, "id")

	// Find films that an actor is in
	var filmActors []*filmactor.FilmActor
	db.DB.Where("actor_id = ?", actorID).Find(&filmActors)

	// Find films based on the film IDs
	var movies []*films.Film
	filmIDs := make([]int, 0, len(filmActors))
	for _, filmActor := range filmActors {
		filmIDs = append(filmIDs, filmActor.FilmID)
	}
	db.DB.Where("film_id IN (?)", filmIDs).Find(&movies)

	// Render the films as a response
	render.RenderList(w, r, films.NewFilmListResponse(movies))
}
