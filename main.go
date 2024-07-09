package main

import (
	"fmt"
	"go-tsmp-map-api/gps"
	"go-tsmp-map-api/truckersmp"
	"log/slog"
	"time"
)

func main() {
	tracker := gps.NewMapAPI()
	logger := slog.Default()

	api := truckersmp.NewAPI()

	/*events, err := api.FetchAllEvents()

	if err != nil {
		panic(err)
	}

	fmt.Println(events.Today[0].ID)*/

	/*event, err := api.FetchSpecifiedEvent("18061")

	if err != nil {
		panic(err)
	}

	fmt.Println(event)*/

	events, err := api.FetchPlayerEvents("558392")

	if err != nil {
		panic(err)
	}

	fmt.Println(events)

	/*user, err := api.FetchPlayerBans("5271915")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	servers, err := api.GetServerList()

	if err != nil {
		panic(err)
	}

	fmt.Println(servers)

	gameTime, err := api.GetGameTime()

	if err != nil {
		panic(err)
	}

	fmt.Println(gameTime)*/

	gps.InitMap()

	chosenId := "anything-you-want"

	playTime := 0

	for {
		user, err := tracker.GetOnlinePlayerData(chosenId)

		if err != nil {
			if playTime != 0 {
				logger.Info(fmt.Sprintf("Target has played for %d seconds!", playTime))
			}
			playTime = 0
			time.Sleep(time.Second * 5)
			continue
		}

		if user == nil {
			if playTime != 0 {
				logger.Info(fmt.Sprintf("Target has played for %d seconds!", playTime))
			}
			playTime = 0
			time.Sleep(time.Second * 5)
			continue
		}

		if len(user) == 0 {
			if playTime != 0 {
				logger.Info(fmt.Sprintf("Target has played for %d seconds!", playTime))
			}
			playTime = 0
			time.Sleep(time.Second * 5)
			continue
		}

		if playTime == 0 {
			logger.Info("Target has joined the server!")
		}

		locationInfo := gps.GetLocationAtCoordinates(user[0].X, user[0].Y)

		logger.Info(fmt.Sprintf("Target is at %s, %s", locationInfo.Country, locationInfo.City))

		playTime += 5
		time.Sleep(time.Second * 5)
	}
}
