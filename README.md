# Go Go GraphicsMagick
A go application that wraps the Graphics Magick C libraries to produce a concurrent image processing application. 

## Performance
I have been able to get conversion and rotation times to around the same speed as the ```gm``` command line utility. I will do further optimisation in future releases to try and bring this time down a little.

## Usage
Check out the ``` examples ``` folder for some example applications. Usage is fairly straight forward.

    err := gm.Convert(_srcImagePath, _destImagePath)
    if err != nil {
            fmt.Println(err)
    }

```_srcImagePath``` and ```_destImagePath``` should both be strings.

## Compiling
To compile any program that uses this library you will need a local copy of the GraphicsMagick C headers for include.

``` hg clone http://hg.code.sf.net/p/graphicsmagick/code GM_SRC ```

``` cd GM_SRC ```

``` ./configure ``` ( or on OS X ``` ./configure CC=clang --enable-shared```)

``` make ```

You also need go1.2 (available from homebrew on OS X and most repos on Linux)

## Development

### Authors

* Lawrence Goldstien: github.com/acodeninja | acode.ninja

### Contribution

##License
MIT License, please see LICENSE file, if none is present please read it [here](https://github.com/acodeninja/goGraphicsMagick/blob/stable/LICENSE)
