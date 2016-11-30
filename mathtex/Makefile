# Env
PWD := $(shell pwd)
PREFIX ?= $(PWD)/build

# Relative paths to mathTeX's cache and work dirs
PATH_CACHE ?= $(PREFIX)/cache/
PATH_WORK ?= $(PREFIX)/work/

# Build mathtex.c
SRC_PATH ?= $(PWD)/mathtex.c
BIN_PATH ?= $(PREFIX)/mathtex.cgi

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

all: clean _build_png test

png: _build_png _done

svg: _build_svg _done

test:
	$(BIN_PATH) -o "$(PATH_CACHE)formula" -m 99 "x^2+y^2"
	$(BIN_PATH) -o "$(PATH_CACHE)switches" -m 99 "\switches"
	$(BIN_PATH) -o "$(PATH_CACHE)environment" -m 99 "\environment"

clean:
	rm -f $(BIN_PATH)
	rm -rf $(PATH_CACHE)
	rm -rf $(PATH_WORK)

_build_png:
	@mkdir -p $(PATH_CACHE) $(PATH_WORK)
	$(CC) $(SRC_PATH) $(CFLAGS_PNG) -o $(BIN_PATH)

_build_svg:
	@mkdir -p $(PATH_CACHE) $(PATH_WORK)
	$(CC) $(SRC_PATH) $(CFLAGS_SVG) -o $(BIN_PATH)

_done:
	@echo "------------------------------------------------"
	@echo " Mathtex - Build Successful"
	@echo "------------------------------------------------"
	@echo " Binary: $(BIN_PATH)"
	@echo " Cache path: $(PATH_CACHE)"
	@echo " Working path: $(PATH_WORK)"
	@echo "------------------------------------------------"

.PHONY: all png svg test clean
