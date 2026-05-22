package main

import (
	"flag"
	"image"
	. "image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"strings"
)

var c64palette = Palette{
	RGBA{0x00, 0x00, 0x00, 0xff}, RGBA{0xff, 0xff, 0xff, 0xff},
	RGBA{0x91, 0x4a, 0x40, 0xff}, RGBA{0x84, 0xc5, 0xcb, 0xff},
	RGBA{0x93, 0x51, 0xb6, 0xff}, RGBA{0x72, 0xb1, 0x4a, 0xff},
	RGBA{0x48, 0x39, 0xaa, 0xff}, RGBA{0xd5, 0xdf, 0x7b, 0xff},
	RGBA{0x99, 0x69, 0x2d, 0xff}, RGBA{0x67, 0x52, 0x00, 0xff},
	RGBA{0xc1, 0x81, 0x78, 0xff}, RGBA{0x60, 0x60, 0x60, 0xff},
	RGBA{0x8a, 0x8a, 0x8a, 0xff}, RGBA{0xb3, 0xec, 0x91, 0xff},
	RGBA{0x86, 0x7a, 0xdd, 0xff}, RGBA{0xb3, 0xb3, 0xb3, 0xff},
}

var (
	outputFmt   string
	outputFile  string
	imageSize   int
	jpegQuality int
)

var imageEncoder func(io.Writer, image.Image) error

func init() {
	log.SetFlags(0)
	log.SetPrefix("avatar: ")
	log.SetOutput(os.Stderr)

	showHelp := flag.Bool("h", false, "show this help and exit")
	shuffle := flag.Bool("r", false, "randomize palette order")
	flag.IntVar(&imageSize, "s", 128, "side of the square in pixels")
	flag.StringVar(&outputFmt, "t", "png", "output format")
	flag.IntVar(&jpegQuality, "q", 90, "jpeg quality [1-100]")
	flag.StringVar(&outputFile, "o", "/dev/stdout", "output file")
	flag.Parse()

	if *showHelp {
		flag.CommandLine.SetOutput(os.Stdout)
		flag.Usage()
		os.Exit(0)
	}

	if imageSize < 4 {
		log.Fatalf("invalid size %d: it must be at least 4\n", imageSize)
	}

	switch strings.ToLower(outputFmt) {
	case "png":
		imageEncoder = encodePng
	case "gif":
		imageEncoder = encodeGif
	case "jpg":
		fallthrough
	case "jpeg":
		if jpegQuality < 1 || jpegQuality > 100 {
			log.Fatalf("invalid jpeg quality %d: it ranges from 1 to 100 inclusive\n", jpegQuality)
		}
		imageEncoder = encodeJpeg
	default:
		log.Fatalf("unknown output format `%s'\n", outputFmt)
	}

	if *shuffle {
		rand.Shuffle(len(c64palette), func(i, j int) {
			c64palette[i], c64palette[j] = c64palette[j], c64palette[i]
		})
	}
}

func main() {
	r := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewRGBA(r)

	cell := imageSize / 4
	for j := 0; j < imageSize; j++ {
		for i := 0; i < imageSize; i++ {
			row := j / cell
			col := i / cell
			k := row*4 + col
			img.Set(i, j, c64palette[k])
		}
	}

	var wr io.Writer
	var err error
	if outputFile == "/dev/stdout" {
		wr = os.Stdout
	} else {
		wr, err = os.Create(outputFile)
		if err != nil {
			log.Fatalln(err)
		}
	}
	err = imageEncoder(wr, img)
	if err != nil {
		log.Fatalln("could not write image:", err)
	}
}

func encodePng(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}

func encodeGif(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, &gif.Options{NumColors: 16})
}

func encodeJpeg(w io.Writer, i image.Image) error {
	return jpeg.Encode(w, i, &jpeg.Options{Quality: jpegQuality})
}
