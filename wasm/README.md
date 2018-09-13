## Build:

`GOARCH=wasm GOOS=js go build -o test.wasm main.go`

## Run just the wasm file:

`sh go_js_wasm_exec test.wasm`

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
