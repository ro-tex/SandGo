package main

import (
	"fmt"
	"syscall/js"
)

// see https://godoc.org/syscall/js

func add(in []js.Value) {
	js.Global().Set("output", js.ValueOf(in[0].Int()+in[1].Int()))
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	c := make(chan struct{}, 0)

	println("Hello, WebAssembly!")
	registerCallbacks()

	<-c
}
