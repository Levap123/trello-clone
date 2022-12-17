package main

import (
	"github.com/Levap123/trello-clone/configs"
	"github.com/Levap123/trello-clone/internal/repository"
	"github.com/Levap123/trello-clone/internal/service"
	"github.com/Levap123/trello-clone/internal/transport/rest"
	"github.com/Levap123/trello-clone/pkg/logger"
	"github.com/Levap123/trello-clone/pkg/server"
)

func main() {
	dbCfg := configs.NewDbConfigs()
	serverCfg := configs.NewServerCgf()
	logger := logger.NewLogger()
	db, err := repository.InitDb(dbCfg)
	if err != nil {
		logger.Err.Fatalln(err.Error())
	}
	if err := db.Ping(); err != nil {
		logger.Err.Fatalln(err.Error())
	}
	repo := repository.NewRepo(db)
	service := service.NewService(repo)
	handler := rest.NewHandler(service, logger)
	server := new(server.Server)
	logger.Info.Printf("server is listening on: http://localhost%s\n", serverCfg.Address)
	if err := server.Run(serverCfg.Address, handler.InitRoutes()); err != nil {
		logger.Err.Fatalln(err.Error())
	}
}
