package main

import (
	"gotello/app/models"
	"time"
)

func main() {
	droneManager := models.NewDroneManager()
	droneManager.TakeOff()
	time.Sleep(5 * time.Second)
	droneManager.Land()
}
