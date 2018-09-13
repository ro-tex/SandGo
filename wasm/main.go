package main

import (
	"strings"
	"syscall/js"
)

func add(in []js.Value) {
	res := js.ValueOf(in[0].Int() + in[1].Int())
	println(res.String())
}

// TODO:
// - urlencode/decode
// - base64 encode/decode

func capitalise(text []js.Value) {
	println("after", strings.ToUpper(text[0].String()))
	js.Global().Set("result", js.ValueOf(strings.ToUpper(text[0].String())))
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("capitalise", js.NewCallback(capitalise))
}

func main() {
	done := make(chan struct{}, 0)

	println("Hello, WebAssembly!")
	registerCallbacks()

	<-done
}
