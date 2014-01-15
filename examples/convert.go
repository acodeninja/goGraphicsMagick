package main

import (
	"flag"
	"fmt"
	gm "github.com/lgoldstien/goGraphicsMagick"
)

var _srcImagePath string
var _destImagePath string

func init() {
	flag.StringVar(&_srcImagePath, "src", "", "The source image you wish to convert from.")
	flag.StringVar(&_destImagePath, "dest", "", "The destination image you wish to output.")
}

func main() {
	flag.Parse()

	err := gm.Rotate(_srcImagePath, _destImagePath)
	if err != nil {
		fmt.Println(err)
	}
}
