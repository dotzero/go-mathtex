package mathtex

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestExists(t *testing.T) {
	if flag, _ := exists(MathtexPath); flag != true {
		t.Fatalf("TestExists MathtexPath: %s - failed", MathtexPath)
	}
	if flag, _ := exists(`/foo/bar/foo.bar`); flag != false {
		t.Fatalf("TestExists /foo/bar/foo.bar - failed")
	}
}

func TestMd5hash(t *testing.T) {
	if md5hash(``) != `d41d8cd98f00b204e9800998ecf8427e` {
		t.Fatalf("TestMd5hash empty - failed")
	}
	if md5hash(`foobar`) != `3858f62230ac3c915f300c664312c63f` {
		t.Fatalf("TestMd5hash foobar - failed")
	}
}

func TestAnalyzeLatex(t *testing.T) {
	conditions := map[string]string{
		`\begin{eqnarray}`:         `\usepackage{eqnarray}`,
		`\begin{eqnarray*}`:        `\usepackage{eqnarray}`,
		`\begin{sequencediagram}`:  `\usepackage{pgf-umlsd}`,
		`\begin{sequencediagram*}`: `\usepackage{pgf-umlsd}`,
		`\addplot`:                 `\usepackage{pgfplots}`,
		`\xymatrix`:                `\usepackage[all]{xy}`,
		`\begin{xy}`:               `\usepackage[all]{xy}`,
		`\begin{picture}(76,20)`:   `\setlength{\unitlength}{1pt}`,
		`\begin{align`:             `\parmode`,
	}

	for cmd := range conditions {
		if out := analyzeLatex(cmd); strings.Contains(out, conditions[cmd]) == false {
			t.Fatalf("TestAnalyzeLatex %s - failed", cmd)
		}
	}
}

func TestCheckBlackList(t *testing.T) {
	var blacklist = []string{
		`\input`,
		`\write`,
		`\usepackage`,
		`\dpi`,
		`\dvips`,
		`\dvipng`,
		`\noquiet`,
		`\msglevel`,
		`\which`,
		`\switches`,
		`\eval`,
		`\advertisement`,
		`\version`,
		`\environment`,
	}

	for _, word := range blacklist {
		if err := checkBlackList(word); err == nil {
			t.Fatalf("TestCheckBlackList %s - failed", word)
		}
	}
	if err := checkBlackList(``); err != nil {
		t.Fatalf("TestCheckBlackList empty - failed")
	}
}

func TestRenderImage(t *testing.T) {
	var err error

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
			log.Println("Failed expression: " + string(content))
			t.Fatalf(err.Error())
		}

		filename, err := RenderImage(string(content))
		if err != nil {
			log.Println("Failed expression: " + string(content))
			t.Fatalf("Mathtex error: " + err.Error())
		}
		log.Printf("PNG: %s", filename)
	}
}
