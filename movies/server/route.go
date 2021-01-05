package server

import (


	"github.com/gorilla/mux"
)

// Router - вызывает обработчики в зависимости от поступившего запроса
func (serv *Server) Router(route *mux.Router) {
	route.HandleFunc("/", serv.moviesListAllHendler)
	route.HandleFunc("/movies/{ID}", serv.movieOneHendler)
	route.HandleFunc("/favicon.ico", serv.faviconHandler)
	route.HandleFunc("/www/style/style.css", serv.styleHehdler)
	route.HandleFunc("/__version__", serv.versionHandler)
	route.HandleFunc("/__heartbeat__", serv.heartbeatHandler)
}
