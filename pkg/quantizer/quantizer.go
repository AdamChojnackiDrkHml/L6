package quantizer

import (
	"fmt"
	"math"
)

//stationary

type Quantizer struct {
	intervals uint8
}

func Quantizer_createQuantizer(bits uint8) *Quantizer {
	return &Quantizer{intervals: pow2(bits)}
}

func (q *Quantizer) Quantizer_getQuant(toQuant uint8) uint8 {
	if toQuant > 255 {
		fmt.Println(toQuant, 255)
		panic("ERROR")
	}

	diff := uint8((256) / int64(q.intervals))

	return toQuant / diff * diff
}

func pow2(a uint8) uint8 {
	return uint8(math.Pow(float64(2), float64(a)))
}
