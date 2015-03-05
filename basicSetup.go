package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	robot := gobot.NewRobot("Eve", func() {
		gobot.Every(500*time.Millisecond, func() {
			fmt.Println("Greeting Human")
		})
	})
	gbot.AddRobot(robot)
	gbot.Start()
}
