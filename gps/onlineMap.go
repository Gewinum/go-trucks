package gps

import (
	"encoding/json"
	"go-tsmp-map-api/utils"
	"math"
	"os"
)

type MapPoint struct {
	PointType string     `json:"type"`
	Name      string     `json:"name"`
	X         float64    `json:"x"`
	Y         float64    `json:"y"`
	Children  []MapPoint `json:"pois"`
}

type LocationInfo struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

var Map []MapPoint

func InitMap() {
	downloadMap()
}

func GetLocationAtCoordinates(x, y float64) LocationInfo {
	var minimalDistance float64
	var chosenCountry MapPoint
	var chosenCity MapPoint

	minimalDistance = 100000000
	for _, countryPoint := range Map {
		for _, cityPoint := range countryPoint.Children {
			calculated := math.Sqrt(math.Pow(cityPoint.X-x, 2) + math.Pow(cityPoint.Y-y, 2))
			if calculated < minimalDistance {
				minimalDistance = calculated
				chosenCountry = countryPoint
				chosenCity = cityPoint
			}
		}
	}

	return LocationInfo{
		City:    chosenCity.Name,
		Country: chosenCountry.Name,
	}
}

func downloadMap() {
	if _, err := os.Stat("./map.json"); err != nil {
		err = utils.DownloadFile("./map.json", "https://map.truckersmp.com/locations_ets2.min.json?v=e9d3436c")
		if err != nil {
			panic(err)
		}
	}

	contentBytes, err := os.ReadFile("./map.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(contentBytes, &Map)

	if err != nil {
		panic(err)
	}
}
