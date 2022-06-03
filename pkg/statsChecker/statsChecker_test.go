package statschecker

import (
	"fmt"
	"os"
	"testing"
)

func TestStatsCheck(t *testing.T) {

	fmt.Println(os.Getwd())

	CheckStat("../../data/input/testy4/example0.tga", "../../data/results/dec")
}
