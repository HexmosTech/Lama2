//go:build wasm

package contoller

import (
	"fmt"
	"os"
	"strings"
	"syscall/js"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/dop251/goja"

	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
)

func ExecuteProcessorBlock(block *gabs.Container, vm *goja.Runtime) httpie.ExResponse {
	return ExecuteProcessorBlockHelper(block)
}

func ExecuteJsCodeWasm(script string) {
	js.Global().Call("eval", script)
}

func ExecuteRequestorBlock(block *gabs.Container, vm *goja.Runtime, opts *lama2cmd.Opts, dir string) httpie.ExResponse {
	return ExecuteRequestorBlockHelper(block)
}

func HandleParsedFile(parsedAPI *gabs.Container) (httpie.ExResponse, *lama2cmd.Opts) {
	return HandleParsedFileHelper(parsedAPI)
}

func ProcessWasmInput(data string) (httpie.ExResponse, *lama2cmd.Opts) {
	apiContent := data
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	if e != nil {
		fmt.Println("Error while parsing API:", e)
	}

	// Print the parsedAPI value
	fmt.Printf("Parsed API: %+v\n", parsedAPI)
	return HandleParsedFile(parsedAPI)
}

func ExecuteRequestorBlockHelper(block *gabs.Container, args ...interface{}) httpie.ExResponse {
	var vm *goja.Runtime
	var opts *lama2cmd.Opts
	var dir string

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
	preprocess.ProcessVarsInBlock(block, vm)
	var cmd []string
	var stdinBody string
	cmd, stdinBody = cmdgen.ConstructCommand(block, opts)
	var resp httpie.ExResponse
	var e1 error
	resp, e1 = cmdexec.ExecCommand(cmd, stdinBody, dir)
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
				ExecuteJsCodeWasm(chainCode)
			}
		} else {
			fmt.Println("Error from ExecCommand", e1)
			os.Exit(1)
		}
	}

	return resp
}
