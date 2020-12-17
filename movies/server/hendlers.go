package server

import (
	"context"
	"fmt"
	"html/template"
	"movies/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

<<<<<<< HEAD

func (serv *Server) moviesListAllHendler(w http.ResponseWriter, r *http.Request) {
=======
func (serv *Server) movieListAllHendler(w http.ResponseWriter, r *http.Request) {
>>>>>>> 6415b6e2e52ac5d527d949fd2c4985d1a34b4de8
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

	if err := tmpl.ExecuteTemplate(w, "Cinema", MoviesSlice); err != nil {
		serv.lg.WithError(err).Infoln("can't show all posts")
	}

}

func (serv *Server) movieOneHendler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("Cinema").ParseFiles("./www/templates/moviePage.html"))

	idVars := mux.Vars(r)
	id := idVars["id"]

	ctx := context.Background()

	oneMovie, err := models.Movies(qm.Where("ID=?", id)).One(ctx, serv.db)
	if err != nil {
		serv.lg.WithError(err).Infof("Can't get this movie")
	}

<<<<<<< HEAD
}

func (serv *Server) movieOneHandler(w http.ResponseWriter, r *http.Request) {
	inputID := r.ParseForm()
	fmt.Println(inputID)
}
=======
	if err := tmpl.ExecuteTemplate(w, "Cinema", oneMovie); err != nil {
		serv.lg.WithError(err).Info("Can't show this movie page")
	}
	fmt.Println(id)
}
>>>>>>> 6415b6e2e52ac5d527d949fd2c4985d1a34b4de8
