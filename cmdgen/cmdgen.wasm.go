//go:build wasm

package cmdgen

import (
	"github.com/HexmosTech/gabs/v2"
)

func ConstructCommand(parsedInput *gabs.Container) ([]string, string) {
	httpv, url, jsonObj, headers, multipartBool, formBool := ConstructCommandHelper(parsedInput)
	res, stdinBody := assembleCmdString(httpv, url, jsonObj, headers, multipartBool, formBool, nil)
	return res, stdinBody
}
