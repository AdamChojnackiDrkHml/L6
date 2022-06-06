package main

import (
	"fmt"
	statschecker "l6/pkg/statsChecker"
	"os"
)

func main() {
	fmt.Println(os.Getwd())

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
	statschecker.CheckStat(inPath, outPath)

}
