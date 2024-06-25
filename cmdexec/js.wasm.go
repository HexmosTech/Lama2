//go:build wasm

package cmdexec

import (
	"syscall/js"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func RunVMCode(jsCode string, vm *goja.Runtime) {
	js.Global().Call("eval", jsCode)
}

// GetJSVm creates a new goja runtime instance
// with console.log enabled
func GetJSVm() *goja.Runtime {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	return vm
}

// The GenerateChainCode function creates a JavaScript code snippet
// that attempts to parse a given HTTP response body as JSON, logging
// success or failure, and stores the result accordingly.
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
