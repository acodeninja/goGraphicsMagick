// Copyright 2014 Lawrence Goldstien lgoldstien@onmylemon.co.uk
// Use of this source code is govered by an MIT license
// Licensing information can be found in the LICENSE file
package goGraphicsMagick

// #cgo CFLAGS: -IGM_SRC
// #cgo pkg-config: GraphicsMagickWand
// #include <stdio.h>
// #include "wand/magick_wand.h"
// int Rotate(char srcPath,char destPath)
// {
//   MagickWand *magick_wand;
//   MagickPassFail status = MagickPass;
//   const char *infile, *outfile;
//   infile=srcPath;
//   outfile=destPath;
//   // Initialize GraphicsMagick API
//   InitializeMagick(*argv);
//   // Allocate Wand handle
//   magick_wand=NewMagickWand();
//   // Read input image file
//   if (status == MagickPass)
//     {
//       status = MagickReadImage(magick_wand,infile);
//     }
//   // Rotate image clockwise 30 degrees with black background
//   if (status == MagickPass)
//     {
//       PixelWand *background;
//       background=NewPixelWand();
//       PixelSetColor(background,"#000000");
//       status = MagickRotateImage(magick_wand,background,30);
//       DestroyPixelWand(background);
//     }
//   // Write output file
//   if (status == MagickPass)
//     {
//       status = MagickWriteImage(magick_wand,outfile);
//     }
//   // Diagnose any error
//   if (status != MagickPass)
//     {
//       char *description;
//       ExceptionType severity;
//       description=MagickGetException(magick_wand,&severity);
//       (void) fprintf(stderr,"%.1024s (severity %d)\n",
//                      description,severity);
//     }
//   // Release Wand handle
//   DestroyMagickWand(magick_wand);
//   // Destroy GraphicsMagick API
//   DestroyMagick();
//   return (status == MagickPass ? 0 : 1);
// }
import "C"

import (
	"errors"
)

func Rotate(srcPath string, destPath string) error {
	var err error = nil

	if srcPath == "" {
		err = errors.New("Failure, there is no source image. Please specify one using -src <src.jpg>.")
	}

	if destPath == "" {
		err = errors.New("Failure, there is no destination image. Please specify one using -dest <dest.png>.")
	}

	C.Rotate()

	return err
}
