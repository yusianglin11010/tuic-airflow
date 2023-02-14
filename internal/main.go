package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/tuic-airflow/internal/config"
	"github.com/yusianglin11010/tuic-airflow/internal/database"
	"github.com/yusianglin11010/tuic-airflow/internal/database/migration"
	"github.com/yusianglin11010/tuic-airflow/internal/handler"
	"github.com/yusianglin11010/tuic-airflow/internal/middleware"
	"github.com/yusianglin11010/tuic-airflow/internal/repository/postgres"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	dbConfig := config.NewDBConfig(logger)
	database.Initialize(dbConfig)
	migration.Migrate()

	restConfig := config.NewRestConfig(logger)

	dbRepo := postgres.NewPGRepo(database.GetDB())
	h := handler.NewHandler(dbRepo)

	server := gin.New()
	server.Use(middleware.AddLoggerToContext(logger))

	server.GET("/alive", h.GetHealth)
	server.GET("/project", h.GetProjects)
	server.Run(restConfig.Port)
}
