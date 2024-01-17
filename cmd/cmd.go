package cmd

import (
	"github.com/AntonyIS/notelify-users-service/config"
	"github.com/AntonyIS/notelify-users-service/internal/adapters/app"
	"github.com/AntonyIS/notelify-users-service/internal/adapters/logger"
	"github.com/AntonyIS/notelify-users-service/internal/adapters/repository/postgres"
	"github.com/AntonyIS/notelify-users-service/internal/core/services"
)

func RunService() {
	// Read application environment and load configurations
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// Initialise console and file logger
	consoleFileLogger := logger.NewLogger()
	// Initialize the logging service
	loggerSvc := services.NewLoggerService(&consoleFileLogger)
	// // Postgres Clien
	databaseRepo, err := postgres.NewPostgresClient(*conf, loggerSvc)
	if err != nil {
		loggerSvc.Error(err.Error())
		panic(err)
	}
	newLoggerService := services.NewLoggingService(conf.LOGGER_URL)
	// Initialize the article service
	articleService := services.NewUserManagementService(databaseRepo, newLoggerService)
	// Run HTTP Server
	app.InitGinRoutes(articleService, loggerSvc, *conf)
	loggerSvc.Close()
}
