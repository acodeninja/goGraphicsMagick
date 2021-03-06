#include <stdio.h>
#include "wand/magick_wand.h"

#define ZERO 48
#define NINE 57
#define MINUS 45
#define DECPNT 46

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

long strtolng_n(char* str, int n)
{
    int sign = 1;
    int place = 1;
    long ret = 0;

    int i;
    for (i = n-1; i >= 0; i--, place *= 10)
    {
        int c = str[i];
        switch (c)
        {
            case MINUS:
                if (i == 0) sign = -1;
                else return -1;
                break;
            default:
                if (c >= ZERO && c <= NINE) ret += (c - ZERO) * place;
                else return -1;
        }
    }

    return sign * ret;
}

double _double_fraction(char* str, int n)
{
    double place = 0.1;
    double ret = 0.0;

    int i;
    for (i = 0; i < n; i++, place /= 10)
    {
        int c = str[i];
        ret += (c - ZERO) * place;
    }
    return ret;
}

double strtodbl(char* str)
{
    int n = 0;
    int sign = 1;
    int d = -1;
    long ret = 0;

    char* temp = str;
    while (*temp != '\0')
    {
        switch (*temp)
        {
            case MINUS:
                if (n == 0) sign = -1;
                else return -1;
                break;
            case DECPNT:
                if (d == -1) d = n;
                else return -1;
                break;
            default:
                if (*temp < ZERO && *temp > NINE) return -1;
        }
        n++;
        temp++;
    }

    if (d == -1)
    {
        return (double)(strtolng_n(str, n));
    }
    else if (d == 0)
    {
        return _double_fraction((str+d+1), (n-d-1));
    }
    else if (sign == -1 && d == 1)
    {
        return (-1)*_double_fraction((str+d+1), (n-d-1));
    }
    else if (sign == -1)
    {
        ret = strtolng_n(str+1, d-1);
        return (-1) * (ret + _double_fraction((str+d+1), (n-d-1)));
    }
    else
    {
        ret = strtolng_n(str, d);
        return ret + _double_fraction((str+d+1), (n-d-1));
    }
}

// The GraphicsMagick based rotate function
int ImageRotate(int argc, char **argv)
{
  MagickWand *magick_wand;
  MagickPassFail status = MagickPass;
  const char *infile, *outfile;
  double degrees;
  infile=argv[1];
  outfile=argv[2];
  degrees=strtodbl(argv[3]);
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
      printf("MagickRotateImage by %f degrees\n", degrees);
      status = MagickRotateImage(magick_wand,background,degrees);
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

// The GraphicsMagick based image conversion function
int ImageConvert(int argc, char **argv)
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