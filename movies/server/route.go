package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router - вызывает обработчики в зависимости от поступившего запроса
func (serv *Server) Router(route *mux.Router) {
	route.Handle("/favicon.ico", http.FileServer(http.Dir("../www")))
	route.HandleFunc("/movies", serv.moviesListAllHendler)
	route.HandleFunc("/movies/{ID}", serv.movieOneHandler)
}