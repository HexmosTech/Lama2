package cmdexec

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/rs/zerolog/log"
)

// GetJSVm creates a new goja runtime instance
// with console.log enabled
func GetJSVm() *goja.Runtime {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	return vm
}

// RunVMCode takes in a JS snippet as a string,
// executes the code in a JS VM, finally checks
// whether there are any errors, and if yes,
// logs the problem.
// Note: the vm runtime remains modified; so if
// you reuse the vm for other operations, the state
// from previous invocations carry over
func RunVMCode(jsCode string, vm *goja.Runtime) {
	_, err := vm.RunString(jsCode)
	if ex, ok := err.(*goja.Exception); ok {
		log.Fatal().Str("Error executing JS processor block", ex.String()).Msg("")
	}
}

// GenerateChainCode takes in an HTTP response body
// and comes up with some JS code to define the
// "magic variable" result. What does the code do?
// The result is stored as a JS object, if the input
// value can be parsed as JSON. Otherwise the value is
// stored as a simple string.
func GenerateChainCode(httpRespBody string) string {
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
