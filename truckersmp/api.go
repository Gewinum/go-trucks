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
	GetServerList() (*[]models.Server, error)
	GetGameTime() (*int64, error)
	FetchAllEvents() (*models.EventList, error)
	FetchSpecifiedEvent(eventId string) (*models.Event, error)
	FetchPlayerEvents(userId string) ([]models.Event, error)
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

func (api truckersMPAPI) GetServerList() (*[]models.Server, error) {
	var responseWrapper models.ServerResponseWrapper

	_, err := getClient().R().SetResult(&responseWrapper).Get(api.apiPath + "/servers")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var serverData []models.Server

	err = mapstructure.Decode(responseWrapper.Response, &serverData)

	if err != nil {
		return nil, err
	}

	return &serverData, nil
}

func (api truckersMPAPI) GetGameTime() (*int64, error) {
	var response models.ServerGameTimeResponse

	_, err := getClient().R().SetResult(&response).Get(api.apiPath + "/game_time")
	if err != nil {
		return nil, err
	}

	return &response.GameTime, nil
}

func (api truckersMPAPI) FetchAllEvents() (*models.EventList, error) {
	var response models.UsualResponseWrapper

	_, err := getClient().R().SetResult(&response).Get(api.apiPath + "/events")

	if err != nil {
		return nil, err
	}

	if response.Error == true {
		return nil, errors.New(*response.Descriptor)
	}

	var eventData models.EventList
	// due to truckersmp api weirdness :/
	err = mapstructure.WeakDecode(response.Response, &eventData)
	if err != nil {
		return nil, err
	}

	return &eventData, nil
}

func (api truckersMPAPI) FetchSpecifiedEvent(eventId string) (*models.Event, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": eventId,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/events/{id}")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var eventData models.Event

	//yup..
	err = mapstructure.WeakDecode(responseWrapper.Response, &eventData)

	if err != nil {
		return nil, err
	}

	return &eventData, nil
}

func (api truckersMPAPI) FetchPlayerEvents(userId string) ([]models.Event, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": userId,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/events/user/{id}")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var eventData []models.Event

	//yup..
	err = mapstructure.WeakDecode(responseWrapper.Response, &eventData)

	if err != nil {
		return nil, err
	}

	return eventData, nil
}

func getClient() *resty.Client {
	return resty.New()
}
