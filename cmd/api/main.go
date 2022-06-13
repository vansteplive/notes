package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vansteplive/notes-app-golang/config"
	"github.com/vansteplive/notes-app-golang/database"
	"github.com/vansteplive/notes-app-golang/pkg/handler"
	"github.com/vansteplive/notes-app-golang/pkg/repository"
	"github.com/vansteplive/notes-app-golang/pkg/service"
	"github.com/vansteplive/notes-app-golang/server"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.InitDB(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("App started ✅")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("\nApp shutting down 🆘")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occured when server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error occured on db connection close: %s", err.Error())
	}
}
