package main

import (
	"fmt"
	statschecker "l6/pkg/statsChecker"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	statschecker.CheckStat("data/input/testy4/example0.tga", "data/results/dec2")

}
