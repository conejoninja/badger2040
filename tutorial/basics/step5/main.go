package main

import (
	"machine"
	"machine/usb/hid/mouse"
	"time"
)

var btnA, btnB, btnC, btnUp, btnDown machine.Pin

const STEP = 8

func main() {

	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	btnA = machine.BUTTON_A
	btnB = machine.BUTTON_B
	btnC = machine.BUTTON_C
	btnUp = machine.BUTTON_UP
	btnDown = machine.BUTTON_DOWN
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	mouseDevice := mouse.Port()

	var x, y int
	for {
		x = 0
		y = 0

		if btnA.Get() {
			x = -STEP
		}
		if btnC.Get() {
			x = STEP
		}
		if btnUp.Get() {
			y = -STEP
		}
		if btnDown.Get() {
			y = STEP
		}

		mouseDevice.Move(x, y)

		if btnB.Get() {
			mouseDevice.Click(mouse.Left)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
