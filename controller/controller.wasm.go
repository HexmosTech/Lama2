//go:build wasm

package contoller

import (
	"fmt"
	"os"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/dop251/goja"

	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/codegen"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
)



func HandleParsedFile(parsedAPI *gabs.Container) (httpie.ExResponse, *lama2cmd.Opts) {
	fmt.Println("HandleParsedFile:")
	fmt.Println("HandleParsedFile:", parsedAPI)
	vm := cmdexec.GetJSVm()
	return HandleParsedFileHelper(parsedAPI, vm)
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

func ProcessConverterInput(data string, ConvertLang string) (string, error) {
	apiContent := data
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	fmt.Println(parsedAPI)
	if e != nil {
		fmt.Println("Error while parsing API:", e)
	}
	snippet := codegen.GenerateTargetCode(ConvertLang, parsedAPI)
	fmt.Println("Generated Snippet:", snippet)
	return snippet, nil
}

func ExecuteRequestorBlockHelper(resp httpie.ExResponse, headersString string, e1 error, vm *goja.Runtime) httpie.ExResponse {
	targetHeader := "text/html"
	isTextHTMLPresent := strings.Contains(headersString, targetHeader)
	if isTextHTMLPresent {
		fmt.Printf("'%s' is present in the headers.\n", targetHeader)
		return resp
	} else {
		fmt.Printf("'%s' is not present in the headers.\n", targetHeader)
		if e1 == nil {
			chainCode := cmdexec.GenerateChainCode(resp.Body)
			cmdexec.RunVMCode(chainCode, vm)
		} else {
			fmt.Println("Error from ExecCommand", e1)
			os.Exit(1)
		}
	}
	return resp
}
