//go:build wasm

package cmdexec

import (
	"syscall/js"
)

func RunVMCode(jsCode string) {
	js.Global().Call("eval", jsCode)
}

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
