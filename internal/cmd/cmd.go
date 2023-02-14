package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yusianglin11010/tuic-airflow/internal/database"
	"github.com/yusianglin11010/tuic-airflow/internal/repository/postgres"
	"github.com/yusianglin11010/tuic-airflow/internal/repository/smartcity"
	"go.uber.org/zap"
)

var FetchProject = &cobra.Command{
	Use:   "fetch",
	Short: "fetch project data",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		db := database.GetDB()
		dbRepo := postgres.NewPGRepo(db)
		if err := dbRepo.DeleteMarker(logger, nil); err != nil {
			panic(err)
		}
		markers, err := smartcity.GetMarkers()
		if err != nil {
			panic(err)
		}
		if err := dbRepo.CreateMarker(logger, nil, markers.Markers); err != nil {
			panic(err)
		}
	},
}
