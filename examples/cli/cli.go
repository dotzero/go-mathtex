package main

import (
	"flag"
	"log"

	mathtex "github.com/dotzero/go-mathtex"
)

// make png \
//     BIN_PATH=/usr/local/bin/mathtex \
//     PATH_CACHE=/var/lib/mathtex/cache/ \
//     PATH_WORK=/var/lib/mathtex/work/

func main() {
	flag.Parse()
	expr := flag.Arg(0)

	mathtex.MathtexPath = "/usr/local/bin/mathtex"
	mathtex.MathtexCachePath = "/var/lib/mathtex/cache/"
	mathtex.MathtexWorkPath = "/var/lib/mathtex/work/"
	mathtex.MathtexMsgLevel = "99"
	mathtex.MathtexOutputExt = "png"

	filename, err := mathtex.RenderImage(expr)
	if err != nil {
		log.Fatalln("Mathtex error: " + err.Error())
	}
	log.Printf("Mathtex image: %s\n", filename)
}
