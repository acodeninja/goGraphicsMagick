#Go Go GraphicsMagick

A go application that wraps the Graphics Magick C libraries to produce a concurrent image processing application. 


##Usage

##Compiling
To compile any program that uses this library you will need a local copy of the GraphicsMagick C headers for include.

``` hg clone http://hg.code.sf.net/p/graphicsmagick/code GM_SRC ```

``` cd GM_SRC ```

``` ./configure ``` ( or on OS X ``` ./configure CC=clang --enable-shared```)

``` make ```

You also need go1.2 (available from homebrew on OS X and most repos on Linux)

##Development

###Authors

* Lawrence Goldstien: github.com/lgoldstien | onmylemon.co.uk

###Contribution

##License
