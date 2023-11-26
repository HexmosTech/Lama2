// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"os"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"

	// "github.com/HexmosTech/lama2/codegen"
	// "github.com/HexmosTech/lama2/lama2cmd"

	// outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"

	// "github.com/HexmosTech/lama2/prettify"
	// "github.com/HexmosTech/lama2/utils"
	// "github.com/rs/zerolog/log"
	"syscall/js"
)

func GetParsedAPIBlocks(parsedAPI *gabs.Container) []*gabs.Container {
	return parsedAPI.S("value").Data().(*gabs.Container).Children()
}

func ExecuteProcessorBlock(block *gabs.Container) {
	b := block.S("value").Data().(*gabs.Container)
	// log.Info().Str("Processor block incoming block", block.String()).Msg("")
	script := b.Data().(string)
	// cmdexec.RunVMCode(script, vm)
	fmt.Println("<-----JS CODE----->:", script)
	// log.Info().Str("Evaluated through syscall js:", script).Msg("")
	js.Global().Call("eval", script)
}

func ExecuteRequestorBlock(block *gabs.Container) httpie.ExResponse {
	preprocess.ProcessVarsInBlock(block)
	// TODO - replace stuff in headers, and varjson and json as well
	cmd, stdinBody := cmdgen.ConstructCommand(block)
	// log.Info().Str("Stdin Body to be passed into httpie", stdinBody).Msg("")
	resp, e1 := cmdexec.ExecCommand(cmd, stdinBody)
	// log.Info().Str("Response from ExecCommand", resp.Body).Msg("")
	if e1 == nil {
		chainCode := cmdexec.GenerateChainCode(resp.Body)
		// log.Info().Str("Evaluated through syscall js:", chainCode).Msg("")
		js.Global().Call("eval", chainCode)
	} else {
		// log.Fatal().Str("Error from ExecCommand", e1.Error())
		os.Exit(1)
	}
	return resp
}

func HandleParsedFile(parsedAPI *gabs.Container) httpie.ExResponse {
	parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	var resp httpie.ExResponse
	for i, block := range parsedAPIblocks {
		// log.Info().Int("Block num", i).Msg("")
		// log.Info().Str("Block getting processed", block.String()).Msg("")
		fmt.Println(i)
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			ExecuteProcessorBlock(block)
		} else if blockType == "Lama2File" {
			resp = ExecuteRequestorBlock(block)
		}
	}
	// if o.Output != "" {
	// 	outputmanager.WriteJSONOutput(resp, o.Output)
	// }
	return resp
}

// Process initiates the following tasks in the given order:
// 1. Parse command line arguments
// 2. Read API file contents
// 3. Expand environment variables in API file
// 4. Parse the API contents
// 5. Generate API request command
// 6. Execute command & retrieve results
// 7. Optionally, post-process and write results to a JSON file
// func Process(version string) {
// 	o := lama2cmd.GetAndValidateCmd(os.Args)
// 	lama2cmd.ArgParsing(o, version)

// 	apiContent := preprocess.GetLamaFileAsString(o.Positional.LamaAPIFile)
// 	_, dir, _ := utils.GetFilePathComponents(o.Positional.LamaAPIFile)
// 	oldDir, _ := os.Getwd()
// 	utils.ChangeWorkingDir(dir)

// 	preprocess.LoadEnvironments(dir)
// 	utils.ChangeWorkingDir(oldDir)
// 	p := parser.NewLama2Parser()
// 	parsedAPI, e := p.Parse(apiContent)
// 	if o.Convert != "" {
// 		codegen.GenerateTargetCode(o.Convert, parsedAPI)
// 		return
// 	}

// 	if o.Prettify {
// 		prettify.Prettify(parsedAPI, p.Context, p.MarkRange, apiContent, o.Positional.LamaAPIFile)
// 		return
// 	}
// 	HandleParsedFile(parsedAPI)
// }

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
	// log.Info().Str("Parsed API", parsedAPI.String()).Msg("")
	return HandleParsedFile(parsedAPI)
}
