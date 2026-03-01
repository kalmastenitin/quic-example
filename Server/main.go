package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requested Via %s\n", r.Proto)
	var response = make(map[string]interface{})
	response["message"] = r.Proto
	w.Header().Set("Content-Type", "text/html")
	json.NewEncoder(w).Encode(response)
}

// AltSvcMiddleware adds the header required for protocol upgrade
func AltSvcMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// "h3" is the identifier for HTTP/3
		// ":443" is the port
		// ma=86400 tells the browser to remember this for 24 hours
		w.Header().Set("Alt-Svc", `h3=":443"; ma=86400`)
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", Hello)

	wrappedMux := AltSvcMiddleware(mux)

	// Serve Http3 over QUIC
	server := http3.Server{
		Handler: wrappedMux,
		Addr:    ":443",
	}

	log.Printf("Starting HTTP/3 server on :443/udp")
	if err := server.ListenAndServeTLS("../host.cert", "../host.key"); err != nil {
		log.Fatalf("QUIC server failed: %v", err)
	}

}
