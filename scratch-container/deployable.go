package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(res, "Hello! Your request's body was: %s", string(body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
