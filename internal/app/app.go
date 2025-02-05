package app

import (
	"actualizer/internal/common/config"
	"actualizer/internal/handler"
	"actualizer/internal/repository"
	"actualizer/internal/repository/postgres"
	"actualizer/internal/server"
	"actualizer/internal/service"
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

func Run() error {
	config.LoadConfig()
	ctx := context.Background()
	serv := new(server.Server)

	connection, err := postgres.NewConnection().Connect(config.AppConfig.DBConfig, ctx)
	if err != nil {
		logrus.Fatalf("error occured while сonnecting DB: %s", err.Error())
		return err
	}

	repo := repository.New(connection)
	scheduler := service.NewService(repo)

	handlers := handler.New(scheduler)
	if err = serv.Run(os.Getenv("API_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running Http-Server: %s", err.Error())
		return err
	}
	return nil
}
