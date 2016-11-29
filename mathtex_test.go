package mathtex

import (
	"strings"
	"testing"
)

func TestFileOutStruct(t *testing.T) {
	f1 := FileOut{Base: "/", Name: "foo", Ext: "bar"}
	if f1.fullname() != "/foo.bar" {
		t.Fatalf("TestFileOutStruct: `%s` == `/foo.bar` - failed", f1.fullname())
	}
	if f1.outpath() != "/foo" {
		t.Fatalf("TestFileOutStruct: `%s` == `/foo` - failed", f1.outpath())
	}

	f2 := FileOut{Name: "foobar"}
	if f2.fullname() != "foobar." {
		t.Fatalf("TestFileOutStruct: `%s` == `foobar` - failed", f2.fullname())
	}
	if f2.outpath() != "foobar" {
		t.Fatalf("TestFileOutStruct: `%s` == `foobar` - failed", f2.outpath())
	}
}

func TestExists(t *testing.T) {
	conditions := map[string]bool{
		`/`:                true,
		`./mathtex.go`:     true,
		`/foo/bar/foo.bar`: false,
		`foobar`:           false,
	}

	for n := range conditions {
		if flag, _ := exists(n); flag != conditions[n] {
			t.Fatalf("TestExists MathtexPath: %s - failed", MathtexPath)
		}
	}
}

func TestMd5hash(t *testing.T) {
	conditions := map[string]string{
		``:       `d41d8cd98f00b204e9800998ecf8427e`,
		`foobar`: `3858f62230ac3c915f300c664312c63f`,
	}

	for n := range conditions {
		if md5hash(n) != conditions[n] {
			t.Fatalf("TestMd5hash `%s` == `%s` - failed", n, conditions[n])
		}
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
		if out := AnalyzeLatex(cmd); strings.Contains(out, conditions[cmd]) == false {
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
		if err := CheckBlackList(word); err == nil {
			t.Fatalf("TestCheckBlackList %s - failed", word)
		}
	}
	if err := CheckBlackList(``); err != nil {
		t.Fatalf("TestCheckBlackList empty - failed")
	}
}
