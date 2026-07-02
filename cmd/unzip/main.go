package main

import (
	"context"
	"flag"

	"github.com/mtkgo/common"
)

func main() {
	var inputFile string
	var outputDir string
	flag.StringVar(&inputFile, "input", "", "")
	flag.StringVar(&outputDir, "output", "", "")
	flag.Parse()

	err := common.UnZip(context.Background(), inputFile, outputDir)
	if err != nil {
		panic(err)
	}
}
