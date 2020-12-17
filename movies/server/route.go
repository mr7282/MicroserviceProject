package server

import (
	// "net/http"

	"github.com/gorilla/mux"
)

// Router - вызывает обработчики в зависимости от поступившего запроса
func (serv *Server) Router(route *mux.Router) {
<<<<<<< HEAD
	route.Handle("/favicon.ico", http.FileServer(http.Dir("../www")))
	route.HandleFunc("/movies", serv.moviesListAllHendler)
	route.HandleFunc("/movies/{ID}", serv.movieOneHandler)
}
=======
	// route.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/favicon.ico", http.FileServer(http.Dir("./www"))))
	route.HandleFunc("/movies", serv.movieListAllHendler)
	route.HandleFunc("/movies/{id}", serv.movieOneHendler)
}
>>>>>>> 6415b6e2e52ac5d527d949fd2c4985d1a34b4de8
