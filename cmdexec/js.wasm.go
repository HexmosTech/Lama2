//go:build wasm

package cmdexec

import (
	"syscall/js"

	"github.com/dop251/goja"
)

func RunVMCode(jsCode string, vm *goja.Runtime) {
	js.Global().Call("eval", jsCode)
}

//  The GenerateChainCode function creates a JavaScript code snippet
//  that attempts to parse a given HTTP response body as JSON, logging
//  success or failure, and stores the result accordingly.
func GenerateChainCode(httpRespBody string) string {
	code := `try {
		result = JSON.parse(String.raw` + "`" + httpRespBody + "`" + `)	
		console.log("Stored as JSON")
	} catch (e) {
		result = String.raw` + "`" + httpRespBody + "`" + ` 	
		console.log(e)
		console.log("Stored as string")
	}`
	return code
}
