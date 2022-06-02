package decoder

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"
)

func TestScanFileHeader(t *testing.T) {

	fmt.Println(os.Getwd())
	byt, _ := ioutil.ReadFile("../../data/output/def")

	d := Decoder_createDecoder(byt, 7)
	// assert.Equal(t, (d.Width), 134)
	// assert.Equal(t, (d.Height), 201)
	d.Decoder_decode()
	file2, err2 := os.OpenFile("../../data/results/dec", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		os.Exit(1)
	}
	defer file2.Close()

	opt := jpeg.Options{
		Quality: 100,
	}

	jpeg.Encode(file2, d.Decoder_getImage(), &opt)
}
