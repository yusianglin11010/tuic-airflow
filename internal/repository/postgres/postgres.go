package postgres

import (
	"github.com/yusianglin11010/tuic-airflow/internal/database"
	"github.com/yusianglin11010/tuic-airflow/internal/database/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBRepo interface {
	GetMarkersByID(logger *zap.Logger, tx *gorm.DB, id string) ([]model.Marker, error)

	CreateMarker(logger *zap.Logger, tx *gorm.DB, markers []model.Marker) error
}

var _ DBRepo = (*pgRepo)(nil)

type pgRepo struct {
	db *database.DB
}

func NewPGRepo(db *database.DB) DBRepo {
	return &pgRepo{
		db: db,
	}
}

func (repo *pgRepo) GetMarkersByID(logger *zap.Logger, tx *gorm.DB, id string) ([]model.Marker, error) {
	if tx == nil {
		tx = database.GetDB().DB
	}
	markers := []model.Marker{}
	if err := tx.Where("project_id = ?", id).Find(&markers).Error; err != nil {
		logger.Error("find marker fail", zap.Error(err))
		return nil, err
	}
	return markers, nil

}

func (repo *pgRepo) CreateMarker(logger *zap.Logger, tx *gorm.DB, markers []model.Marker) error {
	if tx == nil {
		tx = database.GetDB().DB
	}
	for _, marker := range markers {
		if err := tx.Create(&marker).Error; err != nil {
			logger.Error("create marker fail", zap.Error(err))
			return err
		}
	}
	return nil
}
