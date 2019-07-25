package main

import (
	"encoding/base64"
	"net/url"
	"strings"
	"syscall/js"
)

func getElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

// TODO:
// - find a wayy to return a value instead of modifying the DOM

func Capitalise(text []js.Value) {
	result := js.ValueOf(strings.ToUpper(text[0].String()))
	outputEl := getElementById(text[1].String())
	outputEl.Set("value", result)
}

func UrlEncode(text []js.Value) {
	outputEl := getElementById(text[1].String())
	result := js.ValueOf(url.PathEscape(text[0].String()))
	outputEl.Set("value", result)
}

func UrlDecode(text []js.Value) {
	outputEl := getElementById(text[1].String())
	result, err := url.PathUnescape(text[0].String())
	if err != nil {
		outputEl.Set("value", js.ValueOf("Failed to unescape! Input: "+text[0].String()))
	}
	println(result)
	outputEl.Set("value", js.ValueOf(result))
}

func Base64Encode(text []js.Value) {
	outputEl := getElementById(text[1].String())
	result := base64.StdEncoding.EncodeToString([]byte(text[0].String()))
	outputEl.Set("value", js.ValueOf(result))
}

func Base64Decode(text []js.Value) {
	outputEl := getElementById(text[1].String())
	result, err := base64.StdEncoding.DecodeString(text[0].String())
	if err != nil {
		outputEl.Set("value", js.ValueOf("Failed to decode! Input: "+text[0].String()))
	}
	outputEl.Set("value", js.ValueOf(string(result)))
}

func registerCallbacks() {
	js.Global().Set("capitalise", js.NewCallback(Capitalise))

	js.Global().Set("urlEncode", js.NewCallback(UrlEncode))
	js.Global().Set("urlDecode", js.NewCallback(UrlDecode))
	js.Global().Set("base64Encode", js.NewCallback(Base64Encode))
	js.Global().Set("base64Decode", js.NewCallback(Base64Decode))
}

func main() {
	done := make(chan struct{}, 0)

	println("Hello, WebAssembly!")
	registerCallbacks()

	<-done
}
