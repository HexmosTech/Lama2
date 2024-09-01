package helper

import (
	"os"
	"strconv"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	options "github.com/HexmosTech/lama2/options"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/preprocess"
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

func ExecuteRequestorBlock(block *gabs.Container, vm *goja.Runtime, opts *options.Opts, dir string) (httpie.ExResponse, int64) {
	preprocess.ProcessVarsInBlock(block, vm)
	// TODO - replace stuff in headers, and varjson and json as well
	cmd, stdinBody := cmdgen.ConstructCommand(block, opts)
	log.Debug().Str("Stdin Body to be passed into httpie", stdinBody).Msg("")
	resp, responseTime, e1 := cmdexec.ExecCommand(cmd, stdinBody, dir)
	log.Debug().Str("Response from ExecCommand", resp.Body).Msg("")
	if e1 == nil {
		chainCode := cmdexec.GenerateChainCode(resp.Body)
		cmdexec.RunVMCode(chainCode, vm)
	} else {
		log.Fatal().Str("Error from ExecCommand", e1.Error())
		os.Exit(1)
	}
	return resp, responseTime
}

func CalculateMetrics(resp httpie.ExResponse, timeInMs int64, responseTime []outputmanager.ResponseTime, statusCodes []outputmanager.StatusCode, contentSizes []outputmanager.ContentSize) ([]outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize) {
	responseTime = append(responseTime, outputmanager.ResponseTime{Type: "l2block", TimeInMs: timeInMs})
	statusCodes = append(statusCodes, outputmanager.StatusCode{Type: "l2block", Status: resp.StatusCode})
	sizeInBytes, _ := strconv.Atoi(resp.Headers["Content-Length"])
	contentSizes = append(contentSizes, outputmanager.ContentSize{Type: "l2block", SizeInBytes: sizeInBytes})

	return responseTime, statusCodes, contentSizes
}
