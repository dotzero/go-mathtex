package mathtex

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	// MathtexPath contains path to mathtex.cgi
	MathtexPath = "/var/www/mathtex.cgi"
	// MathtexCachePath contains path to mathtex.cgi cache
	MathtexCachePath = "/var/www/cache/"
	// MathtexMsgLevel contains mathtex.cgi message level
	MathtexMsgLevel = "99"
)

// RenderImage render LaTeX expression to PNG
func RenderImage(expr string) (string, error) {
	var (
		cmdArgs []string
		cmdOut  []byte
		err     error
	)

	fileBase := MathtexCachePath + md5hash(expr)
	fileExt := ".png"
	filename := fileBase + fileExt

	err = checkBlackList(expr)
	if err != nil {
		return "", err
	}

	expr = analyzeLatex(expr)
	cmdArgs = []string{expr, "-m", MathtexMsgLevel, "-o", fileBase}
	if cmdOut, err = exec.Command(MathtexPath, cmdArgs...).Output(); err != nil {
		return "", err
	}

	// Debug only
	if MathtexMsgLevel != "0" {
		fmt.Println(string(cmdOut))
	}

	if flag, err := exists(filename); flag != true || err != nil {
		return "", fmt.Errorf("Failed expression `%s`", expr)
	}

	return filename, nil
}

// checkBlackList parse expression and check for dangerous commands
func checkBlackList(expr string) error {
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
		if strings.Contains(expr, word) == true {
			return fmt.Errorf("Sorry, command %s is not available.", word)
		}
	}

	return nil
}

// analyzeLatex parse expression and add usepackage
func analyzeLatex(expr string) string {
	eol := "\n"

	conditions := map[string]string{
		`eqnarray`:        `eqnarray`,
		`sequencediagram`: `pgf-umlsd`,
	}

	for cmd := range conditions {
		if strings.Contains(expr, `\begin{`+cmd+`}`) == true {
			expr = `\usepackage{` + conditions[cmd] + `}` + eol + expr
		} else if strings.Contains(expr, `\begin{`+cmd+`*}`) == true {
			expr = `\usepackage{` + conditions[cmd] + `}` + eol + expr
		}
	}

	if strings.Contains(expr, `\addplot`) == true {
		expr = `\usepackage{pgfplots}` + eol + expr
	}

	if strings.Contains(expr, `\xymatrix`) == true {
		expr = `\usepackage[all]{xy}` + eol + expr
	} else if strings.Contains(expr, `\begin{xy}`) == true {
		expr = `\usepackage[all]{xy}` + eol + expr
	}

	if strings.Contains(expr, `picture`) == true {
		expr = `\setlength{\unitlength}{1pt}` + eol + expr
	}

	if strings.Contains(expr, `\begin{align`) == true {
		expr = `\parmode` + eol + expr
	}

	return expr
}

// md5hash calculate the md5 hash of a string
func md5hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
