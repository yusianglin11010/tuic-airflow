package postgres

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/yusianglin11010/tuic-airflow/internal/config"
	"github.com/yusianglin11010/tuic-airflow/internal/database"
	"github.com/yusianglin11010/tuic-airflow/internal/database/model"
	"go.uber.org/zap"
)

func TestCreateMarker(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := config.NewDBConfig(logger)
	database.Initialize(cfg)
	db := database.GetDB()
	pgRepo := NewPGRepo(db)

	randomID := uuid.New().String()
	markers := []model.Marker{
		{
			ProjectID: randomID,
			Name:      "test",
			Lat:       "12.3",
			Lng:       "45.6",
		},
	}

	err := pgRepo.CreateMarker(logger, nil, markers)

	assert.Nil(t, err)

	data := model.Marker{}
	db.Where("project_id = ?", randomID).First(&data)

	assert.Equal(t, data.ProjectID, randomID)
}

func TestGetMarkersByID(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := config.NewDBConfig(logger)
	database.Initialize(cfg)
	db := database.GetDB()
	pgRepo := NewPGRepo(db)

	randomID := uuid.New().String()
	markers := []model.Marker{
		{
			ProjectID: randomID,
			Name:      "test",
			Lat:       "12.3",
			Lng:       "45.6",
		},
		{
			ProjectID: randomID,
			Name:      "test",
			Lat:       "12.3",
			Lng:       "45.6",
		},
	}
	for _, marker := range markers {
		err := db.Create(&marker).Error
		assert.Nil(t, err)
	}

	res, err := pgRepo.GetMarkersByID(logger, nil, randomID)
	assert.Nil(t, err)
	assert.Equal(t, len(markers), len(res))

}
