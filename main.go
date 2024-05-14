package main

import (
	"machine"
	"tinygo.org/x/drivers/pixel"

	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
)

var display board.Displayer[pixel.Monochrome]
var screen *tinygl.Screen[pixel.Monochrome]
var led machine.Pin
var (
	white = pixel.NewColor[pixel.Monochrome](0xff, 0xff, 0xff)
	black = pixel.NewColor[pixel.Monochrome](0x00, 0x00, 0x00)
)

func main() {
	led = machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	display = board.Display.Configure()
	width, height := display.Size()

	buf := pixel.NewImage[pixel.Monochrome](int(width), int(height)/4)
	screen = tinygl.NewScreen[pixel.Monochrome](display, buf, board.Display.PPI())

	//canvas := gfx.NewCanvas(black, int(width), int(height))
	//screen.SetChild(canvas)

	screen.Update()
}
