package main

import (
	"image/color"
	"machine"
	"machine/usb"
	"machine/usb/adc/midi"
	"time"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const MIDICHANNEL = 10

var display uc8151.Device
var btnA, btnB, btnC, btnUp, btnDown machine.Pin

var black = color.RGBA{1, 1, 1, 255}
var white = color.RGBA{0, 0, 0, 255}

const WIDTH = 296
const HEIGHT = 128

func main() {
	usb.Product = "Badger Drum"

	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

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

	display.ClearBuffer()
	display.Display()

	str := "Go to https://virtualpiano.eu/"
	w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, str)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 30, str, black)
	str = "Select \"Badger Drum\""
	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, str)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 60, str, black)
	str = "as external MIDI keyboard"
	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, str)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 90, str, black)
	display.Display()

	notes := []int{36, 50, 55, 58, 39}
	pressed := make([]bool, 5)

	for {
		if btnA.Get() {
			if !pressed[0] {
				midi.Midi.NoteOn(0, MIDICHANNEL, midi.Note(notes[0]), 64)
				pressed[0] = true
			}
		} else {
			pressed[0] = false
		}
		if btnB.Get() {
			if !pressed[1] {
				midi.Midi.NoteOn(0, MIDICHANNEL, midi.Note(notes[1]), 64)
				pressed[1] = true
			}
		} else {
			pressed[1] = false
		}
		if btnC.Get() {
			if !pressed[2] {
				midi.Midi.NoteOn(0, MIDICHANNEL, midi.Note(notes[2]), 64)
				pressed[2] = true
			}
		} else {
			pressed[2] = false
		}
		if btnUp.Get() {
			if !pressed[3] {
				midi.Midi.NoteOn(0, MIDICHANNEL, midi.Note(notes[3]), 64)
				pressed[3] = true
			}
		} else {
			pressed[3] = false
		}
		if btnDown.Get() {
			if !pressed[4] {
				midi.Midi.NoteOn(0, MIDICHANNEL, midi.Note(notes[4]), 64)
				pressed[4] = true
			}
		} else {
			pressed[4] = false
		}
		time.Sleep(100 * time.Millisecond)
	}
}
