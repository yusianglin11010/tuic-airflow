package smartcity

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/yusianglin11010/tuic-airflow/internal/database/model"
	"github.com/yusianglin11010/tuic-airflow/internal/domain"
)

type Markers struct {
	XMLName xml.Name       `xml:"markers"`
	Markers []model.Marker `xml:"marker"`
}

func GetMarkers() (*Markers, error) {
	url := domain.MarkersResource

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	rss := &Markers{}
	err = xml.NewDecoder(res.Body).Decode(rss)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rss, nil
}
