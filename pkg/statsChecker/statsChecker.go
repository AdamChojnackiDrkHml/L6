package statschecker

import (
	"fmt"
	"image/jpeg"
	"math"
	"os"

	"github.com/ftrvxmtrx/tga"
)

type statschecker struct {
	originBitmap  [][][3]uint8
	decodedBitmap [][][3]uint8
}

func CheckStat(originPath, decodedPath string) {
	file, err := os.Open(originPath)
	if err != nil {
		os.Exit(1)
	}

	bitmap, err2 := tga.Decode(file)

	if err2 != nil {
		os.Exit(1)
	}

	originBitmap := make([][][3]uint8, bitmap.Bounds().Max.Y)
	for i := 0; i < bitmap.Bounds().Max.Y; i++ {
		originBitmap[i] = make([][3]uint8, bitmap.Bounds().Max.X)

		for j := 0; j < bitmap.Bounds().Max.X; j++ {
			r, g, b, _ := bitmap.At(j, i).RGBA()
			originBitmap[i][j] = [3]uint8{uint8(r / 256), uint8(g / 256), uint8(b / 256)}
		}
	}

	file, err = os.Open(decodedPath)
	if err != nil {
		os.Exit(1)
	}

	bitmap, err2 = jpeg.Decode(file)

	if err2 != nil {
		os.Exit(1)
	}

	decodedBitmap := make([][][3]uint8, bitmap.Bounds().Max.Y)
	for i := 0; i < bitmap.Bounds().Max.Y; i++ {
		decodedBitmap[i] = make([][3]uint8, bitmap.Bounds().Max.X)

		for j := 0; j < bitmap.Bounds().Max.X; j++ {
			r, g, b, _ := bitmap.At(j, i).RGBA()
			fmt.Println(r/256, g/256, b/256)

			decodedBitmap[i][j] = [3]uint8{uint8(r / 256), uint8(g / 256), uint8(b / 256)}
		}
	}

	mse := Mse(originBitmap, decodedBitmap)
	mseR := MseSingleCol(originBitmap, decodedBitmap, 0)
	mseG := MseSingleCol(originBitmap, decodedBitmap, 1)
	mseB := MseSingleCol(originBitmap, decodedBitmap, 2)

	snr := Snr(mse, originBitmap)

	fmt.Println("MSE := ", mse)
	fmt.Println("MSER := ", mseR)
	fmt.Println("MSEG := ", mseG)
	fmt.Println("MSEB := ", mseB)
	fmt.Println("SNR := ", snr)
}

func MseSingleCol(originBitmap, decodedBitmap [][][3]uint8, colorBit uint8) float64 {
	sum := 0.0
	for i := 0; i < len(originBitmap); i++ {
		for j := 0; j < len(originBitmap[0]); j++ {
			sum += math.Abs(float64(originBitmap[i][j][colorBit] - decodedBitmap[i][j][colorBit]))
		}
	}

	return sum
}

func Mse(originBitmap, decodedBitmap [][][3]uint8) float64 {
	sum := 0.0
	for i := 0; i < len(originBitmap); i++ {
		for j := 0; j < len(originBitmap[0]); j++ {
			sum += taxicab(originBitmap[i][j], decodedBitmap[i][j])
		}
	}
	return sum / float64(len(originBitmap)*len(originBitmap[0]))
}

func Snr(MSE float64, originBitmap [][][3]uint8) float64 {
	sum := 0.0
	for _, row := range originBitmap {
		for _, pix := range row {
			sum += math.Pow(float64(pix[0]), 2) + math.Pow(float64(pix[1]), 2) + math.Pow(float64(pix[2]), 2)

		}

	}
	return (sum * (1.0 / float64(len(originBitmap)*len(originBitmap[0])))) / MSE
}

func taxicab(vec1, vec2 [3]uint8) float64 {
	sum := 0.0
	vec1F := []float64{float64(vec1[0]), float64(vec1[1]), float64(vec1[2])}
	vec2F := []float64{float64(vec2[0]), float64(vec2[1]), float64(vec2[2])}

	for i := range vec1F {
		sum += math.Abs(vec1F[i] - vec2F[i])
	}

	return sum
}
