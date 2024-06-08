//go:build cli

// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"os"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/codegen"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/prettify"
	"github.com/HexmosTech/lama2/utils"
	"github.com/dop251/goja"
	"github.com/rs/zerolog/log"
)

func ExecuteProcessorBlock(block *gabs.Container, vm *goja.Runtime) httpie.ExResponse {
	// b := block.S("value").Data().(*gabs.Container)
	// log.Debug().Str("Processor block incoming block", block.String()).Msg("")
	// script := b.Data().(string)
	// cmdexec.RunVMCode(script, vm)
	return ExecuteProcessorBlockHelper(block, vm)
}

func ExecuteRequestorBlock(block *gabs.Container, vm *goja.Runtime, opts *lama2cmd.Opts, dir string) httpie.ExResponse {
	// preprocess.ProcessVarsInBlock(block, vm)
	// // TODO - replace stuff in headers, and varjson and json as well
	// cmd, stdinBody := cmdgen.ConstructCommand(block, opts)
	// log.Debug().Str("Stdin Body to be passed into httpie", stdinBody).Msg("")
	// resp, e1 := cmdexec.ExecCommand(cmd, stdinBody, dir)
	// log.Debug().Str("Response from ExecCommand", resp.Body).Msg("")
	// if e1 == nil {
	// 	chainCode := cmdexec.GenerateChainCode(resp.Body)
	// 	cmdexec.RunVMCode(chainCode, vm)
	// } else {
	// 	log.Fatal().Str("Error from ExecCommand", e1.Error())
	// 	os.Exit(1)
	// }
	// return resp
	return ExecuteRequestorBlockHelper(block, vm, opts, dir)
}

func HandleParsedFile(parsedAPI *gabs.Container, o *lama2cmd.Opts, dir string) httpie.ExResponse {
	// parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	// vm := cmdexec.GetJSVm()
	// var resp httpie.ExResponse
	// for i, block := range parsedAPIblocks {
	// 	log.Debug().Int("Block num", i).Msg("")
	// 	log.Debug().Str("Block getting processed", block.String()).Msg("")
	// 	blockType := block.S("type").Data().(string)
	// 	if blockType == "processor" {
	// 		ExecuteProcessorBlock(block, vm)
	// 	} else if blockType == "Lama2File" {
	// 		resp = ExecuteRequestorBlock(block, vm, o, dir)
	// 	}
	// }
	// if o.Output != "" {
	// 	outputmanager.WriteJSONOutput(resp, o.Output)
	// }
	return HandleParsedFileHelper(parsedAPI, o, dir)
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
	HandleParsedFile(parsedAPI, o, dir)
}

func ExecuteRequestorBlockHelper(block *gabs.Container, args ...interface{}) httpie.ExResponse {
	var vm *goja.Runtime
	var opts *lama2cmd.Opts
	var dir string
	var cmd []string
	var stdinBody string
	if len(args) > 0 {
		for _, arg := range args {
			switch v := arg.(type) {
			case *goja.Runtime:
				vm = v
			case *lama2cmd.Opts:
				opts = v
			case string:
				dir = v
			}
		}
	}

	if vm != nil {
		preprocess.ProcessVarsInBlock(block, vm)
	} else {
		preprocess.ProcessVarsInBlock(block, nil)
	}

	if opts != nil {
		cmd, stdinBody = cmdgen.ConstructCommand(block, opts)
	} else {
		cmd, stdinBody = cmdgen.ConstructCommand(block, opts)
	}

	var resp httpie.ExResponse
	var e1 error
	if dir != "" {
		resp, e1 = cmdexec.ExecCommand(cmd, stdinBody, dir)
	} else {
		resp, e1 = cmdexec.ExecCommand(cmd, stdinBody, "")
	}

	headers := resp.Headers
	var headersString string
	for key, value := range headers {
		headersString += fmt.Sprintf("%s: %s\n", key, value)
	}

	targetHeader := "text/html"
	isTextHTMLPresent := strings.Contains(headersString, targetHeader)

	if isTextHTMLPresent {
		fmt.Printf("'%s' is present in the headers.\n", targetHeader)
		return resp
	} else {
		fmt.Printf("'%s' is not present in the headers.\n", targetHeader)
		if e1 == nil {
			chainCode := cmdexec.GenerateChainCode(resp.Body)
			if vm != nil {
				cmdexec.RunVMCode(chainCode, vm)
			} else {
				cmdexec.RunVMCode(chainCode, vm)
				// ExecuteJsCodeWasm(chainCode)
				// js.Global().Call("eval", chainCode)
			}
		} else {
			fmt.Printf("Error from ExecCommand", e1)
			os.Exit(1)
		}
	}

	return resp
}
