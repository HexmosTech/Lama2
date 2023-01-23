// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"os"
	"path"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/importer"
	"github.com/HexmosTech/lama2/lama2cmd"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/utils"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/rs/zerolog/log"
)

func expandHeaders(block *gabs.Container, vm *goja.Runtime) {
	headerMap := block.S("details", "headers")
	log.Debug().Str("HeaderMap", headerMap.String()).Msg("")
	if headerMap == nil {
		return
	}
	newHeaderMap := gabs.New()
	for k, v := range headerMap.ChildrenMap() {
		log.Trace().Strs("Header pair", []string{k, " = ", v.String()}).Msg("")
		key := preprocess.ExpandEnv(k, vm)
		val := preprocess.ExpandEnv(v.Data().(*gabs.Container).Data().(string), vm)
		valWrap := gabs.New()
		valWrap.Set(val)
		newHeaderMap.Set(valWrap, key)
	}
	block.Delete("details", "headers")
	block.Set(newHeaderMap, "details", "headers")
	log.Debug().Str("Expanded Header block", block.String()).Msg("")
}

func expandURL(block *gabs.Container, vm *goja.Runtime) {
	b := block.S("url", "value").Data().(string)
	log.Debug().Str("Url block", b).Msg("")
	url := preprocess.ExpandEnv(b, vm)
	block.Delete("url", "value")
	block.Set(url, "url", "value")
}

func expandJSON(block *gabs.Container, vm *goja.Runtime) {
	log.Debug().Str("JSON block to be expanded", block.String()).Msg("")
	dataBlock := block.S("details", "ip_data")
	if dataBlock == nil {
		return
	}
	dataBlockStr := dataBlock.String()
	dataBlockStr = preprocess.ExpandEnv(dataBlockStr, vm)
	log.Debug().Str("Expanded JSON data block", dataBlockStr).Msg("")
	processedBlock, err := gabs.ParseJSON([]byte(dataBlockStr))
	if err != nil {
		log.Error().Str("Preprocess JSON block issue", "").Msg("")
		return
	}
	block.Delete("details", "ip_data")
	block.Set(processedBlock, "details", "ip_data")
	log.Debug().Str("Processed JSON block", block.String()).Msg("")
}

func runVMCode(chainCode string, vm *goja.Runtime) {
	_, err := vm.RunString(chainCode)
	if ex, ok := err.(*goja.Exception); ok {
		log.Fatal().Str("Error executing JS processor block", ex.String()).Msg("")
	}
}

func generateChainCode(resp string) string {
	code := `try {
		result = JSON.parse(String.raw` + "`" + resp + "`" + `)	
		console.log("Stored as JSON")
	} catch (e) {
		result = String.raw` + "`" + resp + "`" + ` 	
		console.log(e)
		console.log("Stored as string")
	}`
	log.Debug().Str("Chain code generated", code).Msg("")
	return code
}

func GetParsedAPIBlocks(parsedAPI *gabs.Container) []*gabs.Container {
	return parsedAPI.S("value").Data().(*gabs.Container).Children()
}

func GetJSVm() *goja.Runtime {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	return vm
}

func ExecuteProcessorBlock(block *gabs.Container, vm *goja.Runtime) {
	b := block.S("value").Data().(*gabs.Container)
	log.Debug().Str("Processor block incoming block", block.String()).Msg("")
	script := b.Data().(string)
	runVMCode(script, vm)
}

func ProcessVarsInBlock(block *gabs.Container, vm *goja.Runtime) {
	expandURL(block, vm)
	expandHeaders(block, vm)
	expandJSON(block, vm)
}

func ExecuteRequestorBlock(block *gabs.Container, vm *goja.Runtime, opts *lama2cmd.Opts, dir string) httpie.ExResponse {
	ProcessVarsInBlock(block, vm)
	// TODO - replace stuff in headers, and varjson and json as well
	cmd, stdinBody := cmdgen.ConstructCommand(block, opts)
	log.Debug().Str("Stdin Body to be passed into httpie", stdinBody).Msg("")
	resp, e1 := cmdexec.ExecCommand(cmd, stdinBody, dir)
	log.Debug().Str("Response from ExecCommand", resp.Body).Msg("")
	if e1 == nil {
		chainCode := generateChainCode(resp.Body)
		runVMCode(chainCode, vm)
	} else {
		log.Fatal().Str("Error from ExecCommand", e1.Error())
		os.Exit(1)
	}
	return resp
}

func ArgParsing(o lama2cmd.Opts, version string) {
	if o.Version {
		fmt.Println(version)
		return
	}
	if o.Update {
		utils.UpdateSelf()
		return
	}
	if len(o.PostmanFile) > 0 {
		if len(o.LamaDir) > 0 {
			importer.PostmanImporter(o.PostmanFile, o.LamaDir)
			os.Exit(1)
		}
		log.Fatal().Msg("To convert Postman export to Lama2, try: l2 -p PostmanFile -l Lama2Dir")
	}
	if len(o.LamaDir) > 0 {
		if len(o.PostmanFile) > 0 {
			importer.PostmanImporter(o.PostmanFile, o.LamaDir)
			os.Exit(1)
		}
		log.Fatal().Msg("To convert Postman export to Lama2, try: l2 -p PostmanFile -l Lama2Dir")
	}

}

func HandleParsedFile(parsedAPI *gabs.Container, o *lama2cmd.Opts, dir string) {
	parsedAPIblocks := GetParsedAPIBlocks(parsedAPI)
	vm := GetJSVm()
	var resp httpie.ExResponse
	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			ExecuteProcessorBlock(block, vm)
		} else if blockType == "Lama2File" {
			resp = ExecuteRequestorBlock(block, vm, o, dir)
		}
	}
	if o.Output != "" {
		outputmanager.WriteJSONOutput(resp, o.Output)
	}

}

// Process initiates the following tasks in the given order:
// 1. Parse command line arguments
// 2. Read API file contents
// 3. Expand environment variables in API file
// 4. Parse the API contents
// 5. Generate API request command
// 6. Execute command & retrieve results
// 7. Optionally, post-process and write results to a JSON file
func Process(version string) {
	o := lama2cmd.GetAndValidateCmd(os.Args)
	apiContent := preprocess.GetLamaFileAsString(o.Positional.LamaAPIFile)
	_, dir, _ := utils.GetFilePathComponents(o.Positional.LamaAPIFile)
	preprocess.LoadElfEnv(path.Join(dir, "l2.env"))
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)

	if e != nil {
		log.Fatal().
			Str("Type", "Controller").
			Str("LamaFile", o.Positional.LamaAPIFile).
			Str("Error", e.Error()).
			Msg(fmt.Sprint("Parse Error"))
	}
	log.Debug().Str("Parsed API", parsedAPI.String()).Msg("")
	HandleParsedFile(parsedAPI, o, dir)
}
