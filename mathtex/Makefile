# Env
PWD := $(shell pwd)
PREFIX ?= $(PWD)/build

# Build mathtex.c
CC_SRC_PATH ?= $(PWD)/mathtex.c
CC_BIN_PATH ?= $(PREFIX)/mathtex.cgi

# Relative paths to mathTeX's cache and work dirs
PATH_CACHE ?= $(PREFIX)/cache/
PATH_WORK ?= $(PREFIX)/work/

# Compile flags
CC ?= cc
CFLAGS_PNG = -DLATEX=\"`which latex`\" \
	-DDVIPNG=\"`which dvipng`\" \
	-DCACHE=\"$(PATH_CACHE)\" \
	-DWORK=\"$(PATH_WORK)\" \
	-DWHITEPACKAGES=\"whitepackages\" \
	-DDPI=\"120\" \
	-DPNG
CFLAGS_SVG = -DLATEX=\"`which latex`\" \
	-DDVISVGM=\"`which dvisvgm`\" \
	-DCACHE=\"$(PATH_CACHE)\" \
	-DWORK=\"$(PATH_WORK)\" \
	-DWHITEPACKAGES=\"whitepackages\" \
	-DSVG

all: clean build test

build: png

png:
	mkdir -p $(PREFIX) $(PATH_CACHE) $(PATH_WORK)
	$(CC) $(CC_SRC_PATH) $(CFLAGS_PNG) -o $(CC_BIN_PATH)

svg:
	mkdir -p $(PREFIX) $(PATH_CACHE) $(PATH_WORK)
	$(CC) $(CC_SRC_PATH) $(CFLAGS_SVG) -o $(CC_BIN_PATH)

test:
	$(CC_BIN_PATH) -o "$(PATH_CACHE)formula" -m 99 "x^2+y^2"
	$(CC_BIN_PATH) -o "$(PATH_CACHE)switches" -m 99 "\switches"
	$(CC_BIN_PATH) -o "$(PATH_CACHE)environment" -m 99 "\environment"

clean:
	rm -f $(CC_BIN_PATH)
	rm -rf $(PATH_CACHE)
	rm -rf $(PATH_WORK)

.PHONY: build png svg test clean
