package main

import (
	"time"
	"image/color"
	"machine"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinydraw"
)

func main() {
	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	btnA := machine.BUTTON_A
	btnB := machine.BUTTON_B
	btnC := machine.BUTTON_C
	btnUp := machine.BUTTON_UP
	btnDown := machine.BUTTON_DOWN
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

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

	// Draw circles to represent each of the buttons
	tinydraw.FilledCircle(&display, 40, 110, 12, color.RGBA{1, 1, 1, 255})  // A
	tinydraw.FilledCircle(&display, 145, 110, 12, color.RGBA{1, 1, 1, 255}) // B
	tinydraw.FilledCircle(&display, 250, 110, 12, color.RGBA{1, 1, 1, 255}) // C
	tinydraw.FilledCircle(&display, 270, 30, 12, color.RGBA{1, 1, 1, 255})  // UP
	tinydraw.FilledCircle(&display, 270, 70, 12, color.RGBA{1, 1, 1, 255})  // DOWN

	// Show the buffer on the screen
	display.Display()

	// aux variables to save the state since the e-ink display takes too much time to refresh
	pressed := make([]bool, 5)
	oldPressed := make([]bool, 5)
	refresh := false

	for {
		refresh = false
		if btnA.Get() {
			pressed[0] = true
			tinydraw.Circle(&display, 40, 110, 14, color.RGBA{1, 1, 1, 255})
		} else {
			pressed[0] = false
			tinydraw.Circle(&display, 40, 110, 14, color.RGBA{0, 0, 0, 255})
		}
		if btnB.Get() {
			pressed[1] = true
			tinydraw.Circle(&display, 145, 110, 14, color.RGBA{1, 1, 1, 255})
		} else {
			pressed[1] = false
			tinydraw.Circle(&display, 145, 110, 14, color.RGBA{0, 0, 0, 255})
		}
		if btnC.Get() {
			pressed[2] = true
			tinydraw.Circle(&display, 250, 110, 14, color.RGBA{1, 1, 1, 255})
		} else {
			pressed[2] = false
			tinydraw.Circle(&display, 250, 110, 14, color.RGBA{0, 0, 0, 255})
		}
		if btnUp.Get() {
			pressed[3] = true
			tinydraw.Circle(&display, 270, 30, 14, color.RGBA{1, 1, 1, 255})
		} else {
			pressed[3] = false
			tinydraw.Circle(&display, 270, 30, 14, color.RGBA{0, 0, 0, 255})
		}
		if btnDown.Get() {
			pressed[4] = true
			tinydraw.Circle(&display, 270, 70, 14, color.RGBA{1, 1, 1, 255})
		} else {
			pressed[4] = false
			tinydraw.Circle(&display, 270, 70, 14, color.RGBA{0, 0, 0, 255})
		}

		// Check if any button has changed state
		// also, save the current state to the old state var
		for i:=0;i<5;i++ {
			if oldPressed[i]!=pressed[i] {
				refresh = true
				oldPressed[i] = pressed[i]
			}
		}

		// Refresh only if needed
		if refresh {
			display.Display()
		}
		time.Sleep(100*time.Millisecond)
	}
}
