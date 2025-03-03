//go:build wasm

package contoller

import (
	"fmt"
	"os"
	"strings"
	"syscall/js"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/rs/zerolog/log"

	"github.com/HexmosTech/lama2/codegen"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"

	"github.com/HexmosTech/lama2/cmdexec"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	preprocess "github.com/HexmosTech/lama2/preprocess"
)

var worker js.Value

func HandleParsedFile(parsedAPI *gabs.Container) (httpie.ExResponse, *lama2cmd.Opts, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize, error) {
	fmt.Println("WASM: HandleParsedFile:")
	fmt.Println("WASM: HandleParsedFile:", parsedAPI)
	return HandleParsedFileHelper(parsedAPI)
}

func ProcessWasmInput(data string) (httpie.ExResponse, *lama2cmd.Opts, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize, error) {
	fmt.Println("WASM: ProcessWasmInput")
	apiContent := data
	p := parser.NewLama2Parser()
	fmt.Printf("WASM: apicontent %+v\n", apiContent)
	parsedAPI, e := p.Parse(apiContent)
	if e != nil {
		fmt.Println("Error while parsing API:", e)
	}

	// Print the parsedAPI value
	fmt.Printf("WASM: Parsed API: %+v\n", parsedAPI)
	return HandleParsedFile(parsedAPI)
}

func ProcessConverterInput(data string, ConvertLang string) (string, error) {
	fmt.Println("WASM: ProcessConverterInput")
	apiContent := data
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	fmt.Println("WASM: Parsed API:", parsedAPI)
	if e != nil {
		fmt.Println("WASM: Error while parsing API:", e)
	}
	snippet := codegen.GenerateTargetCode(ConvertLang, parsedAPI)
	fmt.Println("WASM: Generated Snippet:", snippet)
	return snippet, nil
}

func ExecuteRequestorBlockHelper(resp httpie.ExResponse, headersString string, e1 error, vm interface{}) httpie.ExResponse {
	fmt.Println("WASM: ExecuteRequestorBlockHelper")
	targetHeader := "text/html"
	isTextHTMLPresent := strings.Contains(headersString, targetHeader)
	if isTextHTMLPresent {
		fmt.Printf("WASM: '%s' is present in the headers.\n", targetHeader)
		return resp
	} else {
		fmt.Printf("WASM: '%s' is not present in the headers.\n", targetHeader)
		if e1 == nil {
			chainCode := cmdexec.GenerateChainCode(resp.Body)
			preprocess.RunCodeInWorker(chainCode)
		} else {
			fmt.Println("WASM: Error from ExecCommand", e1)
			os.Exit(1)
		}
	}
	return resp
}

func processBlocks(parsedAPIblocks []*gabs.Container, o *lama2cmd.Opts, dir string) (httpie.ExResponse, *lama2cmd.Opts, []outputmanager.ResponseTime, []outputmanager.StatusCode, []outputmanager.ContentSize, error) {
	fmt.Println("WASM: processBlocks")
	worker = preprocess.InitWebWorker() // Initialize the web worker
	var resp httpie.ExResponse
	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		switch blockType {
		case "processor":
			ExecuteProcessorBlock(block)
		case "Lama2File":
			resp, _, _ = processLama2FileBlock(block, worker, o, dir)
		}
	}
	fmt.Println("WASM: processBlocks returning")
	return resp, o, nil, nil, nil, nil
}

func ExecuteProcessorBlock(block *gabs.Container) {
	fmt.Println("WASM: ExecuteProcessorBlock")
	b := block.S("value").Data().(*gabs.Container)
	log.Debug().Str("Processor block incoming block", block.String()).Msg("")
	script := b.Data().(string)
	preprocess.RunCodeInWorker(script)
	fmt.Println("WASM: ExecuteProcessorBlock returning")
}
