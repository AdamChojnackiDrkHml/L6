package coder

import (
	"fmt"
	"image"
	"l6/pkg/quantizer"
	"strconv"
)

const (
	RED   = 0
	GREEN = 1
	BLUE  = 2
)

type Coder struct {
	rgbBitMap  [][][3]uint8
	filterLow  [][][3]uint8
	filterHigh [][][3]uint8
	width      uint32
	height     uint32
	q          *quantizer.Quantizer
	Outfile    []byte
}

func Coder_createCoder(bitmap image.Image, colors uint8) *Coder {
	coder := &Coder{}

	//PLACEHOLDER
	fmt.Println(bitmap.Bounds().Max.Y + 1)
	fmt.Println(bitmap.Bounds().Max.X + 1)
	fmt.Println(len(coder.rgbBitMap))
	coder.rgbBitMap = make([][][3]uint8, bitmap.Bounds().Max.Y)
	coder.filterLow = make([][][3]uint8, bitmap.Bounds().Max.Y)
	coder.filterHigh = make([][][3]uint8, bitmap.Bounds().Max.Y)
	coder.Outfile = make([]byte, 0)
	for i := 0; i < bitmap.Bounds().Max.Y; i++ {
		coder.rgbBitMap[i] = make([][3]uint8, bitmap.Bounds().Max.X)
		coder.filterLow[i] = make([][3]uint8, bitmap.Bounds().Max.X)
		coder.filterHigh[i] = make([][3]uint8, bitmap.Bounds().Max.X)

		for j := 0; j < bitmap.Bounds().Max.X; j++ {
			r, g, b, _ := bitmap.At(j, i).RGBA()
			coder.rgbBitMap[i][j] = [3]uint8{uint8(r / 256), uint8(g / 256), uint8(b / 256)}
		}
	}
	coder.q = quantizer.Quantizer_createQuantizer(colors)
	coder.width = uint32(bitmap.Bounds().Max.X)
	coder.height = uint32(bitmap.Bounds().Max.Y)
	fmt.Println(coder.width * coder.height * 3)

	return coder
}

func (c *Coder) Coder_encode() {
	filLow, filHigh := c.filter()

	pixLowArray := c.differentialCoding(filLow)

	pixHighArray := c.quantify(filHigh)
	c.Outfile = append(c.Outfile, []byte(strconv.Itoa(int(c.width)))...)
	c.Outfile = append(c.Outfile, []byte(" ")...)
	c.Outfile = append(c.Outfile, []byte(strconv.Itoa(int(c.height)))...)
	c.Outfile = append(c.Outfile, []byte(" ")...)

	for i := range pixLowArray {
		c.Outfile = append(c.Outfile, byte(pixLowArray[i][RED]))
		c.Outfile = append(c.Outfile, byte(pixLowArray[i][GREEN]))
		c.Outfile = append(c.Outfile, byte(pixLowArray[i][BLUE]))

		c.Outfile = append(c.Outfile, byte(pixHighArray[i][RED]))
		c.Outfile = append(c.Outfile, byte(pixHighArray[i][GREEN]))
		c.Outfile = append(c.Outfile, byte(pixHighArray[i][BLUE]))

	}

}

func (c *Coder) filter() ([][3]uint8, [][3]uint8) {

	filLow := make([][3]uint8, 0)
	filHigh := make([][3]uint8, 0)
	prev := [3]uint8{0, 0, 0}
	for i := 0; i < int(c.height); i++ {
		for j := 0; j < int(c.width); j++ {

			r, g, b := c.rgbBitMap[i][j][RED], c.rgbBitMap[i][j][GREEN], c.rgbBitMap[i][j][BLUE]

			Yn := [3]uint8{r/2 + prev[RED]/2, g/2 + prev[GREEN]/2, b/2 + prev[BLUE]/2}
			Zn := [3]uint8{r - Yn[RED], g - Yn[GREEN], b - Yn[BLUE]}

			filLow = append(filLow, Yn)
			filHigh = append(filHigh, Zn)
			prev = [3]uint8{r, g, b}
		}
	}
	filLowTrunc := make([][3]uint8, 0)
	filHighTrunc := make([][3]uint8, 0)

	for i := 0; i < len(filLow); i += 2 {
		filLowTrunc = append(filLowTrunc, filLow[i])
		filHighTrunc = append(filHighTrunc, filHigh[i])
	}

	return filLowTrunc, filHighTrunc
}

func (c *Coder) differentialCoding(pixels [][3]uint8) [][3]uint8 {
	result := make([][3]uint8, len(pixels))
	result[0] = pixels[0]
	quantum := c.quantify(pixels)

	for i := 1; i < len(pixels); i++ {
		result[i][RED] = (pixels[i][RED] - quantum[i-1][RED])
		result[i][GREEN] = (pixels[i][GREEN] - quantum[i-1][GREEN])
		result[i][BLUE] = (pixels[i][BLUE] - quantum[i-1][BLUE])
	}
	result = c.quantify(result)

	return result
}

func (c *Coder) quantify(bytes [][3]uint8) [][3]uint8 {
	quantized := make([][3]uint8, len(bytes))

	for i := range quantized {
		quantized[i] = [3]uint8{0, 0, 0}
		quantized[i][RED] = c.q.Quantizer_getQuant(bytes[i][RED])
		quantized[i][GREEN] = c.q.Quantizer_getQuant(bytes[i][GREEN])
		quantized[i][BLUE] = c.q.Quantizer_getQuant(bytes[i][BLUE])

	}

	return quantized
}
