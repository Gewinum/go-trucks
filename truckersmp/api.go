package truckersmp

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"go-tsmp-map-api/truckersmp/models"
)

type API interface {
	FetchPlayerInformation(id string) (*models.Player, error)
	FetchPlayerBans(id string) (*[]models.PlayerBan, error)
	GetServerList() (*[]models.Server, error)
	GetGameTime() (*int64, error)
	FetchAllEvents() (*models.EventList, error)
	FetchSpecifiedEvent(eventId string) (*models.Event, error)
	FetchPlayerEvents(userId string) ([]models.Event, error)
	FetchVTCs() (*models.VTCList, error)
	FetchSpecifiedVTC(id string) (*models.DetailedVTCInfo, error)
	FetchVTCNews(id string) (*models.VTCNewsList, error)
	FetchSpecifiedVTCNews(id, newsId string) (*models.VTCDetailedNews, error)
	FetchVTCRoleList(id string) (*models.VTCRoleList, error)
	FetchSpecifiedVTCRole(id, roleId string) (*models.VTCRole, error)
	FetchVTCMemberList(id string) (*models.VTCMemberList, error)
	FetchSpecifiedVTCMember(id, userId string) (*models.VTCPlayerInfo, error)
	FetchVTCEventList(id string) (*[]models.Event, error)
	FetchSpecifiedVTCEvent(id, eventId string) (*models.Event, error)
	FetchGameInformation() (*models.GameInfo, error)
	FetchGameRules() (*models.RulesInfo, error)
}

type truckersMPAPI struct {
	apiPath string
}

func NewAPI() API {
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

func (api truckersMPAPI) FetchVTCs() (*models.VTCList, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetResult(&responseWrapper).Get(api.apiPath + "/vtc")

	if err != nil {
		return nil, err
	}

	var list models.VTCList

	err = mapstructure.Decode(responseWrapper.Response, &list)

	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (api truckersMPAPI) FetchSpecifiedVTC(id string) (*models.DetailedVTCInfo, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": id,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var detailedVTCData models.DetailedVTCInfo

	err = mapstructure.Decode(responseWrapper.Response, &detailedVTCData)
	if err != nil {
		return nil, err
	}

	return &detailedVTCData, nil
}

func (api truckersMPAPI) FetchVTCNews(id string) (*models.VTCNewsList, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": id,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/news")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcNews models.VTCNewsList

	err = mapstructure.Decode(responseWrapper.Response, &vtcNews)
	if err != nil {
		return nil, err
	}

	return &vtcNews, nil
}

func (api truckersMPAPI) FetchSpecifiedVTCNews(id, newsId string) (*models.VTCDetailedNews, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id":      id,
		"news_id": newsId,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/news/{news_id}")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcNews models.VTCDetailedNews

	err = mapstructure.Decode(responseWrapper.Response, &vtcNews)
	if err != nil {
		return nil, err
	}

	return &vtcNews, nil
}

func (api truckersMPAPI) FetchVTCRoleList(id string) (*models.VTCRoleList, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": id,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/roles")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcNews models.VTCRoleList

	err = mapstructure.Decode(responseWrapper.Response, &vtcNews)
	if err != nil {
		return nil, err
	}

	return &vtcNews, nil
}

func (api truckersMPAPI) FetchSpecifiedVTCRole(id, roleId string) (*models.VTCRole, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id":      id,
		"role_id": roleId,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/role/{role_id}")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcNews models.VTCRole

	err = mapstructure.Decode(responseWrapper.Response, &vtcNews)
	if err != nil {
		return nil, err
	}

	return &vtcNews, nil
}
func (api truckersMPAPI) FetchVTCMemberList(id string) (*models.VTCMemberList, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": id,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/members")

	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcMemberList models.VTCMemberList

	err = mapstructure.WeakDecode(responseWrapper.Response, &vtcMemberList)

	if err != nil {
		return nil, err
	}

	return &vtcMemberList, nil
}

func (api truckersMPAPI) FetchSpecifiedVTCMember(id, memberId string) (*models.VTCPlayerInfo, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id":        id,
		"member_id": memberId,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/member/{member_id}")

	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcPlayerInfo models.VTCPlayerInfo

	err = mapstructure.WeakDecode(responseWrapper.Response, &vtcPlayerInfo)

	if err != nil {
		return nil, err
	}

	return &vtcPlayerInfo, nil
}

func (api truckersMPAPI) FetchVTCEventList(id string) (*[]models.Event, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id": id,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/events")
	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcEventList []models.Event

	err = mapstructure.Decode(responseWrapper.Response, &vtcEventList)
	if err != nil {
		return nil, err
	}

	return &vtcEventList, nil
}
func (api truckersMPAPI) FetchSpecifiedVTCEvent(id, eventId string) (*models.Event, error) {
	var responseWrapper models.UsualResponseWrapper

	_, err := getClient().R().SetPathParams(map[string]string{
		"id":       id,
		"event_id": eventId,
	}).SetResult(&responseWrapper).Get(api.apiPath + "/vtc/{id}/events/{event_id}")

	if err != nil {
		return nil, err
	}

	if responseWrapper.Error == true {
		return nil, errors.New(*responseWrapper.Descriptor)
	}

	var vtcEvent models.Event

	err = mapstructure.Decode(responseWrapper.Response, &vtcEvent)
	if err != nil {
		return nil, err
	}

	return &vtcEvent, nil
}

func (api truckersMPAPI) FetchGameInformation() (*models.GameInfo, error) {
	var gameInfo models.GameInfo

	_, err := getClient().R().SetResult(&gameInfo).Get(api.apiPath + "/version")

	if err != nil {
		return nil, err
	}

	return &gameInfo, nil
}

func (api truckersMPAPI) FetchGameRules() (*models.RulesInfo, error) {
	var rulesInfo models.RulesInfo

	_, err := getClient().R().SetResult(&rulesInfo).Get(api.apiPath + "/rules")
	if err != nil {
		return nil, err
	}

	return &rulesInfo, nil
}

func getClient() *resty.Client {
	return resty.New()
}
