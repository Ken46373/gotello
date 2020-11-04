package main

import (
	"gotello/config"
	"gotello/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")

	/*
		droneManager := models.NewDroneManager()
		droneManager.TakeOff()
		time.Sleep(5 * time.Second)
		droneManager.Land()
	*/
}
