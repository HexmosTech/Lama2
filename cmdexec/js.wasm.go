//go:build wasm

package cmdexec

import (
	"fmt"
	"syscall/js"
	"github.com/rs/zerolog/log"
)

var worker js.Value

// InitWebWorker initializes the Web Worker and returns the worker instance.
func InitWebWorker() js.Value {
	if worker.IsUndefined() {
		worker = js.Global().Get("Worker").New("worker.js")
		worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			result := args[0].Get("data").Get("result")
			err := args[0].Get("data").Get("error")
			if err.String() != "null" {
				fmt.Println("Error from web worker:", err)
			} else {
				fmt.Println("Result from web worker:", result)
			}
			return nil
		}))
	}
	return worker
}

// GetJSVm initializes the Web Worker (this replaces the Goja VM creation)
func GetJSVm() interface{} {
	return InitWebWorker()
}

// RunVMCode takes in a JS snippet as a string, executes the code in a JS VM using Web Worker
func RunVMCode(jsCode string, vm interface{}) {
	worker := vm.(js.Value)
	resultChan := make(chan js.Value)
	worker.Call("postMessage", map[string]interface{}{
		"code": jsCode,
	})

	js.Global().Get("addEventListener").Invoke("message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resultChan <- args[0].Get("data")
		return nil
	}))

	result := <-resultChan
	err := result.Get("error").String()
	if err != "null" {
		log.Fatal().Str("Error executing JS processor block", err).Msg("")
	}
}

// GenerateChainCode takes in an HTTP response body and comes up with some JS code to define the "magic variable" result.
func GenerateChainCode(httpRespBody string) string {
	fmt.Println("WW: Value of httpRespBody is",httpRespBody)
	code := `try {
		result = JSON.parse(String.raw` + "`" + httpRespBody + "`" + `)	
		console.log("Stored as JSON")
	} catch (e) {
		result = String.raw` + "`" + httpRespBody + "`" + ` 	
		console.log(e)
		console.log("Stored as string")
	}`
	log.Debug().Str("Chain code generated", code).Msg("")
	return code
}
