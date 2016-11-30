## go-mathtex

[![Build Status](https://travis-ci.org/dotzero/go-mathtex.svg?branch=master)](https://travis-ci.org/dotzero/go-mathtex)
[![Go Report Card](https://goreportcard.com/badge/github.com/dotzero/go-mathtex)](https://goreportcard.com/report/github.com/dotzero/go-mathtex)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

This is a Go package that allows rendering LaTeX documents to images using [mathTeX](http://www.forkosh.com/mathtex.html). Also parse LaTeX expression and check for dangerous commands and add some instructions for proper rendering.

## Before install

At first of all, you must install `texlive` package, follow the instructions below to install TeX Live on Debian-based linux.

```bash
sudo apt-get update && apt-get install -y --no-install-recommends texlive-full
```

Then compile `mathtex` to binary, follow the instructions below.

```bash
git clone https://github.com/dotzero/go-mathtex
cd go-mathtex/mathtex
make png \
    BIN_PATH=/usr/local/bin/mathtex \
    PATH_CACHE=/var/lib/mathtex/cache/ \
    PATH_WORK=/var/lib/mathtex/work/
```

This commands will compile mathtex to `/usr/local/bin/mathtex` and also create two dirs `/var/lib/mathtex/cache` and `/var/lib/mathtex/work`. If you prefer to render SVG instead of PNG, then follow the instructions below.

```bash
git clone https://github.com/dotzero/go-mathtex
cd go-mathtex/mathtex
make svg \
    BIN_PATH=/usr/local/bin/mathtex \
    PATH_CACHE=/var/lib/mathtex/cache/ \
    PATH_WORK=/var/lib/mathtex/work/
```

## Install

```bash
go get github.com/dotzero/go-mathtex
```

## Usage

```go
package main

import (
    "log"

    mathtex "github.com/dotzero/go-mathtex"
)

func main() {
    mathtex.MathtexPath = "/var/www/mathtex.cgi" // path to mathtex binary
    mathtex.MathtexCachePath = "/var/www/cache/" // path to mathtex cache files
    mathtex.MathtexWorkPath = "/var/www/work/" // path to mathtex work files
    mathtex.MathtexMsgLevel = "0" // mathtex message level
    mathtex.MathtexOutputExt = "png" // mathtex output file extenstion

    expr := `x^2+y^2`

    // Check for pre-rendered file
    if filename, err := mathtex.CheckRenderCache(expr); err == nil {
        log.Printf("Mathtex pre-rendered file: %s\n", filename)
        return
    }

    // Pre-rendered file not found, try to render it
    filename, err := mathtex.RenderImage(expr)
    if err != nil {
        log.Fatalln("Mathtex error: " + err.Error())
        return
    }
    log.Printf("Mathtex rendered file: %s\n", filename)
}
```

## Licenses

* mathTeX: [GNU GPL](mathtex/COPYING)
* go-mathtex: [MIT](LICENSE)
