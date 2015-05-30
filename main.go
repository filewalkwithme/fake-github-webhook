package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var targetHost string
var payloadFile string

func init() {
	flag.StringVar(&targetHost, "host", "", "The target host")
	flag.StringVar(&payloadFile, "file", "", "The payload file")
	flag.Parse()
}

func main() {
	if targetHost == "" {
		flag.Usage()
	} else {
		sendPayload(targetHost, payloadFile)
	}
}

func sendPayload(url string, file string) (resp *http.Response, err error) {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	buf := strings.NewReader(string(fileContents))
	return http.Post(url, "application/json", buf)
}
