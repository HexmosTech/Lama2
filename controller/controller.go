// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"log"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/dop251/goja"
)

func GetParsedAPIBlocks(parsedAPI *gabs.Container) []*gabs.Container {
	return parsedAPI.S("value").Data().(*gabs.Container).Children()
}

func extractArgs(args []interface{}) (*goja.Runtime, *lama2cmd.Opts, string) {
	var vm *goja.Runtime
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
	return vm, o, dir
}

func processLama2FileBlock(block *gabs.Container, vm *goja.Runtime, o *lama2cmd.Opts, dir string) httpie.ExResponse {
	var resp httpie.ExResponse
	if vm != nil && o != nil && dir != "" {
		resp = ExecuteRequestorBlock(block, vm, o, dir)
	} else {
		resp = ExecuteRequestorBlock(block, vm, o, dir)
	}
	return resp
}

func processBlocks(parsedAPIblocks []*gabs.Container, vm *goja.Runtime, o *lama2cmd.Opts, dir string) (httpie.ExResponse, *lama2cmd.Opts) {
	var resp httpie.ExResponse

	for i, block := range parsedAPIblocks {
		log.Println("Block num", i)
		log.Println("Block getting processed", block.String())
		blockType := block.S("type").Data().(string)
		switch blockType {
		case "processor":
			resp = ExecuteProcessorBlock(block, vm)
		case "Lama2File":
			resp = processLama2FileBlock(block, vm, o, dir)
		}
	}
	return resp, o
}

func HandleParsedFileHelper(parsedAPI *gabs.Container, args ...interface{}) (httpie.ExResponse, *lama2cmd.Opts) {
	parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	vm, o, dir := extractArgs(args)
	return processBlocks(parsedAPIblocks, vm, o, dir)
}

func ExecuteProcessorBlockHelper(block *gabs.Container, vm ...*goja.Runtime) httpie.ExResponse {
	b := block.S("value").Data().(*gabs.Container)
	script := b.Data().(string)

	// Check if vm is provided
	if len(vm) > 0 && vm[0] != nil {
		log.Println("Executing with VM")
		cmdexec.RunVMCode(script, vm[0])
	} else {
		cmdexec.RunVMCode(script, nil)
	}
	return httpie.ExResponse{} // Provide the actual ExResponse value you intend to return
}
