package main

import (
	"flag"
	"image/color"
	"image/png"
	"log"
	"os"

	"encoding/base64"

	"fmt"

	dither "github.com/makeworld-the-better-one/dither/v2"
)

func main() {
	filepath := flag.String("filepath", "", "Fullpath of the image")
	noDither := flag.Bool("nodither", false, "To avoid dithering the image")
	showImage := flag.Bool("showimage", false, "Display the image")
	flag.Parse()

	if *filepath == "" {
		fmt.Println("-filepath can not be empty if -conf=custom")
		return
	}

	file, _ := os.Open(*filepath)
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal("failed to decode image file", err)
	}

	if !*noDither {
		palette := []color.Color{
			color.Black,
			color.White,
		}

		// Create ditherer
		d := dither.NewDitherer(palette)
		d.Matrix = dither.FloydSteinberg

		// Dither the image, attempting to modify the existing image
		// If it can't then a dithered copy will be returned.
		img = d.Dither(img)

		f2, err := os.Create("dither.png")
		if err != nil {
			panic(err)
		}

		err = png.Encode(f2, img)
		if err != nil {
			panic(err)
		}
	}

	buffer := make([]uint8, int(img.Bounds().Max.Y)*int(img.Bounds().Max.X)/8)

	bbb := 0
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < (img.Bounds().Max.Y / 8); y++ {
			var buf uint8
			for i := uint8(0); i < 8; i++ {
				r, g, b, _ := img.At(img.Bounds().Max.X-x-1, (y*8)+int(i)).RGBA()
				if r == 0 && g == 0 && b == 0 {
					buf |= 1 << (7 - i)
					if *showImage {
						fmt.Print("*")
					}
				} else {
					if *showImage {
						fmt.Print(" ")
					}
				}
			}
			buffer[bbb] = buf
			bbb++
		}
		if *showImage {
			fmt.Println("")
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Use this if you plan to use an array directly in the code")
	for b := 0; b < len(buffer); b++ {
		fmt.Printf("0x%X,", buffer[b])
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Use this for the ldflags main.ProfilePic")
	fmt.Println(base64.StdEncoding.EncodeToString(buffer))
}
