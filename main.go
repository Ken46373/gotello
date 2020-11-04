package main

import (
	"gotello/app/models"
	"gotello/config"
	"gotello/utils"
	"time"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	droneManager := models.NewDroneManager()
	droneManager.TakeOff()
	time.Sleep(10 * time.Second)
	droneManager.Patrol()
	time.Sleep(30 * time.Second)
	droneManager.Patrol()
	time.Sleep(10 * time.Second)
	droneManager.Land()
}
