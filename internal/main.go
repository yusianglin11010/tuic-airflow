package main

import (
	"encoding/xml"

	"github.com/yusianglin11010/tuic-airflow/internal/config"
	"github.com/yusianglin11010/tuic-airflow/internal/database"
	"github.com/yusianglin11010/tuic-airflow/internal/database/migration"
	"go.uber.org/zap"
)

type Marker struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	// Address string `xml:"address,attr"`
	Lat string `xml:"lat,attr"`
	Lng string `xml:"lng,attr"`
	// Locname string `xml:"locname,attr"`
	// Img     string `xml:"img,attr"`
	// Intro   string `xml:"intro,attr"`
	// Icon    string `xml:"icon,attr"`
	// Type    string `xml:"type,attr"`
}

type Markers struct {
	XMLName xml.Name `xml:"markers"`
	Markers []Marker `xml:"marker"`
}

func main() {
	logger, _ := zap.NewProduction()
	cfg := config.NewDBConfig(logger)
	database.Initialize(cfg)
	migration.Migrate()
}
