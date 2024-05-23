//go:build wasm

package contoller

import (
	"fmt"
	"os"
	"strings"

	"reflect"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"

	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"

	"syscall/js"
)

func GetParsedAPIBlocks(parsedAPI *gabs.Container) []*gabs.Container {
	return parsedAPI.S("value").Data().(*gabs.Container).Children()
}

func ExecuteProcessorBlock(block *gabs.Container) {
	b := block.S("value").Data().(*gabs.Container)
	script := b.Data().(string)
	js.Global().Call("eval", script)
}

func ExecuteRequestorBlock(block *gabs.Container) httpie.ExResponse {
	preprocess.ProcessVarsInBlock(block)
	// TODO - replace stuff in headers, and varjson and json as well
	cmd, stdinBody := cmdgen.ConstructCommand(block)
	resp, e1 := cmdexec.ExecCommand(cmd, stdinBody)
	printFields(resp)
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
			js.Global().Call("eval", chainCode)
		} else {
			os.Exit(1)
		}
	}

	return resp
}

func printFields(data interface{}) {
	v := reflect.ValueOf(data)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fmt.Printf("The key values available : %s: %v\n", fieldName, fieldValue)
	}
}

func HandleParsedFile(parsedAPI *gabs.Container) httpie.ExResponse {
	parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	var resp httpie.ExResponse
	for i, block := range parsedAPIblocks {
		fmt.Println(i)
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			ExecuteProcessorBlock(block)
		} else if blockType == "Lama2File" {
			resp = ExecuteRequestorBlock(block)
		}
	}
	return resp
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
