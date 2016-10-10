package main

import (
	"flag"
	"log"

	mathtex "github.com/dotzero/go-mathtex"
)

func main() {
	flag.Parse()
	expr := flag.Arg(0)

	mathtex.MathtexMsgLevel = "0"
	filename, err := mathtex.RenderImage(expr)
	if err != nil {
		log.Fatalln("Mathtex error: " + err.Error())
	}
	log.Printf("Mathtex image: %s\n", filename)
}
