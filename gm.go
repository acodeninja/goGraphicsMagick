// Copyright 2014 Lawrence Goldstien lgoldstien@onmylemon.co.uk
// Use of this source code is govered by an MIT license
// Licensing information can be found in the LICENSE file
package goGraphicsMagick

// #cgo CFLAGS: -IGM_SRC
// #cgo pkg-config: GraphicsMagickWand
/*
#include <stdio.h>
#include "wand/magick_wand.h"

// Some C helper functions
// Thanks to https://groups.google.com/forum/#!topic/golang-nuts/pQueMFdY0mk
static char**makeCharArray(int size) {
    return calloc(sizeof(char*), size);
}

static void setArrayString(char **a, char *s, int n) {
    a[n] = s;
}

static void freeCharArray(char **a, int size) {
    int i;
    for (i = 0; i < size; i++)
            free(a[i]);
    free(a);
}

// The GraphicsMagick based rotate function
int GMRotate(int argc, char **argv)
{

int i;

for (i = 0; i < argc; i++) {
    printf("%s\n", argv[i]);
}
  MagickWand *magick_wand;
  MagickPassFail status = MagickPass;
  const char *infile, *outfile;
  infile=argv[1];
  outfile=argv[2];
  // Initialize GraphicsMagick API
  InitializeMagick(*argv);
  // Allocate Wand handle
  magick_wand=NewMagickWand();
  // Read input image file
  if (status == MagickPass)
    {
      status = MagickReadImage(magick_wand,infile);
    }
  // Rotate image clockwise 30 degrees with black background
  if (status == MagickPass)
    {
      PixelWand *background;
      background=NewPixelWand();
      PixelSetColor(background,"#000000");
      status = MagickRotateImage(magick_wand,background,30);
      DestroyPixelWand(background);
    }
  // Write output file
  if (status == MagickPass)
    {
      status = MagickWriteImage(magick_wand,outfile);
    }
  // Diagnose any error
  if (status != MagickPass)
    {
      char *description;
      ExceptionType severity;
      description=MagickGetException(magick_wand,&severity);
      (void) fprintf(stderr,"%.1024s (severity %d)\n",
                     description,severity);
    }
  // Release Wand handle
  DestroyMagickWand(magick_wand);
  // Destroy GraphicsMagick API
  DestroyMagick();
  return (status == MagickPass ? 0 : 1);
}
*/
import "C"

import (
	"errors"
)

// Initialise some of the standard variables we will need
var goArgv []string
var goArgc int
var cArgv **C.char
var cArgc C.int

func goSliceStringToCharStar(goSliceString []string) (cCharStar **C.char) {
	cCharStar = C.makeCharArray(C.int(len(goSliceString)))

	for i, s := range goSliceString {
		C.setArrayString(cCharStar, C.CString(s), C.int(i))
	}

	return cCharStar
}

// The rotate function
// Rotates an image and puts on an optional background
func Rotate(srcPath string, destPath string) error {
	var err error = nil

	if srcPath == "" {
		err = errors.New("Failure, there is no source image. Please specify one using -src <src.jpg>.")
	}

	if destPath == "" {
		err = errors.New("Failure, there is no destination image. Please specify one using -dest <dest.png>.")
	}

	// Append the source and destination to argv
	goArgv = append(goArgv, destPath)
	goArgv = append(goArgv, srcPath)

	cArgv = goSliceStringToCharStar(goArgv)
	defer C.freeCharArray(cArgv, C.int(len(goArgv)))
	cArgc = C.int(len(goArgv))

	C.GMRotate(cArgc, cArgv)

	return err
}
