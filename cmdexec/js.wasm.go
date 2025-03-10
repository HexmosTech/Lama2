//go:build wasm

package cmdexec

import (
	"fmt"
	"syscall/js"

	"github.com/rs/zerolog/log"
)

// RunVMCode takes in a JS snippet as a string, executes the code in a JS VM using Web Worker
func RunVMCode(jsCode string, vm interface{}) {
	fmt.Println("WASM: Running VM code")
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
	fmt.Println("WASM: VM code executed successfully")
}

// GenerateChainCode takes in an HTTP response body and comes up with some JS code to define the "magic variable" result.
func GenerateChainCode(httpRespBody string) string {
	fmt.Println("WASM: Generating chain code")
	code := `try {
		result = JSON.parse(String.raw` + "`" + httpRespBody + "`" + `)	
		console.log("Stored as JSON")
	} catch (e) {
		result = String.raw` + "`" + httpRespBody + "`" + ` 	
		console.log(e)
		console.log("Stored as string")
	}`
	fmt.Println("WASM: Chain code generated", code)
	return code
}
