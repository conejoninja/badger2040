package main

import (
	"machine"
	"machine/usb"
	"machine/usb/adc/midi"

	"time"
)

type chord struct {
	name  string
	notes []midi.Note
}

var (
	buttonC  = machine.BUTTON_A
	buttonG  = machine.BUTTON_B
	buttonAm = machine.BUTTON_C
	buttonF  = machine.BUTTON_DOWN

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	chords = []struct {
		name  string
		notes []midi.Note
	}{
		{name: "C ", notes: []midi.Note{midi.C3, midi.E3, midi.G3}},
		{name: "G ", notes: []midi.Note{midi.G3, midi.B3, midi.D4}},
		{name: "Am", notes: []midi.Note{midi.A3, midi.C4, midi.E4}},
		{name: "F ", notes: []midi.Note{midi.F3, midi.A3, midi.C4}},
	}

	pressed bool
)

func main() {
	usb.Product = "TinyGo Chorder"

	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	pressed := false
	index := 0

	for {
		if btnA.Get() {
			if !pressed {
				for _, note := range chords[index].notes {
					midi.Port().NoteOn(midicable, midichannel, note, velocity)
				}
				pressed = true
			}
		} else {
			if pressed {
				for _, note := range chords[index].notes {
					midi.Port().NoteOff(midicable, midichannel, note, velocity)
				}
				index = (index + 1) % len(chords)
				pressed = false
			}
		}

		time.Sleep(time.Millisecond * 100)
	}
}
