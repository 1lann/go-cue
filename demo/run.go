package main

import (
	"fmt"
	"github.com/1lann/go-cue"
	"image/color"
	"time"
)

func main() {
	details, err := cue.Initialize()

	fmt.Println("\nCUE SDK and server details:\n")
	fmt.Println("SDK version: " + details.SDKVersion)
	fmt.Println("Server version: " + details.ServerVersion)
	if details.BreakingChanges {
		fmt.Println("Has breaking changes: yes")
	} else {
		fmt.Println("Has breaking changes: no")
	}

	if err != nil {
		fmt.Println("\nThe following errors were encountered:")
		fmt.Println(err)
	} else {
		fmt.Println("\nNo errors were detected.")
	}

	fmt.Println("\n------------------------------------------------------------\n")

	devices, err := cue.GetAllDevices()
	fmt.Println(len(devices), "device(s) were detected.\n")

	for i, device := range devices {
		fmt.Println("Device #", i)
		fmt.Println("Type:", device.Type)
		fmt.Println("Model:", device.Model, "\n")
	}

	if err != nil {
		fmt.Println("The following errors were encountered:")
		fmt.Println(err)
	} else {
		fmt.Println("No errors were detected.")
	}

	fmt.Println("\n------------------------------------------------------------\n")

	leds, err := cue.InitializeLeds()

	fmt.Println("\n------------------------------------------------------------\n")

	fmt.Println(leds, "leds were detected.\n")

	if err != nil {
		fmt.Println("The following errors were encountered:")
		fmt.Println(err)
	} else {
		fmt.Println("No errors were detected.")
	}

	fmt.Println("\n------------------------------------------------------------\n")

	led, found := cue.GetLedById(cue.CLK_P)

	if !found {
		fmt.Println("The P key could not be found!")
	} else {
		fmt.Println("The P key will be attempted to be lit up with my favorite shade of green.")
		if err = led.SetColor(color.RGBA{0x4c, 0xaf, 0x50, 0xff}); err != nil {
			fmt.Println("The following errors were encountered:")
			fmt.Println(err)
		} else {
			fmt.Println("No errors were detected.")
		}
	}

	fmt.Println("\n------------------------------------------------------------\n")

	fmt.Println("-- END OF DEMO TEST --")

	time.Sleep(time.Minute)
}
