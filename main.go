package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8889")

	work := func() {
		drone.TakeOff()

		gobot.After(10*time.Second, func() {
			drone.FrontFlip()
		})

		gobot.After(20*time.Second, func() {
			drone.BackFlip()
		})

		gobot.After(30*time.Second, func() {
			drone.RightFlip()
		})

		gobot.After(40*time.Second, func() {
			drone.LeftFlip()
		})

		gobot.After(50*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work)

	robot.Start()
}
