package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {

	// similar to curl -k
	transport := &http3.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			NextProtos:         []string{"h3"},
		},
	}
	defer transport.Close()

	// Initialize Http Client
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get("https://127.0.0.1/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Printf("Protocol: %s\n", resp.Proto)
	fmt.Printf("Status: %s\n", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Body: %s\n", body[:100])
}
