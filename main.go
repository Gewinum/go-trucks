package main

import (
	"fmt"
	"go-tsmp-map-api/library"
	"log/slog"
	"time"
)

func main() {
	tracker := library.NewMapAPI()
	logger := slog.Default()

	library.InitMap()

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

		locationInfo := library.GetLocationAtCoordinates(user[0].X, user[0].Y)

		logger.Info(fmt.Sprintf("Target is at %s, %s", locationInfo.Country, locationInfo.City))

		playTime += 5
		time.Sleep(time.Second * 5)
	}
}
