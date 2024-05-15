package main

import (
	"time"

	"tinygo.org/x/drivers/pixel"

	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/gfx"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/mono"
)

func main() {
	run(board.Display.Configure())
}

func run[T pixel.Color](display board.Displayer[T]) {
	board.Display.SetBrightness(board.Display.MaxBrightness())
	time.Sleep(1 * time.Second)

	width, height := display.Size()
	scalePercent := board.Display.PPI() * 100 / 120

	// Initialize the screen.
	buf := pixel.NewImage[T](int(width), int(height)/4)
	screen := tinygl.NewScreen[T](display, buf, board.Display.PPI())

	var (
		white = pixel.NewColor[T](0xFF, 0xFF, 0xFF)
		black = pixel.NewColor[T](0x00, 0x00, 0x00)
	)
	theme := mono.New(style.NewScale(scalePercent), screen)
	helloText := theme.NewText("Hi")
	canvas := gfx.NewCanvas(black, 96, 96)

	vBox := theme.NewVBox(helloText, canvas)

	screen.SetChild(vBox)
	screen.Update()

	brickW := gfx.NewRect(white, 0, 0, 288, 10)
	canvas.Add(brickW)

	for {
		//brickW.Move(0, 0)
		helloText.SetText("Goodbye")
		screen.Update()
		display.Display()
		time.Sleep(500 * time.Millisecond)
		helloText.SetText("Hello")
		screen.Update()
		display.Display()
		time.Sleep(500 * time.Millisecond)
	}
}
