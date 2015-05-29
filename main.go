package main

import (
	"flag"
	"net/http"
	"strings"
)

var targetHost string

func init() {
	flag.StringVar(&targetHost, "host", "", "The target host")
	flag.Parse()
}

func main() {
	if targetHost == "" {
		flag.Usage()
	} else {
		sendPayload(targetHost)
	}
}

func sendPayload(url string) (resp *http.Response, err error) {
	buf := strings.NewReader(payload)
	return http.Post(url, "application/json", buf)
}
