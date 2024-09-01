//go:build cli

package cmdgen

import (
	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/rs/zerolog/log"
)

func ConstructCommand(parsedInput *gabs.Container, o *lama2cmd.Opts) ([]string, string) {
	log.Info().Str("ParsedInput", parsedInput.String()).Msg("")
	httpv, url, jsonObj, headers, multipartBool, formBool := ConstructCommandHelper(parsedInput)
	res, stdinBody := assembleCmdString(httpv, url, jsonObj, headers, multipartBool, formBool, o)
	return res, stdinBody
}
