package main

import (
	_ "embed"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

// This code allows you to re-convert a bitmap variable file back

// must build and run the program first, and generate a splash.bin file using
// ./gopherbadgeimg splash <image file here>

//go:embed splash.bin
var tainigo []byte

func TestMain(t *testing.T) {
	decodeToPng()
}

func decodeToPng() {
	dst := image.NewRGBA(image.Rect(0, 0, 246, 128))
	outPng, _ := os.Create("splash.png")
	defer outPng.Close()
	for j := 0; j < 246; j++ {
		for i := 0; i < 128; i++ {
			offset := i + j*128
			bit := tainigo[offset/8] & (1 << uint(7-offset%8))
			if bit != 0 {
				dst.Set(245-j, i, color.RGBA{255, 255, 255, 255})
			} else {
				dst.Set(245-j, i, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	png.Encode(outPng, dst)
}
