package main

import (
	"machine"

	"image/color"
	"time"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinydraw"
)

const (
	startx       = 10
	starty       = 20
	radius int16 = 6
)

func handleDisplay() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display := uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation:    uc8151.ROTATION_270,
		Speed:       uc8151.TURBO,
		FlickerFree: true,
	})

	display.ClearDisplay()

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	for {
		for i, key := range keys {
			tinydraw.FilledCircle(&display, startx+radius+32*int16(i), 48-radius-1, radius, black)
			if key.pressed {
				tinydraw.FilledCircle(&display, startx+radius+32*int16(i), 48-radius-1, radius, white)
			} else {
				tinydraw.Circle(&display, startx+radius+32*int16(i), 48-radius-1, radius, white)
			}
		}

		display.Display()

		time.Sleep(200 * time.Millisecond)
	}
}
