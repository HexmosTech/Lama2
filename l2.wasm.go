//go:build wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	controller "github.com/HexmosTech/lama2/controller"
)

var version string

func init() {
	// Initialize version if not set
	if len(version) == 0 {
		version = "vUnset"
	}
}

func main() {
	// Set up the WebAssembly function
	js.Global().Set("lama2Wasm", js.FuncOf(lama2Wasm))
	js.Global().Set("goWebRequestFunc", wasmLamaPromise())
	select {}
}

func lama2Wasm(this js.Value, i []js.Value) interface{} {
	go func() {
		controller.ProcessWasmInput("GET\nhttps://httpbin.org/get")
	}()
	return "data from api:"
}

func wasmLamaPromise() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		inputdata := args[0].String()
		// proxyinputdata := args[1].String()
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			go func() {
				result := controller.ProcessWasmInput(inputdata)
				resultJSON, err := json.Marshal(result)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				resolve.Invoke(js.ValueOf(string(resultJSON)))
			}()
			return nil
		})
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
