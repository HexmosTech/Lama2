//go:build cli

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
	"github.com/HexmosTech/lama2/codegen"
	"github.com/HexmosTech/lama2/lama2cmd"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/prettify"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

func HandleParsedFile(parsedAPI *gabs.Container, o *lama2cmd.Opts, dir string) (httpie.ExResponse, *lama2cmd.Opts, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize, error) {
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

	apiContent, err := preprocess.GetLamaFileAsString(o.Positional.LamaAPIFile)
	if err != nil {
		log.Fatal().Str("Type", "Preprocess").Msg(fmt.Sprint("Couldn't read: ", o.Positional.LamaAPIFile))
	}
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
	res, out, responseTime, statusCodes, contentSizes, err := HandleParsedFile(parsedAPI, o, dir)
	if err != nil {
		log.Fatal().Str("Type", "Controller").Msg(fmt.Sprint("Error: ", err))
	}
	if out.Output != "" {
		outputmanager.WriteJSONOutput(res, out.Output, responseTime, statusCodes, contentSizes)
	}
}

func processBlocks(parsedAPIblocks []*gabs.Container, o *lama2cmd.Opts, dir string) (httpie.ExResponse, *lama2cmd.Opts, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize) {
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
		switch blockType {
		case "processor":
			ExecuteProcessorBlock(block, vm)
		case "Lama2File":
			resp, timeInMs = processLama2FileBlock(block, vm, o, dir)
			log.Info().Str("ResponseTime", fmt.Sprintf("%dms", timeInMs)).Msg("")
			responseTime, statusCodes, contentSizes = CalculateMetrics(resp, timeInMs, responseTime, statusCodes, contentSizes)
		}
	}
	return resp, o, responseTime, statusCodes, contentSizes
}

func ExecuteRequestorBlockHelper(resp httpie.ExResponse, headersString string, e1 error, vm interface{}) httpie.ExResponse {
	if e1 == nil {
		chainCode := cmdexec.GenerateChainCode(resp.Body)
		cmdexec.RunVMCode(chainCode, vm)
	} else {
		fmt.Printf("Error from ExecCommand", e1)
		os.Exit(1)
	}
	return resp
}

func ExecuteProcessorBlock(block *gabs.Container, vm interface{}) {
	b := block.S("value").Data().(*gabs.Container)
	log.Debug().Str("Processor block incoming block", block.String()).Msg("")
	script := b.Data().(string)
	cmdexec.RunVMCode(script, vm)
}

func CalculateMetrics(resp httpie.ExResponse, timeInMs int64, responseTime []outputmanager.ResponseTime, statusCodes []outputmanager.StatusCode, contentSizes []outputmanager.ContentSize) ([]outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize) {
	responseTime = append(responseTime, outputmanager.ResponseTime{Type: "l2block", TimeInMs: timeInMs})
	statusCodes = append(statusCodes, outputmanager.StatusCode{Type: "l2block", Status: resp.StatusCode})
	sizeInBytes, _ := strconv.Atoi(resp.Headers["Content-Length"])
	contentSizes = append(contentSizes, outputmanager.ContentSize{Type: "l2block", SizeInBytes: sizeInBytes})

	return responseTime, statusCodes, contentSizes
}
