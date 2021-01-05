package main

import (
	"database/sql"
	"movies/config"
	"movies/server"
	"movies/version"
	"os"

	"github.com/sirupsen/logrus"
	// "honnef.co/go/tools/config"

	_ "github.com/go-sql-driver/mysql"
)

// NewLogger - Создаёт новый логгер
func NewLogger() *logrus.Logger {
	lg := logrus.New()
	lg.SetReportCaller(false)
	lg.SetFormatter(&logrus.TextFormatter{})
	lg.SetLevel(logrus.DebugLevel)
	return lg
}


func main() {
	lg:= NewLogger()

	launchMode := config.LaunchMode(os.Getenv("LAUNCH_MODE"))
	if len(launchMode) == 0 {
		launchMode = config.LocalEnv
	}
	lg.Infof("LAUNCH MODE: %v", launchMode)

	cfg, err := config.Load(launchMode, "./config")
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("CONFIG: %+v", cfg)

	info := server.VersionInfo{
		Version: version.Version,
		Build: version.Build,
		Commit: version.Commit,
	}

	db, err := sql.Open("mysql", "root:1240608@tcp(88.210.21.76:3360)/cinema?parseTime=true")
	if err != nil {
		lg.WithError(err).Fatal("can`t connect to DataBase")
	}
	defer db.Close()


	serv := server.NewServer(db, lg, info, cfg.Port)
	go func ()  {
		serv.StartServer()
	}()

	for {}

	//TODO нужно изучить и реализовать остановку сервера с помощью системных прерываний через канал!!!
}


