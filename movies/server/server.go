package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server - объект сервера
type Server struct {
	db *sql.DB
	lg *logrus.Logger

}


// NewServer - создает новый экземпляр сервера
func NewServer(db *sql.DB, lg *logrus.Logger) *Server{
	return &Server{
		db: db,
		lg: lg,
	}
}

// StartServer - поднимает сервер, обрабатывает ошибку в случае неудачи
func (serv *Server) StartServer() {
	route := mux.NewRouter()
	serv.Router(route)

	logrus.Info("Microservice Movies launched...")
	serv.lg.WithError(http.ListenAndServe(":8080", route)).Fatal("the server doesn`t start!")
}