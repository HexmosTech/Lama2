//go:build wasm

package contoller

import (
	"fmt"
	"syscall/js"
)

var worker js.Value

// func initWebWorker() {
// 	worker = js.Global().Get("Worker").New("worker.js")

// 	worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		result := args[0].Get("data")
// 		// Handle the result from the web worker
// 		fmt.Println("Result from web worker:", result)
// 		return nil
// 	}))
// }

func initWebWorker() js.Value {
	if worker.IsUndefined() {
		worker = js.Global().Get("Worker").New("worker.js")
		worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			result := args[0].Get("data")
			// Handle the result from the web worker
			fmt.Println("Result from web worker:", result)
			return nil
		}))
	}
	return worker
}


func runCodeInWorker(chainCode string) {
	worker.Call("postMessage", chainCode)
}
