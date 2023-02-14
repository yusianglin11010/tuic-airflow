package main

import (
	"github.com/spf13/cobra"
	"github.com/yusianglin11010/tuic-airflow/internal/cmd"
	"github.com/yusianglin11010/tuic-airflow/internal/config"
	"github.com/yusianglin11010/tuic-airflow/internal/database"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	dbConfig := config.NewDBConfig(logger)
	database.Initialize(dbConfig)
	defer database.End()
	// defer database.Close()
	rootCmd := &cobra.Command{Use: "./app-name"}
	rootCmd.AddCommand(cmd.FetchProject)
	rootCmd.Execute()
}
