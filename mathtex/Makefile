# Env
PWD := $(shell pwd)
PREFIX ?= /var/www

# Build mathtex.c
CC_SRC_PATH ?= $(PWD)/mathtex.c
CC_BIN_PATH ?= $(PREFIX)/mathtex.cgi

# Relative paths to mathTeX's cache and work dirs
PATH_CACHE ?= $(PREFIX)/cache/
PATH_WORK ?= $(PREFIX)/work/

# Compile flags
CC ?= cc
CFLAGS = -DLATEX=\"`which latex`\" \
	-DDVIPNG=\"`which dvipng`\" \
	-DCACHE=\"$(PATH_CACHE)\" \
	-DWORK=\"$(PATH_WORK)\" \
	-DWHITEPACKAGES=\"whitepackages\" \
	-DDPI=\"120\" \
	-DPNG

all: build

build:
	mkdir -p $(PREFIX) $(PATH_CACHE) $(PATH_WORK)
	$(CC) $(CC_SRC_PATH) $(CFLAGS) -o $(CC_BIN_PATH)

test:
	$(CC_BIN_PATH) -o "$(PATH_CACHE)test" -m 99 "x^2+y^2"
	@echo "---------------------------------------------"
	@echo "File: $(PATH_CACHE)test.png"
	@ls -lah $(PATH_CACHE)test.png
	@echo "---------------------------------------------"

clean:
ifneq ("$(wildcard $(CC_BIN_PATH))","")
	rm $(CC_BIN_PATH)
endif
ifneq ("$(wildcard $(PATH_CACHE))","")
	rm -rf $(PATH_CACHE)
endif
ifneq ("$(wildcard $(PATH_WORK))","")
	rm -rf $(PATH_WORK)
endif

.PHONY: all build test clean
