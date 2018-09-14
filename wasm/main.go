package main

import (
	"strings"
	"syscall/js"
)

func getElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

// TODO:
// - urlencode/decode
// - base64 encode/decode

func capitalise(text []js.Value) {
	result := js.ValueOf(strings.ToUpper(text[0].String()))
	outputEl := getElementById(text[1].String())
	outputEl.Set("value", result)
}

func registerCallbacks() {
	js.Global().Set("capitalise", js.NewCallback(capitalise))
}

func main() {
	done := make(chan struct{}, 0)

	println("Hello, WebAssembly!")
	registerCallbacks()

	<-done
}
