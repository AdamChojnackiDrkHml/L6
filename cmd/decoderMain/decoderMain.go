package main

import (
	"fmt"
	"io/ioutil"
	"l6/pkg/decoder"
	"os"

	"github.com/ftrvxmtrx/tga"
)

func main() {
	fmt.Println(os.Getwd())
	var inPath string
	var outPath string
	if len(os.Args) < 3 {
		inPath = "../../data/output/def"
		outPath = "../../data/results/decoded"
	} else {
		inPath = os.Args[1]
		outPath = os.Args[2]

	}
	byt, _ := ioutil.ReadFile(inPath)

	d := decoder.Decoder_createDecoder(byt)
	// assert.Equal(t, (d.Width), 134)
	// assert.Equal(t, (d.Height), 201)
	d.Decoder_decode()
	file2, err2 := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err2 != nil {
		os.Exit(1)
	}

	img := d.Decoder_getImage()
	tga.Encode(file2, img)

}
