// Copyright 2014 Lawrence Goldstien lgoldstien@onmylemon.co.uk
// Use of this source code is govered by an MIT license
// Licensing information can be found in the LICENSE file
package goGraphicsMagick

// #cgo CFLAGS: -IGM_SRC
// #cgo pkg-config: GraphicsMagickWand
// #include "_c_utils.c"
import "C"

import (
	"errors"
  "fmt"
)

// Initialise some of the standard variables we will need
var goArgv []string
var goArgc int
var cArgv **C.char
var cArgc C.int


func checkPaths(srcPath, destPath string) error {
  var err error = nil
  if srcPath == "" {
    err = errors.New("Failure, there is no source image. Please specify one using -src <src.jpg>.")
  }

  if destPath == "" {
    err = errors.New("Failure, there is no destination image. Please specify one using -dest <dest.png>.")
  }
  return err
}

// C helper function to convert a slice of strings to a **char for argv
func goSliceStringToCharStar(goSliceString []string) (cCharStar **C.char) {
	cCharStar = C.makeCharArray(C.int(len(goSliceString)))

	for i, s := range goSliceString {
		C.setArrayString(cCharStar, C.CString(s), C.int(i))
	}

	return cCharStar
}

// The rotate function
// Rotates an image and puts on an optional background
func Rotate(srcPath, destPath, degrees string) error {
	var err error = nil

  // Check the source paths exist
  err = checkPaths(srcPath, destPath)
  if err != nil {
    panic("The src or destination paths have not been set.")
  }

	// Append the source and destination to argv
	goArgv = append(goArgv, "")
	goArgv = append(goArgv, srcPath)
  goArgv = append(goArgv, destPath)
  goArgv = append(goArgv, degrees)

	cArgv = goSliceStringToCharStar(goArgv)
	// defer C.freeCharArray(cArgv, C.int(len(goArgv)))
	cArgc = C.int(len(goArgv))

	C.ImageRotate(cArgc, cArgv)

	return err
}

// The conversion function
// Converts an image from one format to another
func Convert(srcPath, destPath string) error {
  var err error = nil

  if srcPath == "" {
    err = errors.New("Failure, there is no source image. Please specify one using -src <src.jpg>.")
  }

  if destPath == "" {
    err = errors.New("Failure, there is no destination image. Please specify one using -dest <dest.png>.")
  }

  // Append the source and destination to argv
  goArgv = append(goArgv, "")
  goArgv = append(goArgv, srcPath)
  goArgv = append(goArgv, destPath)

  fmt.Println(goArgv)

  cArgv = goSliceStringToCharStar(goArgv)
  // defer C.freeCharArray(cArgv, C.int(len(goArgv)))
  cArgc = C.int(len(goArgv))

  C.ImageConvert(cArgc, cArgv)

  return err
}