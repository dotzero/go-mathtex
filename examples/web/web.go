package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	mathtex "github.com/dotzero/go-mathtex"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func responseJSON(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := response{Code: code, Message: message}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error with json encoder")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.RawQuery) < 1 {
		responseJSON(w, http.StatusBadRequest, "Bad formula")
		return
	}

	filename, err := mathtex.RenderImage(r.URL.RawQuery)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, "Bad formula")
		log.Println("Mathtex RenderImage error: " + err.Error())
		return
	}

	img, err := ioutil.ReadFile(filename)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, "Bad formula")
		log.Println("ReadFile error: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(img)))
	if _, err := w.Write(img); err != nil {
		fmt.Println("Failed to reply with image")
	}
}

func main() {
	portFlag := flag.Int("p", 8888, "Port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*portFlag), nil))
}
