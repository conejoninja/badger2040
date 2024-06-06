package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display := uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.MEDIUM,
		Blocking: true,
	})

	// Clear the screen to black
	display.ClearBuffer()

	// Write "Hello" 10 pixels from the right and 50 pixels from the top, in mellon green
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 50, "Hello", color.RGBA{R: 1, G: 1, B: 1, A: 255})
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 40, 80, "Gophers!", color.RGBA{R: 1, G: 1, B: 1, A: 255})

	// Show the buffer on the screen
	display.Display()
}
