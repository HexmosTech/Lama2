package main

import (
	contoller "github.com/HexmosTech/lama2/controller"
	"syscall/js"
)

var version string

func main() {
	if len(version) == 0 {
		version = "vUnset"
	}
	// contoller.Process(version)
	// apiContent := "GET\nhttps://httpbin.org/get"
	// contoller.ProcessWebInput(apiContent, version)
	js.Global().Set("wasmLamaRequest", js.FuncOf(parseWasmRequest))
	js.Global().Set("helloWasm", js.FuncOf(hello))
	select {} 
}

func parseWasmRequest(this js.Value,args []js.Value) interface{} {
	contoller.ProcessWasmInput(args[0].String())
	return js.ValueOf("Processed")
}

func hello(this js.Value, p []js.Value) interface{} {
    return js.ValueOf("Hello from Go WASM!")
}
