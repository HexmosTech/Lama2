// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
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


func GetParsedAPIBlocks(parsedAPI *gabs.Container) []*gabs.Container {
	return parsedAPI.S("value").Data().(*gabs.Container).Children()
}

func ExecuteProcessorBlock(block *gabs.Container, vm *goja.Runtime) {
	b := block.S("value").Data().(*gabs.Container)
	log.Debug().Str("Processor block incoming block", block.String()).Msg("")
	script := b.Data().(string)
	cmdexec.RunVMCode(script, vm)
}

func ExecuteRequestorBlock(block *gabs.Container, vm *goja.Runtime, opts *lama2cmd.Opts, dir string) (httpie.ExResponse,int64) {
	preprocess.ProcessVarsInBlock(block, vm)
	// TODO - replace stuff in headers, and varjson and json as well
	cmd, stdinBody := cmdgen.ConstructCommand(block, opts)
	log.Debug().Str("Stdin Body to be passed into httpie", stdinBody).Msg("")
	resp,responseTime, e1 := cmdexec.ExecCommand(cmd, stdinBody, dir)
	log.Debug().Str("Response from ExecCommand", resp.Body).Msg("")
	if e1 == nil {
		chainCode := cmdexec.GenerateChainCode(resp.Body)
		cmdexec.RunVMCode(chainCode, vm)
	} else {
		log.Fatal().Str("Error from ExecCommand", e1.Error())
		os.Exit(1)
	}
	return resp,responseTime
}

func HandleParsedFile(parsedAPI *gabs.Container, o *lama2cmd.Opts, dir string) {
	parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	vm := cmdexec.GetJSVm()
	var resp httpie.ExResponse
	var responseTime []outputmanager.ResponseTime
	var statusCodes []outputmanager.StatusCode
	var contentSizes []outputmanager.ContentSize
	var timeInMs int64
	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			ExecuteProcessorBlock(block, vm)
		} else if blockType == "Lama2File" {
			resp, timeInMs = ExecuteRequestorBlock(block, vm, o, dir)
			log.Info().Str("ResponseTime", fmt.Sprintf("%dms", timeInMs)).Msg("")
			responseTime, statusCodes, contentSizes = CalculateMetrics(resp, timeInMs, responseTime, statusCodes, contentSizes)
		}
	}
	if o.Output != "" {
		outputmanager.WriteJSONOutput(resp, responseTime, statusCodes, contentSizes, o.Output)
	}
}

func CalculateMetrics(resp httpie.ExResponse, timeInMs int64, responseTime []outputmanager.ResponseTime, statusCodes []outputmanager.StatusCode, contentSizes []outputmanager.ContentSize) ([]outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize) {
	responseTime = append(responseTime, outputmanager.ResponseTime{Type: "l2block", TimeInMs: timeInMs})
	statusCodes = append(statusCodes, outputmanager.StatusCode{Type: "l2block", Status: resp.StatusCode})
	sizeInBytes, _ := strconv.Atoi(resp.Headers["Content-Length"])
	contentSizes = append(contentSizes, outputmanager.ContentSize{Type: "l2block", SizeInBytes: sizeInBytes})
	
	return responseTime, statusCodes, contentSizes
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
