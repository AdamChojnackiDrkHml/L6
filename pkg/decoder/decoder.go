package decoder

import (
	"fmt"
	"image"
	"image/color"
	"l6/pkg/quantizer"
	"strconv"
)

const (
	RED   = 0
	GREEN = 1
	BLUE  = 2
)

type Decoder struct {
	RgbBitMap  [][][3]uint8
	filterLow  [][3]uint8
	filterHigh [][3]uint8
	Width      int
	Height     int
	q          *quantizer.Quantizer
	Infile     []byte
}

func Decoder_createDecoder(bytes []byte) *Decoder {
	decoder := &Decoder{}

	//PLACEHOLDER
	numString := ""
	for {
		if bytes[0] == byte(' ') {
			bytes = bytes[1:]
			break
		}
		numString += string(bytes[0])
		bytes = bytes[1:]
	}
	decoder.Width, _ = strconv.Atoi(numString)
	numString = ""
	for {
		if bytes[0] == byte(' ') {
			bytes = bytes[1:]

			break
		}
		numString += string(bytes[0])
		bytes = bytes[1:]
	}
	decoder.Height, _ = strconv.Atoi(numString)
	decoder.Infile = bytes

	pixHighArray := make([][3]uint8, 0)
	pixLowArray := make([][3]uint8, 0)

	for i := 0; i < len(bytes); i += 6 {
		pixLowArray = append(pixLowArray, [3]uint8{uint8(bytes[i]), uint8(bytes[i+1]), uint8(bytes[i+2])})
		pixHighArray = append(pixHighArray, [3]uint8{uint8(bytes[i+3]), uint8(bytes[i+4]), uint8(bytes[i+5])})
	}

	decoder.filterLow = pixLowArray
	decoder.filterHigh = pixHighArray
	// decoder.q = quantizer.Quantizer_createQuantizer(colors)

	return decoder
}

func (d *Decoder) Decoder_decode() {
	d.diffDecode()
	d.lowHighDecoding()
}

func (d *Decoder) diffDecode() {
	pixels := make([][3]uint8, len(d.filterLow))
	pixels[0] = d.filterLow[0]
	for i := 1; i < len(pixels); i++ {
		pixels[i][RED] = pixels[i-1][RED] + d.filterLow[i][RED]
		pixels[i][GREEN] = pixels[i-1][GREEN] + d.filterLow[i][GREEN]
		pixels[i][BLUE] = pixels[i-1][BLUE] + d.filterLow[i][BLUE]
	}

	d.filterLow = pixels
}

func (d *Decoder) lowHighDecoding() {

	decoded := make([][3]uint8, d.Width*d.Height)
	decoded[0][RED] = d.filterLow[0][RED] + d.filterHigh[0][RED]
	decoded[0][GREEN] = d.filterLow[0][GREEN] + d.filterHigh[0][GREEN]
	decoded[0][BLUE] = d.filterLow[0][BLUE] + d.filterHigh[0][BLUE]

	for i := 1; i < len(d.filterLow); i++ {
		decoded[2*i][RED] = d.filterLow[i][RED] + d.filterHigh[i][RED]
		decoded[2*i][GREEN] = d.filterLow[i][GREEN] + d.filterHigh[i][GREEN]
		decoded[2*i][BLUE] = d.filterLow[i][BLUE] + d.filterHigh[i][BLUE]

		decoded[(2*i)-1][RED] = d.filterLow[(i)][RED] - d.filterHigh[(i)][RED]
		decoded[(2*i)-1][GREEN] = d.filterLow[(i)][GREEN] - d.filterHigh[(i)][GREEN]
		decoded[(2*i)-1][BLUE] = d.filterLow[(i)][BLUE] - d.filterHigh[(i)][BLUE]
	}
	fmt.Println(len(decoded))
	d.RgbBitMap = make([][][3]uint8, d.Height)
	for i := 0; i < d.Height; i++ {
		d.RgbBitMap[i] = make([][3]uint8, d.Width)
		for j := 0; j < d.Width; j++ {
			index := d.Width*i + j
			r, g, b := decoded[index][RED], decoded[index][GREEN], decoded[index][BLUE]
			d.RgbBitMap[i][j] = [3]uint8{r, g, b}
		}
	}
}

func (d *Decoder) Decoder_getImage() image.Image {

	width := int(d.Width)
	height := int(d.Height)
	fmt.Println(width, height)
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewNRGBA(image.Rectangle{upLeft, lowRight})

	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {

			pixcols := d.RgbBitMap[i][j]
			col := color.NRGBA{uint8(pixcols[RED]), uint8(pixcols[GREEN]), uint8(pixcols[BLUE]), 255.0}
			img.Set(j, i, col)
		}

	}

	return img
}
