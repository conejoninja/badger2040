package main

import (
	"machine"
	"machine/usb"
	"machine/usb/adc/midi"

	"time"

	"tinygo.org/x/drivers/pcf8591"
)

var (
	note = midi.C3

	midicable   uint8 = 0
	midichannel uint8 = 1
	velocity    uint8 = 0x40

	pressed = false
)

var i2c = machine.I2C0

func main() {
	usb.Product = "TinyGo Therestick"

	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	pressed := false

	i2c.Configure(machine.I2CConfig{})
	adc := pcf8591.New(i2c)
	adc.Configure()

	stickX, stickY := adc.CH0, adc.CH1

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

		// x axis for modulation, y axis for pitch bend
		x, y := stickX.Get(), stickY.Get()
		println(x, y)

		if pressed {
			// scale x to range 0x0 thru 0xff
			sx := 0xFF * uint32(x) / 0xFFFF
			midi.Port().ControlChange(midicable, midichannel, midi.CCModulationWheel, uint8(sx))

			// scale y to range 0x0 thru 0x3FFF
			sy := 0x3FFF * uint32(y) / 0xFFFF
			midi.Port().PitchBend(midicable, midichannel, uint16(sy))
		}

		time.Sleep(time.Millisecond * 100)
	}
}
