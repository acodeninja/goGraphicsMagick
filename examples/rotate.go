package main

import (
	"flag"
	"fmt"
	"C"
	gm "github.com/lgoldstien/goGraphicsMagick"
)

var _srcImagePath string
var _destImagePath string
var _degrees string

func init() {
	flag.StringVar(&_srcImagePath, "src", "", "The source image you wish to convert from.")
	flag.StringVar(&_destImagePath, "dest", "", "The destination image you wish to output.")
	flag.StringVar(&_degrees, "deg", "0", "The desired degrees of rotation.")
}

func main() {
	flag.Parse()

	err := gm.Rotate(_srcImagePath, _destImagePath, _degrees)
	if err != nil {
		fmt.Println(err)
	}
}
