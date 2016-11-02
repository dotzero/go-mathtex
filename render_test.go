// +build extend
// How to use: go test -v -tags extend

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

	MathtexMsgLevel = "0" // Set Verbosity level

	fixturesDir := pwd + `/fixtures`
	files, _ := ioutil.ReadDir(fixturesDir)
	for _, f := range files {
		fixture := fixturesDir + `/` + f.Name()
		log.Printf("Fixture: %s", fixture)

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

		log.Printf("PNG: %s", filename)
	}
}
