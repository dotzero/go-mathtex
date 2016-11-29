// +build testpng
// How to use: go test -v -tags testpng

package mathtex

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestRenderImage(t *testing.T) {
	var (
		err      error
		filename string
	)

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(err.Error())
	}

	MathtexPath = pwd + "/mathtex/build/mathtex.cgi"
	MathtexCachePath = pwd + "/mathtex/build/cache/"
	MathtexWorkPath = pwd + "/mathtex/build/work/"
	MathtexMsgLevel = "0"
	MathtexOutputExt = "png"

	fixturesDir := pwd + `/fixtures`
	files, _ := ioutil.ReadDir(fixturesDir)
	for _, f := range files {
		fixture := fixturesDir + `/` + f.Name()
		log.Printf("Input fixture: %s", fixture)

		content, err := ioutil.ReadFile(fixture)
		if err != nil {
			t.Fatalf(err.Error())
		}

		filename, err = RenderImage(string(content))
		if err != nil {
			log.Println("Failed expression: " + string(content))
			t.Fatalf("RenderImage error: " + err.Error())
		}

		filename, err = CheckRenderCache(string(content))
		if err != nil {
			t.Fatalf("CheckRenderCache error: " + err.Error())
		}

		log.Printf("Output %s: %s", MathtexOutputExt, filename)
	}
}
