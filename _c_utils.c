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
int ImageRotate(int argc, char **argv)
{
  MagickWand *magick_wand;
  MagickPassFail status = MagickPass;
  const char *infile, *outfile;
  infile=argv[1];
  outfile=argv[2];
  printf("%s\n", infile);
  printf("%s\n", outfile);
  // Initialize GraphicsMagick API
  printf("InitializeMagick...\n");
  InitializeMagick(*argv);
  // Allocate Wand handle
  printf("NewMagickWand\n");
  magick_wand=NewMagickWand();
  // Read input image file
  if (status == MagickPass)
    {
      printf("MagickPass\n");
      printf("MagickReadImage\n");
      status = MagickReadImage(magick_wand,infile);
    }
  // Rotate image clockwise 30 degrees with black background
  if (status == MagickPass)
    {
      printf("MagickPass\n");
      printf("PixelWand\n");
      PixelWand *background;
      printf("NewPixelWand\n");
      background=NewPixelWand();
      printf("PixelSetColor\n");
      PixelSetColor(background,"#000000");
      printf("MagickRotateImage\n");
      status = MagickRotateImage(magick_wand,background,30);
      printf("DestroyPixelWand\n");
      DestroyPixelWand(background);
    }
  // Write output file
  if (status == MagickPass)
    {
      printf("MagickPass\n");
      printf("MagickWriteImage\n");
      printf("%s\n", outfile);
      status = MagickWriteImage(magick_wand,outfile);
    }
  // Diagnose any error
  if (status != MagickPass)
    {
      printf("!MagickPass\n");
      char *description;
      ExceptionType severity;
      printf("MagickGetException\n");
      description=MagickGetException(magick_wand,&severity);
      (void) fprintf(stderr,"%.1024s (severity %d)\n",
                     description,severity);
    }
  // Release Wand handle
  printf("DestroyMagickWand\n");
  DestroyMagickWand(magick_wand);
  // Destroy GraphicsMagick API
  printf("DestroyMagick\n");
  DestroyMagick();
  return (status == MagickPass ? 0 : 1);
}