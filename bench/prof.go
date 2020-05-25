package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

func main() {
	go func() {
		fmt.Println(
			http.ListenAndServe(
				"localhost:6060",
				http.HandlerFunc(pprof.Index),
			),
		)
	}()
	// go tool pprof http://localhost:6060/debug/pprof/heap
}
