package server

import (
	"context"
	"fmt"
	"html/template"
	"movies/models"
	"movies/render"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (serv *Server) versionHandler(w http.ResponseWriter, _ *http.Request){
	type Data struct {
		Status string `json:"status"`
		Version string `json:"version"`
		Commit string `json:"commit"`
		Build string `json:"build"`
	}

	data := Data{
		Status: strconv.Itoa(http.StatusOK),
		Version: serv.VersionInfo.Version,
		Commit: serv.VersionInfo.Commit,
		Build: serv.VersionInfo.Build,
	}

	render.RenderJson(w, data)
}

func (serv *Server) heartbeatHandler(w http.ResponseWriter, _ *http.Request) {
	render.RenderJson(w, http.StatusOK)
}


func (serv *Server) moviesListAllHendler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("Cinema").ParseFiles("./www/templates/index.html"))

	ctx := context.Background()

	allMovies, err := models.Movies().All(ctx, serv.db)
	if err != nil {
		serv.lg.WithError(err).Info("Can't get a list of all movies")
	}

	if err := tmpl.ExecuteTemplate(w, "Cinema", allMovies); err != nil {
		serv.lg.WithError(err).Infoln("can't show all posts")
		render.RenderJsonERR(w, err, http.StatusNoContent)
	}
}

func (serv *Server) movieOneHendler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("Cinema").ParseFiles("./www/templates/moviePage.html"))

	idVars := mux.Vars(r)
	id := idVars["ID"]
	fmt.Println(id)

	ctx := context.Background()

	oneMovie, err := models.Movies(qm.Where("ID=?", id)).One(ctx, serv.db)
	if err != nil {
		serv.lg.WithError(err).Infof("Can't get this movie")
	}

	if err := tmpl.ExecuteTemplate(w, "Cinema", oneMovie); err != nil {
		serv.lg.WithError(err).Info("Can't show this movie page")
	}
}

func (serv *Server) faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("www", "favicon.ico"))
}


func (serv *Server) styleHehdler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("www", "style", "style.css"))
}