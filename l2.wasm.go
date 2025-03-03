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
	fmt.Println("WASM: Starting main")
	// Set the global JavaScript property "goWebRequestFunc" to the result of wasmLamaPromise
	js.Global().Set("goWebRequestFunc", wasmLamaPromise())
	js.Global().Set("goCmdConvertFunc", wasmCodeConverter())

	// Block the main function to keep the Go WebAssembly running
	select {}
}

func wasmLamaPromise() js.Func {
	fmt.Println("WASM: Creating wasmLamaPromise")
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		liveapiManifest := args[2].String()
		js.Global().Set("liveapiManifest", liveapiManifest)
		inputdata := args[0].String()
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			go func() {
				result, _, _, _, _, _ := controller.ProcessWasmInput(inputdata)
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

func wasmCodeConverter() js.Func {
	fmt.Println("WASM: Creating wasmCodeConverter")
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		inputdata := args[0].String()
		convertLang := args[1].String()
		fmt.Println("inputdata", inputdata)
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			go func() {
				result, _ := controller.ProcessConverterInput(inputdata, convertLang)
				resolve.Invoke(js.ValueOf(result))
			}()
			return nil
		})
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
