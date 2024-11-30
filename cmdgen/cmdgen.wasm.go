//go:build wasm

package cmdgen

import (
	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/lama2cmd"
	"syscall/js"
)

func ConstructCommand(parsedInput *gabs.Container, o *lama2cmd.Opts) ([]string, string) {
	httpv, url, jsonObj, headers, multipartBool, formBool := ConstructCommandHelper(parsedInput)
	res, stdinBody := assembleCmdString(httpv, url, jsonObj, headers, multipartBool, formBool, nil)

	LaBearerAuthToken := js.Global().Get("LaBearerAuthToken").String()
	authHeaderExists := false
	for _, header := range res {
		if len(header) >= len("Authorization:") && header[:len("Authorization:")] == "Authorization:" {
			authHeaderExists = true
			break
		}
	}

	if !authHeaderExists {
		res = append(res, "Authorization: Bearer "+LaBearerAuthToken)
	}
	return res, stdinBody
}
