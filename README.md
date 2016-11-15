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
make build
```

This commands will compile mathtex to `/var/www/mathtex.cgi` and also create two dirs `/var/www/cache` and `/var/www/work`.
You can change this behavior with `make build PREFIX="/path/to/"`.

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

    filename, err := mathtex.RenderImage(`x^2+y^2`)
    if err != nil {
        log.Fatalln("Mathtex error: " + err.Error())
    }
    log.Printf("Mathtex image: %s\n", filename)
}
```

## Licenses

* mathTeX: [GNU GPL](mathtex/COPYING)
* go-mathtex: [MIT](LICENSE)
