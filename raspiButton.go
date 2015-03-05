package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	// Create a new Gobot, this is first setup
	// Gobot contains robots, trap, Commander, Eventer
	//Gobot is the main type of your Gobot application and contains a collection of Robots,
	//API commands and Events
	gbot := gobot.NewGobot()

	// Create a new Raspi adaptor
	// RaspiAdaptor contains name, revision, 12cLocation, digitalPins, i2cDevice
	// Create new RaspiAdaptor with name "raspi"
	r := raspi.NewRaspiAdaptor("raspi")

	// Create a new ButtonDriver
	// ButtonDriver contains Active, pin, name, halt, interval, connection, gobot.Eventer
	// This button named "button" in the pin #11
	button := gpio.NewButtonDriver(r, "button", "11")

	// Create a new LedDriver
	// LedDriver represents a digital LED
	// LedDriver contains pin, name, connection, high, gobot.Commander
	// This LedDriver named "led", with pin #7
	led := gpio.NewLedDriver(r, "led", "7")

	// Create work for the robot
	work := func() {

		// when the button is pressed (push):
		gobot.On(button.Event("push"), func(data interface{}) {
			fmt.Println("button pressed") // show to the terminal indicates the button is pressed
			led.On()                      // turn on the LED
		})

		// when the button is released
		gobot.On(button.Event("release"), func(data interface{}) {
			fmt.Println("button released") // show to the terminal indicates the button is released
			led.Off()                      // turn off the LED
		})
	}

	// Create a new Robot
	// Robot contains Name, Work, connections, devices, Commander, and Eventer
	// []Connection: Connections which are automatically started and stopped with the robot
	// []Device: Devices which are automaticallu started and stopped with the robot
	// The work routine the robot will execute once all devices and connections have been initialized
	// and started
	// A name will be automatically generated if no name is supplied
	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button, led},
		work,
	)

	// AddRobot adds a new robot to the internal collection of robots. Returns the added robot
	gbot.AddRobot(robot)

	// Start calls the Start method in its collection of robots, and stops all robots on reception of a SIGINT
	// Start will block the execution of your main function until it receives the SIGINT
	gbot.Start()
}
