package truckersmp

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"go-tsmp-map-api/truckersmp/models"
)

type TruckersMPAPI interface {
	FetchPlayerInformation(id string) (*models.Player, error)
	FetchPlayerBans(id string) (*[]models.PlayerBan, error)
}

type truckersMPAPI struct {
	apiPath string
}

func NewAPI() TruckersMPAPI {
	return &truckersMPAPI{
		apiPath: "https://api.truckersmp.com/v2",
	}
}

func (api truckersMPAPI) FetchPlayerInformation(id string) (*models.Player, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": id,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/player/{id}")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var userData models.Player
	err = mapstructure.Decode(responseWrapper.Response, &userData)

	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func (api truckersMPAPI) FetchPlayerBans(id string) (*[]models.PlayerBan, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": id,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/bans/{id}")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var bansData []models.PlayerBan

	err = mapstructure.Decode(responseWrapper.Response, &bansData)

	if err != nil {
		return nil, err
	}

	return &bansData, nil
}

func getClient() *resty.Client {
	return resty.New()
}
