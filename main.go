package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	buf := strings.NewReader(payload)
	resp, err := http.Post("http://127.0.0.1:8080", "application/json", buf)
	fmt.Printf("%v\n", resp)
	fmt.Printf("%v\n", err)
}
