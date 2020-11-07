package server

import (
	// "net/http"

	"github.com/gorilla/mux"
)

// Router - вызывает обработчики в зависимости от поступившего запроса
func (serv *Server) Router(route *mux.Router) {
	// route.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/favicon.ico", http.FileServer(http.Dir("./www"))))
	route.HandleFunc("/movies", serv.movieListAllHendler)
	route.HandleFunc("/movies/{id}", serv.movieOneHendler)
}
