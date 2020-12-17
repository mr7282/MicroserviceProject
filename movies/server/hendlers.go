package server

import (
	"context"
	"fmt"
	"html/template"
	"movies/models"
	"net/http"

	"github.com/volatiletech/sqlboiler/boil"
)


func (serv *Server) moviesListAllHendler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("Cinema").ParseFiles("./www/templates/index.html"))


	ctx := context.Background()

	allMovies, err := models.Movies().All(ctx, serv.db)
	if err != nil {
		serv.lg.WithError(err).Info("Can't get a list of all movies")
	}

	var MoviesSlice []models.Movie
	for _, movie := range allMovies {
		MoviesSlice = append(MoviesSlice, *movie)
	}

	var data = struct{
		Movies []models.Movie
	}{
		Movies: MoviesSlice,
	}

	boil.DebugMode = true

	if err := tmpl.ExecuteTemplate(w, "Cinema", data); err != nil {
		serv.lg.WithError(err).Infoln("can't show all posts")
	}

}

func (serv *Server) movieOneHandler(w http.ResponseWriter, r *http.Request) {
	inputID := r.ParseForm()
	fmt.Println(inputID)
}