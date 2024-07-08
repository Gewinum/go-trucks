package library

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

type MapAPI struct {
}

func NewMapAPI() MapAPI {
	return MapAPI{}
}

type TruckersMPResponse struct {
	Data    any  `json:"Data"`
	Success bool `json:"Success"`
}

type PlayerData struct {
	Name       string  `json:"Name"`
	X          float64 `json:"X"`
	Y          float64 `json:"Y"`
	Heading    float64 `json:"Heading"`
	MpId       int64   `json:"MpId"`
	PlayerId   int64   `json:"PlayerId"`
	ServerId   int64   `json:"ServerId"`
	ServerType int64   `json:"ServerType"`
	Time       int64   `json:"Time"`
}

func (mapApi MapAPI) GetOnlinePlayerData(nameOrId string) ([]PlayerData, error) {
	var responseWrapper TruckersMPResponse

	builtUrl := fmt.Sprintf("https://tracker.ets2map.com/v3/playersearch?string=%s", nameOrId)
	_, err := getClient().R().SetResult(&responseWrapper).Get(builtUrl)

	if err != nil {
		return nil, err
	}

	var playerData []PlayerData
	err = mapstructure.Decode(responseWrapper.Data, &playerData)
	if err != nil {
		return nil, errors.New("user wasn't found")
	}

	return playerData, nil
}

func getClient() *resty.Client {
	client := resty.New()
	return client
}
