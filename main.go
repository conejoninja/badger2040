package main

import (
	"time"

	"tinygo.org/x/drivers/pixel"

	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/mono"
)

func main() {
	run(board.Display.Configure())
}

func run[T pixel.Color](display board.Displayer[T]) {
	time.Sleep(1 * time.Second)

	width, height := display.Size()
	scalePercent := board.Display.PPI() * 100 / 120

	// Initialize the screen.
	buf := pixel.NewImage[T](int(width), int(height)/4)
	screen := tinygl.NewScreen[T](display, buf, board.Display.PPI())

	theme := mono.New(style.NewScale(scalePercent), screen)
	header := theme.NewText("Hello, TinyGL")
	helloText := theme.NewText("Hi")
	vBox := theme.NewVBox(header, helloText)

	screen.SetChild(vBox)
	screen.Update()
}
