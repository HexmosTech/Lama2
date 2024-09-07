// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/lama2cmd"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/preprocess"
)

func HandleParsedFileHelper(parsedAPI *gabs.Container, args ...interface{}) (httpie.ExResponse, *lama2cmd.Opts, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize, error) {
    parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
    o, dir := extractArgs(args)
    resp, opts, responseTimes, statusCode, contentSize := processBlocks(parsedAPIblocks, o, dir)
    return resp, opts, responseTimes, statusCode, contentSize, nil
}

func GetParsedAPIBlocks(parsedAPI *gabs.Container) []*gabs.Container {
	return parsedAPI.S("value").Data().(*gabs.Container).Children()
}

func extractArgs(args []interface{}) (*lama2cmd.Opts, string) {
	var o *lama2cmd.Opts
	var dir string

	if len(args) > 0 {
		for _, arg := range args {
			switch v := arg.(type) {
			case *lama2cmd.Opts:
				o = v
			case string:
				dir = v
			}
		}
	}
	return o, dir
}

func processLama2FileBlock(block *gabs.Container, vm interface{}, o *lama2cmd.Opts, dir string) (httpie.ExResponse, int64) {
	// fmt.Println("WW: block value before:", block.String())
	preprocess.ProcessVarsInBlock(block, vm)
	// fmt.Println("WW: block value after:", block.String())

	cmd, stdinBody := cmdgen.ConstructCommand(block, o)
	// fmt.Println("WW: cmd:", cmd)
	// fmt.Println("WW: stdinBody:", stdinBody)
	var resp httpie.ExResponse
	var e1 error
	resp, responseTime, e1 := cmdexec.ExecCommand(cmd, stdinBody, dir)
	headers := resp.Headers
	var headersString string
	for key, value := range headers {
		headersString += fmt.Sprintf("%s: %s\n", key, value)
	}

	resp = ExecuteRequestorBlockHelper(resp, headersString, e1, vm)
	return resp, responseTime
}

// func processBlocks(parsedAPIblocks []*gabs.Container, o *lama2cmd.Opts, dir string) (httpie.ExResponse, *lama2cmd.Opts) {
// 	vm := cmdexec.GetJSVm()
// 	var resp httpie.ExResponse
// 	for i, block := range parsedAPIblocks {
// 		log.Debug().Int("Block num", i).Msg("")
// 		log.Debug().Str("Block getting processed", block.String()).Msg("")
// 		blockType := block.S("type").Data().(string)
// 		switch blockType {
// 		case "processor":
// 			ExecuteProcessorBlock(block, vm)
// 		case "Lama2File":
// 			resp = processLama2FileBlock(block, vm, o, dir)
// 		}
// 	}
// 	return resp, o
// }


