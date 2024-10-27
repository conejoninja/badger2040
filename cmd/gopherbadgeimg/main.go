package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
	"strings"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
)

// flags for determining what to do
var (
	disableDithering bool
	outMode          string
	show             bool
	ratio            string
)

func main() {
	flag.BoolVar(&disableDithering, "disable-dithering", false, "disables dithering")
	flag.BoolVar(&show, "show", false, "paints dot-matrix-style art to the screen representing the image")
	flag.StringVar(
		&outMode,
		"outmode",
		"",
		"set the resize mode to one of: rice, bin, base64, or none",
	)
	flag.StringVar(
		&ratio,
		"ratio",
		"",
		"set the aspect ratio to predefined values including 'profile' or splash', or a custom value specified in the format of <height>x<width>.",
	)
	flag.Parse()
	if flag.NArg() != 1 {
		log.Printf("args: %v\n\n", flag.Args())
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s <input_image>:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"\nExamples:\n%s -outmode bin -ratio splash tainigo_128.png\n%s -outmode rice -ratio 128x128 -disable-dithering -show image.jpg\n",
			os.Args[0],
			os.Args[0],
		)
		os.Exit(1)
		return
	}
	infile := flag.Args()[0]
	if _, err := os.Stat(infile); err != nil {
		log.Fatalf("could not stat %v: %v", infile, err)
	}
	sourceImage, err := LoadImg(infile)
	if err != nil {
		log.Fatalf("error loading source image: %v", err)
	}
	var imgBits []byte

	var x, y int
	switch ratio {
	case "profile":
		// profile image is 128x128
		x, y = 128, 128
	case "splash":
		// splash image is 246x128
		x, y = 246, 128
	case "":
		log.Println("error: a ratio must be provided.")
		Usage()
		return
	default:
		x, y, err = ParseRatio(ratio)
		if err != nil {
			log.Println(err.Error())
			Usage()
			// The Usage function calls os.Exit(1) but LSPs and static analyzers often don't
			// pick up on that, so it's good practice to return from the caller anyway
			// For the sake of consistency, we return after toplevel os.Exit calls as well
			return
		}
	}
	imgBits = ImgToBytes(x, y, sourceImage)
	switch outMode {
	case "rice":
		err = WriteToGoFile(fmt.Sprintf("%s-generated.go", ratio), ratio, imgBits)
		if err != nil {
			log.Fatalf("error writing image to file: %v", err)
		}
	case "bin":
		err = WriteToBinFile(fmt.Sprintf("%s.bin", ratio), imgBits)
		if err != nil {
			log.Fatalf("error writing image to file: %v", err)
		}
	case "base64":
		fmt.Println(EncodeToString(imgBits))
	case "none":
		// this option is useful if you want to preview the file without creating it
	default:
		log.Printf("error: invalid outmode `%s`\n\n", outMode)
		Usage()
		return
	}
	if show {
		PrintImg(x, y, imgBits)
	}
}

// LoadImg loads and decodes filename into image.Image pointer
func LoadImg(infile string) (*image.Image, error) {
	f, err := os.Open(infile)
	if err != nil {
		return nil, err
	}
	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return &src, nil
}

func ParseRatio(rstr string) (int, int, error) {
	rstr = strings.ToLower(rstr)
	pixels := strings.Split(rstr, "x")
	if len(pixels) != 2 {
		return 0, 0, errors.New("invalid ratio string provided")
	}
	x, err := strconv.Atoi(pixels[0])
	if err != nil {
		return 0, 0, errors.Join(errors.New("error: could not parse x coordinate count"), err)
	}
	y, err := strconv.Atoi(pixels[1])
	if err != nil {
		return 0, 0, errors.Join(errors.New("error: could not parse y coordinate count"), err)
	}
	return x, y, nil
}

// Usage prints a proper example of usage for when the user misuses the program.
//
// Usage also calls exit(1) to terminate the program with an error code.
func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s <input_image>:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(
		flag.CommandLine.Output(),
		"\nExamples:\n%s input.png -outmode bin -ratio profile\n%s input.jpg -outmode rice -ratio 128x128 -disable-dithering -show\n",
		os.Args[0],
		os.Args[0],
	)
	os.Exit(1)
}

// PrintImg prints an `*` for each marked bit
//
// It writes to stderr so that it doesn't conflict with the base64 output
func PrintImg(x, y int, imageBits []byte) {
	// create a matrix of bools to represent the image pixels
	matrix := make([][]bool, y)
	for i := range matrix {
		matrix[i] = make([]bool, x)
	}
	// unpack the bits into the matrix
	count := 0
	for dy := 0; dy < y; dy++ {
		for dx := 0; dx < x; dx++ {
			matrix[dy][dx] = (imageBits[calcIndex(count)] & (1 << calcBit(count))) != 0
			count++
		}
	}

	// print the matrix to stderr
	for dy := 0; dy < y; dy++ {
		for dx := 0; dx < x; dx++ {
			if matrix[dy][dx] {
				fmt.Fprint(os.Stderr, "*")
			} else {
				fmt.Fprint(os.Stderr, " ")
			}
		}
		fmt.Fprintln(os.Stderr)
	}
}
