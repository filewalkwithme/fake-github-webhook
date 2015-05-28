package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendPayload(t *testing.T) {
	echoServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, string(body))
	}))
	defer echoServer.Close()

	resEchoServer, err := sendPayload(echoServer.URL)

	if err != nil {
		t.Fatalf("Error sending payload to echoServer: %v", err)
	}

	respPayload, err := ioutil.ReadAll(resEchoServer.Body)
	resEchoServer.Body.Close()
	if err != nil {
		t.Fatalf("Error receiving payload response from EchoServer: %v", err)
	}

	if string(respPayload) != payload {
		t.Fatal("Payload sent and received from EchoServer are not equals.")
	}
}
