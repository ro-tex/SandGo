## Build:

`GOARCH=wasm GOOS=js go build -o lib.wasm main.go`

## Run just the wasm file:

`sh go_js_wasm_exec lib.wasm`

## Run the server:

`go run server.go`

## Copy over the HTML & JS support boilerplate:

`cp $(go env GOROOT)/misc/wasm/wasm_exec.{html,js} .`

## Minimal JS to run Go Wasm:

```
const go = new Go();
WebAssembly.instantiateStreaming(fetch("bin.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance);
  });
```

## Working with DOM elements:

API docs: https://godoc.org/syscall/js

Get the value of an element:
`val := js.Global().Get("document").Call("getElementById", "some_id").Get("value").String()`

Set the value of a variable:
`js.Global().Set("output", js.ValueOf(5))`

Set the value of an element:
`js.Global().Get("document").Call("getElementById", "some_output_id").Set("value", res)`
