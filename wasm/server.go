package main

import (
	"flag"
	"log"
	"net/http"
	// "strings"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

// func handler(resp http.ResponseWriter, req *http.Request) {
// 	if strings.HasSuffix(req.URL.Path, ".wasm") {
// 		resp.Header().Set("content-type", "application/wasm")
// 	}
//
// 	http.FileServer(http.Dir(*dir)).ServeHTTP(resp, req)
// }

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	// log.Fatal(http.ListenAndServe(*listen, http.HandlerFunc(handler)))
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}
