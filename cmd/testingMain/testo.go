package main

import (
	"fmt"
	"os"

	"github.com/ftrvxmtrx/tga"
)

func main() {
	fmt.Println(os.Getwd())
	var inPath string

	if len(os.Args) < 4 {
		inPath = "../../data/input/testy4/example0.tga"
	} else {
		inPath = os.Args[1]

	}

	file, err := os.Open(inPath)

	if err != nil {
		os.Exit(1)
	}

	img, err2 := tga.Decode(file)

	if err2 != nil {
		os.Exit(1)
	}

	file2, err3 := os.OpenFile("../../data/results/dec2", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err3 != nil {
		os.Exit(1)
	}

	err4 := tga.Encode(file2, img)

	if err4 != nil {
		os.Exit(1)
	}
}
