//go:build cli

// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"os"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/codegen"
	"github.com/HexmosTech/lama2/lama2cmd"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/prettify"
	"github.com/HexmosTech/lama2/utils"
	"github.com/dop251/goja"
	"github.com/rs/zerolog/log"
)

func ExecuteProcessorBlock(block *gabs.Container, vm *goja.Runtime) httpie.ExResponse {
	return ExecuteProcessorBlockHelper(block, vm)
}

// func ExecuteRequestorBlock(block *gabs.Container, vm *goja.Runtime, opts *lama2cmd.Opts, dir string) httpie.ExResponse {
// 	return ExecuteRequestorBlockHelper(block, vm, opts, dir)
// }

func HandleParsedFile(parsedAPI *gabs.Container, o *lama2cmd.Opts, dir string) (httpie.ExResponse, *lama2cmd.Opts) {
	vm := cmdexec.GetJSVm()
	return HandleParsedFileHelper(parsedAPI, o, dir, vm)
}

// Process initiates the following tasks in the given order:
// 1. Parse command line arguments
// 2. Read API file contents
// 3. Expand environment variables in API file
// 4. Parse the API contents
// 5. Generate API request command
// 6. Execute command & retrieve results
// 7. Optionally, post-process and write results to a JSON file
func Process(version string) {
	o := lama2cmd.GetAndValidateCmd(os.Args)
	lama2cmd.ArgParsing(o, version)

	apiContent := preprocess.GetLamaFileAsString(o.Positional.LamaAPIFile)
	_, dir, _ := utils.GetFilePathComponents(o.Positional.LamaAPIFile)
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(dir)

	preprocess.LoadEnvironments(dir)
	utils.ChangeWorkingDir(oldDir)
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	if o.Convert != "" {
		codegen.GenerateTargetCode(o.Convert, parsedAPI)
		return
	}

	if o.Prettify {
		prettify.Prettify(parsedAPI, p.Context, p.MarkRange, apiContent, o.Positional.LamaAPIFile)
		return
	}

	if e != nil {
		log.Fatal().
			Str("Type", "Controller").
			Str("LamaFile", o.Positional.LamaAPIFile).
			Str("Error", e.Error()).
			Msg("Parse Error")
	}
	log.Debug().Str("Parsed API", parsedAPI.String()).Msg("")
	res, out := HandleParsedFile(parsedAPI, o, dir)
	if out.Output != "" {
		outputmanager.WriteJSONOutput(res, out.Output)
	}
}

func ExecuteRequestorBlockHelper(resp httpie.ExResponse, headersString string, e1 error, vm *goja.Runtime) httpie.ExResponse {
	if e1 == nil {
		chainCode := cmdexec.GenerateChainCode(resp.Body)
		cmdexec.RunVMCode(chainCode, vm)
	} else {
		fmt.Printf("Error from ExecCommand", e1)
		os.Exit(1)
	}
	return resp
}
