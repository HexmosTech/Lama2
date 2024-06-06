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
			case *goja.Runtime:
				vm = v
			case *lama2cmd.Opts:
				o = v
			case string:
				dir = v
			}
		}
	}
	return vm, o, dir
}

// func processProcessorBlock(block *gabs.Container, vm *goja.Runtime) httpie.ExResponse {
// 	if vm != nil {
// 		ExecuteProcessorBlockCli(block, vm)
// 	} else {
// 		ExecuteProcessorBlockWasm(block)
// 	}
// 	return httpie.ExResponse{}
// }

func processLama2FileBlock(block *gabs.Container, vm *goja.Runtime, o *lama2cmd.Opts, dir string) httpie.ExResponse {
	var resp httpie.ExResponse
	if vm != nil && o != nil && dir != "" {
		resp = ExecuteRequestorBlock(block, vm, o, dir)
	} else {
		resp = ExecuteRequestorBlock(block, nil, nil, "")
	}
	return resp
}

func processBlocks(parsedAPIblocks []*gabs.Container, vm *goja.Runtime, o *lama2cmd.Opts, dir string) httpie.ExResponse {
	var resp httpie.ExResponse

	for i, block := range parsedAPIblocks {
		log.Println("Block num", i)
		log.Println("Block getting processed", block.String())
		blockType := block.S("type").Data().(string)
		switch blockType {
		case "processor":
			// resp = processProcessorBlock(block, vm)
			resp = ExecuteProcessorBlock(block, vm)
		case "Lama2File":
			resp = processLama2FileBlock(block, vm, o, dir)
		}
	}

	return resp
}

func HandleParsedFileHelper(parsedAPI *gabs.Container, args ...interface{}) httpie.ExResponse {
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

// func ExecuteRequestorBlockHelper(block *gabs.Container, args ...interface{}) httpie.ExResponse {
// 	var vm *goja.Runtime
// 	var opts *lama2cmd.Opts
// 	var dir string

// 	if len(args) > 0 {
// 		for _, arg := range args {
// 			switch v := arg.(type) {
// 			case *goja.Runtime:
// 				vm = v
// 			case *lama2cmd.Opts:
// 				opts = v
// 			case string:
// 				dir = v
// 			}
// 		}
// 	}

// 	if vm != nil {
// 		preprocess.ProcessVarsInBlock(block, vm)
// 	} else {
// 		preprocess.ProcessVarsInBlock(block)
// 	}

// 	var cmd, stdinBody string
// 	if opts != nil {
// 		cmd, stdinBody = cmdgen.ConstructCommand(block, opts)
// 	} else {
// 		cmd, stdinBody = cmdgen.ConstructCommand(block)
// 	}

// 	var resp httpie.ExResponse
// 	var e1 error
// 	if dir != "" {
// 		resp, e1 = cmdexec.ExecCommand(cmd, stdinBody, dir)
// 	} else {
// 		resp, e1 = cmdexec.ExecCommand(cmd, stdinBody)
// 	}

// 	headers := resp.Headers
// 	var headersString string
// 	for key, value := range headers {
// 		headersString += fmt.Sprintf("%s: %s\n", key, value)
// 	}

// 	targetHeader := "text/html"
// 	isTextHTMLPresent := strings.Contains(headersString, targetHeader)

// 	if isTextHTMLPresent {
// 		fmt.Printf("'%s' is present in the headers.\n", targetHeader)
// 		return resp
// 	} else {
// 		fmt.Printf("'%s' is not present in the headers.\n", targetHeader)
// 		if e1 == nil {
// 			chainCode := cmdexec.GenerateChainCode(resp.Body)
// 			if vm != nil {
// 				cmdexec.RunVMCode(chainCode, vm)
// 			} else {
// 				ExecuteJsCodeWasm(chainCode)
// 				// js.Global().Call("eval", chainCode)
// 			}
// 		} else {
// 			log.Fatal("Error from ExecCommand", e1)
// 			os.Exit(1)
// 		}
// 	}

// 	return resp
// }
