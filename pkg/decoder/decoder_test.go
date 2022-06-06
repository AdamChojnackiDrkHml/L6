package decoder

import (
	"fmt"
	"io/ioutil"
	statschecker "l6/pkg/statsChecker"
	"os"
	"testing"

	"github.com/ftrvxmtrx/tga"
)

func TestScanFileHeader(t *testing.T) {

	fmt.Println(os.Getwd())
	byt, _ := ioutil.ReadFile("../../data/output/def")

	d := Decoder_createDecoder(byt)
	// assert.Equal(t, (d.Width), 134)
	// assert.Equal(t, (d.Height), 201)
	d.Decoder_decode()
	file2, err2 := os.OpenFile("../../data/results/dec2", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		os.Exit(1)
	}
	defer file2.Close()
	file3, err3 := os.Open("../../data/input/testy4/example0.tga")

	if err3 != nil {
		os.Exit(1)
	}

	bitmap, err4 := tga.Decode(file3)
	if err4 != nil {
		os.Exit(1)
	}

	rgbBitMap := make([][][3]uint8, bitmap.Bounds().Max.Y)
	for i := 0; i < bitmap.Bounds().Max.Y; i++ {
		rgbBitMap[i] = make([][3]uint8, bitmap.Bounds().Max.X)

		for j := 0; j < bitmap.Bounds().Max.X; j++ {
			r, g, b, _ := bitmap.At(j, i).RGBA()
			rgbBitMap[i][j] = [3]uint8{uint8(r / 256), uint8(g / 256), uint8(b / 256)}
		}
	}

	statschecker.AAA(rgbBitMap, d.RgbBitMap)
	// opt := jpeg.Options{
	// 	Quality: 100,
	// }
	img := d.Decoder_getImage()
	tga.Encode(file2, img)
}
