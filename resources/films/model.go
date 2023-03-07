package films

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

type Film struct {
	FilmId      int    `gorm:"type:smallint;primaryKey"`
	Title       string `gorm:"type:varchar(45)"`
	Description string `gorm:"type:text"`
	ReleaseYear int    `gorm:"type:smallint;default:1"`
	LanguageId  int    `gorm:"type:tinyint"`
	Rating      string `gorm:"type:longtext"`
	Length      int    `gorm:"type:bigint"`
}

func (Film) TableName() string {
	return "film"
}

type FilmRequest struct {
	*Film
}

func (f *FilmRequest) Bind(r *http.Request) error {
	log.Println(f.Film.Title)
	if f.Film == nil {
		return errors.New("missing required Actor fields")
	}

	f.Film.Title = strings.ToUpper(f.Film.Title)
	return nil
}

type FilmResponse struct {
	*Film
}

func NewFilmResponse(film *Film) *FilmResponse {
	return &FilmResponse{film}
}

func NewFilmListResponse(films []*Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (f *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
