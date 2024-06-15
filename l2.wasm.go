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
	//Set the global JavaScript property "goWebRequestFunc" to the result of wasmLamaPromise
	js.Global().Set("goWebRequestFunc", wasmLamaPromise())
	// Block the main function to keep the Go WebAssembly running
	select {}
}

func wasmLamaPromise() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		inputdata := args[0].String()
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			go func() {
				result, _ := controller.ProcessWasmInput(inputdata)
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
