package cmdexec

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/rs/zerolog/log"
)

func GetJSVm() *goja.Runtime {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	return vm
}

func RunVMCode(chainCode string, vm *goja.Runtime) {
	_, err := vm.RunString(chainCode)
	if ex, ok := err.(*goja.Exception); ok {
		log.Fatal().Str("Error executing JS processor block", ex.String()).Msg("")
	}
}

func GenerateChainCode(resp string) string {
	code := `try {
		result = JSON.parse(String.raw` + "`" + resp + "`" + `)	
		console.log("Stored as JSON")
	} catch (e) {
		result = String.raw` + "`" + resp + "`" + ` 	
		console.log(e)
		console.log("Stored as string")
	}`
	log.Debug().Str("Chain code generated", code).Msg("")
	return code
}
