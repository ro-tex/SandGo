package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	var srv = http.Server{Addr: ":6080"}
	http2.ConfigureServer(&srv, nil) //Enable http2

	srv.ListenAndServeTLS("certs/localhost.cert", "certs/localhost.key")
}
