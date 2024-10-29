package main

import (
	"machine"
	"machine/usb"
	"machine/usb/adc/midi"

	"time"
)

var (
	note = midi.C3

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	pressed = false
)

func main() {
	usb.Product = "TinyGo One Note"

	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	pressed := false

	for {
		if btnA.Get() {
			if !pressed {
				midi.Port().NoteOn(midicable, midichannel, note, velocity)
				pressed = true
			}
		} else {
			midi.Port().NoteOff(midicable, midichannel, note, velocity)
			pressed = false
		}

		time.Sleep(time.Millisecond * 100)
	}
}
