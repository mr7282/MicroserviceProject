package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server - объект сервера
type Server struct {
	VersionInfo

	db *sql.DB
	lg *logrus.Logger

	port string

}

//VersionInfo - описывает версию сборки
type VersionInfo struct {
	Version string
	Commit string
	Build string
}


// NewServer - создает новый экземпляр сервера
func NewServer(db *sql.DB, lg *logrus.Logger, info VersionInfo, port string) *Server{
	return &Server{
		VersionInfo: info,
		port: port,
		db: db,
		lg: lg,
	}
}

// StartServer - поднимает сервер, обрабатывает ошибку в случае неудачи
func (serv *Server) StartServer() {
	route := mux.NewRouter()
	serv.Router(route)

	logrus.Info("Microservice Movies launched...")
	serv.lg.WithError(http.ListenAndServe(serv.port, route)).Fatal("the server doesn`t start!")
}