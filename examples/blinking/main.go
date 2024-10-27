package main

import (
	//"gotchi/assets"
	"machine"
	"time"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/uc8151"

	"github.com/conejoninja/badger2040/examples/blinking/assets"
)

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display := uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation:    uc8151.ROTATION_270,
		Speed:       uc8151.MEDIUM,
		Blocking:    true,
		FlickerFree: true,
	})

	display.ClearDisplay()

	animation := [][]uint8{
		assets.Idle1,
		assets.Idle2,
		assets.Idle3,
	}

	var current int

	for {
		img := pixel.NewImageFromBytes[pixel.Monochrome](64, 64, []byte(animation[current]))
		display.DrawBitmap(0, 0, img)

		display.Display()
		display.WaitUntilIdle()

		time.Sleep(500 * time.Millisecond)
		current += 1
		if current == len(animation) {
			current = 0
		}
	}
}
