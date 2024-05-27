// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall/js"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/dop251/goja"
)

func GetParsedAPIBlocks(parsedAPI *gabs.Container) []*gabs.Container {
	return parsedAPI.S("value").Data().(*gabs.Container).Children()
}

func HandleParsedFileHelper(parsedAPI *gabs.Container, args ...interface{}) ExResponse {
	parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	var resp ExResponse
	var vm *goja.Runtime
	var o *Opts
	var dir string

	if len(args) > 0 {
		for _, arg := range args {
			switch v := arg.(type) {
			case *goja.Runtime:
				vm = v
			case *Opts:
				o = v
			case string:
				dir = v
			}
		}
	}

	for i, block := range parsedAPIblocks {
		log.Println("Block num", i)
		log.Println("Block getting processed", block.String())
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			if vm != nil {
				ExecuteProcessorBlock(block, vm)
			} else {
				ExecuteProcessorBlock(block)
			}
		} else if blockType == "Lama2File" {
			if vm != nil && o != nil && dir != "" {
				resp = ExecuteRequestorBlock(block, vm, o, dir)
			} else {
				resp = ExecuteRequestorBlock(block)
			}
		}
	}

	if o != nil && o.Output != "" {
		outputmanager{}.WriteJSONOutput(resp, o.Output)
	}

	return resp
}

func ExecuteProcessorBlockHelper(block *gabs.Container, vm ...*goja.Runtime) {
	b := block.S("value").Data().(*gabs.Container)
	script := b.Data().(string)

	// Check if vm is provided
	if len(vm) > 0 && vm[0] != nil {
		log.Println("Executing with VM")
		cmdexec.RunVMCode(script, vm[0])
	} else {
		log.Println("Executing without VM")
		js.Global().Call("eval", script)
	}
}

func ExecuteRequestorBlockHelper(block *gabs.Container, args ...interface{}) httpie.ExResponse {
	var vm *goja.Runtime
	var opts *Opts
	var dir string

	if len(args) > 0 {
		for _, arg := range args {
			switch v := arg.(type) {
			case *goja.Runtime:
				vm = v
			case *Opts:
				opts = v
			case string:
				dir = v
			}
		}
	}

	if vm != nil {
		preprocess.ProcessVarsInBlock(block, vm)
	} else {
		preprocess.ProcessVarsInBlock(block)
	}

	var cmd, stdinBody string
	if opts != nil {
		cmd, stdinBody = cmdgen.ConstructCommand(block, opts)
	} else {
		cmd, stdinBody = cmdgen.ConstructCommand(block)
	}

	var resp ExResponse
	var e1 error
	if dir != "" {
		resp, e1 = cmdexec.ExecCommand(cmd, stdinBody, dir)
	} else {
		resp, e1 = cmdexec.ExecCommand(cmd, stdinBody)
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
				js.Global().Call("eval", chainCode)
			}
		} else {
			log.Fatal("Error from ExecCommand", e1)
			os.Exit(1)
		}
	}

	return resp
}
