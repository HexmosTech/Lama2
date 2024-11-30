//go:build wasm

package cmdgen

import (
	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/lama2cmd"
	"fmt"
    "syscall/js"
)


var myPassedValue string
func setMyValue(this js.Value, p []js.Value) interface{} {
    myPassedValue = p[0].String() // Get the first parameter
    fmt.Println("DBGX: Received value from JS in Go:", myPassedValue)
    return nil
}

func ConstructCommand(parsedInput *gabs.Container, o *lama2cmd.Opts) ([]string, string) {
	httpv, url, jsonObj, headers, multipartBool, formBool := ConstructCommandHelper(parsedInput)
	res, stdinBody := assembleCmdString(httpv, url, jsonObj, headers, multipartBool, formBool, nil)
	return res, stdinBody
}
