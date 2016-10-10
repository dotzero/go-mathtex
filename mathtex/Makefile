PREFIX ?= /var/www/

SRC_NAME ?= mathtex.c
BIN_NAME ?= $(PREFIX)mathtex.cgi

PATH_CACHE ?= $(PREFIX)cache/
PATH_WORK ?= $(PREFIX)tmp/

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
	$(CC) $(SRC_NAME) $(CFLAGS) -o $(BIN_NAME)

clean:
ifneq ("$(wildcard $(BIN_NAME))","")
	rm $(BIN_NAME)
endif
ifneq ("$(wildcard $(PATH_CACHE))","")
	rm -rf $(PATH_CACHE)
endif
ifneq ("$(wildcard $(PATH_WORK))","")
	rm -rf $(PATH_WORK)
endif

.PHONY: all build clean
