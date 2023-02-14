package migration

import (
	"fmt"

	"github.com/yusianglin11010/tuic-airflow/internal/database"
	"github.com/yusianglin11010/tuic-airflow/internal/database/model"
)

func Migrate() {
	db := database.GetDB()
	if err := db.AutoMigrate(&model.Marker{}); err != nil {
		panic(fmt.Sprintf("migrate table error %v", err))
	}
}
