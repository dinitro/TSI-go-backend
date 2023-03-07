package filmactor

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type FilmActor struct {
	ActorID int `gorm:"column:actor_id;primary_key"`
	FilmID  int `gorm:"column:film_id;primary_key"`
}

func (FilmActor) TableName() string {
	return "film_actor"
}

type FilmActorRequest struct {
	*FilmActor
}

func (a *FilmActorRequest) Bind(r *http.Request) error {
	if a.FilmActor == nil {
		return errors.New("missing required fields")
	}

	return nil
}

type FilmActorResponse struct {
	*FilmActor
}

func NewFilmActorResponse(filmactor *FilmActor) *FilmActorResponse {
	return &FilmActorResponse{filmactor}
}

func NewFilmActorListResponse(filmactors []*FilmActor) []render.Renderer {
	list := []render.Renderer{}
	for _, filmactor := range filmactors {
		list = append(list, NewFilmActorResponse(filmactor))
	}
	return list
}

func (a *FilmActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
