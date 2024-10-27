package main

import (
	"image"
	"image/color"
	"log"

	"github.com/makeworld-the-better-one/dither"
	"golang.org/x/image/draw"
)

// ImgToBytes resizes an image to the requested size and converts it to a bitmap byte slice
func ImgToBytes(x, y int, inputImg *image.Image) []byte {
	// work on values not pointers
	src := *inputImg
	// create a new, rectangular image that's the size we want
	dst := image.NewRGBA(image.Rect(0, 0, x, y))
	// use NearestNeighbor algo to fit our original image into the smaller (or bigger!?) image
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	// Our e-ink display uses one bit for each pixel, on or off.
	// Therefore, we need one bit for each pixel.
	// Since we have a byte slice, and 8 bytes per bit, divide by 8
	imageBits := make([]byte, x*y/8)

	// Again, on or off, white or black are our only color options
	palette := []color.Color{
		color.Black,
		color.White,
	}

	if disableDithering {
		// don't dither image if flag is set, useful for some images which are already black and white
	} else {
		// using our palette, create a dithering struct
		// and dither our image to get some false shading.
		// read more here: https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering
		d := dither.NewDitherer(palette)
		d.Matrix = dither.FloydSteinberg
		dithered := d.Dither(dst)
		// this nil check is necessary since the library will often write
		// the dithered image to dst, but not always. Read their docs for more info
		if dithered != nil {
			var ok bool
			dst, ok = dithered.(*image.RGBA)
			// docs claim image is guaranteed to be of this type when not nil, but it's good to check anyway
			if !ok {
				log.Fatalf("error: typeof dithered should have been `*image.RGBA` but was `%T`", dithered)
			}
		}
	}

	count := 0
	for dy := 0; dy < y; dy++ {
		for dx := 0; dx < x; dx++ {
			// grab image point, determine if bit should be 1 or a 0
			r, g, b, _ := dst.At(dx, dy).RGBA()
			if r+g+b == 0 {
				// use bit shifting + integer division & modulo arithmetic to change
				// the individual bits we want to set
				//			imageBits[(i*y+j)/8] = imageBits[(i*y+j)/8] | (1 << uint(7-(i*y+j)%8))
				imageBits[calcIndex(count)] |= (1 << calcBit(count))
			}
			count++
		}
	}
	return imageBits
}

func calcIndex(c int) int {
	return c / 8
}

func calcBit(c int) uint {
	return uint(7 - c%8)
}
