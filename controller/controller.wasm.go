//go:build wasm

package contoller

import (
	"fmt"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"

	"github.com/HexmosTech/lama2/parser"
)

func ExecuteProcessorBlock(block *gabs.Container) {
	// b := block.S("value").Data().(*gabs.Container)
	// script := b.Data().(string)
	// js.Global().Call("eval", script)
	ExecuteProcessorBlockHelper(block)
}

func ExecuteRequestorBlock(block *gabs.Container) httpie.ExResponse {
	// preprocess.ProcessVarsInBlock(block)
	// // TODO - replace stuff in headers, and varjson and json as well
	// cmd, stdinBody := cmdgen.ConstructCommand(block)
	// resp, e1 := cmdexec.ExecCommand(cmd, stdinBody)
	// headers := resp.Headers
	// var headersString string
	// for key, value := range headers {
	// 	headersString += fmt.Sprintf("%s: %s\n", key, value)
	// }

	// targetHeader := "text/html"
	// isTextHTMLPresent := strings.Contains(headersString, targetHeader)

	// if isTextHTMLPresent {
	// 	fmt.Printf("'%s' is present in the headers.\n", targetHeader)
	// 	return resp
	// } else {
	// 	fmt.Printf("'%s' is not present in the headers.\n", targetHeader)
	// 	if e1 == nil {
	// 		chainCode := cmdexec.GenerateChainCode(resp.Body)
	// 		js.Global().Call("eval", chainCode)
	// 	} else {
	// 		os.Exit(1)
	// 	}
	// }

	// return resp
	return ExecuteRequestorBlockHelper(block)
}

func HandleParsedFile(parsedAPI *gabs.Container) httpie.ExResponse {
	// parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	// var resp httpie.ExResponse
	// for i, block := range parsedAPIblocks {
	// 	fmt.Println(i)
	// 	blockType := block.S("type").Data().(string)
	// 	if blockType == "processor" {
	// 		ExecuteProcessorBlock(block)
	// 	} else if blockType == "Lama2File" {
	// 		resp = ExecuteRequestorBlock(block)
	// 	}
	// }
	// return resp
	HandleParsedFileHelper(parsedAPI)
}

func ProcessWasmInput(data string) httpie.ExResponse {
	apiContent := data
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	if e != nil {
		fmt.Println("Error while parsing API:", e)
	}

	// Print the parsedAPI value
	fmt.Printf("Parsed API: %+v\n", parsedAPI)
	fmt.Println(parsedAPI)
	return HandleParsedFile(parsedAPI)
}
