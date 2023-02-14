package model

type Marker struct {
	ProjectID   string `gorm:"column:project_id" xml:"id,attr"`
	Name string `gorm:"column:name" xml:"name,attr"`
	Lat  string `gorm:"column:lat" xml:"lat,attr"`
	Lng  string `gorm:"column:lng" xml:"lng,attr"`
}
