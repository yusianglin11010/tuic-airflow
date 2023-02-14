package database

import (
	"fmt"

	"github.com/yusianglin11010/tuic-airflow/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var db DB

func Initialize(cfg *config.DBConfig) {
	db.initialize(cfg)
}

func (d *DB) initialize(cfg *config.DBConfig) {
	if d.DB == nil {
		db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("connect to db error: %v", err))
		}
		d.DB = db
	}
}

func GetDB() *DB {
	return &db
}

func End() {
	db.close()
}

func (d *DB) close() {

	sql, err := d.DB.DB()
	if err != nil {
		panic(err)
	}
	sql.Close()
	d.DB = nil

}
