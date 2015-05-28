package main

import (
	"net/http"
	"strings"
)

func main() {
	sendPayload("http://127.0.0.1:8080")
}

func sendPayload(url string) (resp *http.Response, err error) {
	buf := strings.NewReader(payload)
	return http.Post(url, "application/json", buf)
}
