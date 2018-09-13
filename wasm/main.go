package main

import (
	"syscall/js"
)

// see https://godoc.org/syscall/js

func add(in []js.Value) {

	// val := js.Global().Get("document").Call("getElementById", "some_id").Get("value").String()

	res := js.ValueOf(in[0].Int() + in[1].Int())
	// js.Global().Set("output", res)
	println(res.String())

	// js.Global().Get("document").Call("getElementById", "some_output_id").Set("value", res)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	done := make(chan struct{}, 0)

	println("Hello, WebAssembly!")
	registerCallbacks()

	<-done
}
