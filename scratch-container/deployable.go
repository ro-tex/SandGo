package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var settings struct {
	Verbose bool
}

func handler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	if settings.Verbose {
		fmt.Fprintf(res, "Hello! Your request's body was: %s (verbose!)", string(body))
	} else {
		fmt.Fprintf(res, "Hello! Your request's body was: %s", string(body))
	}
}

func main() {
	verboseFlag := flag.Bool("v", false, "Enable verbose logging.")
	flag.Parse()
	settings.Verbose = *verboseFlag

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
