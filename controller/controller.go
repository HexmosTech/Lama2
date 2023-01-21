// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/importer"
	"github.com/HexmosTech/lama2/lama2cmd"
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
	fmt.Println(headerMap)
	if headerMap == nil {
		return
	}
	newHeaderMap := gabs.New()
	for k, v := range headerMap.ChildrenMap() {
		fmt.Println(k, " = ", v)
		key := preprocess.ExpandEnv(k, vm)
		val := preprocess.ExpandEnv(v.Data().(*gabs.Container).Data().(string), vm)
		valWrap := gabs.New()
		valWrap.Set(val)
		newHeaderMap.Set(valWrap, key)
	}
	block.Delete("details", "headers")
	block.Set(newHeaderMap, "details", "headers")
	fmt.Println(block)

}

func expandUrl(block *gabs.Container, vm *goja.Runtime) {
	b := block.S("url", "value").Data().(string)
	fmt.Println(b)
	url := preprocess.ExpandEnv(b, vm)
	block.Delete("url", "value")
	block.Set(url, "url", "value")
}

func expandJSON(block *gabs.Container, vm *goja.Runtime) {
	fmt.Println("In expandJSON", block)
	dataBlock := block.S("details", "ip_data")
	if dataBlock == nil {
		return
	}
	dataBlockStr := dataBlock.String()
	dataBlockStr = preprocess.ExpandEnv(dataBlockStr, vm)
	fmt.Println(dataBlockStr)
	processedBlock, err := gabs.ParseJSON([]byte(dataBlockStr))
	if err != nil {
		fmt.Println("Couldn't expand env variables on top of the JSON. Try TODO option") // TODO
	}
	block.Delete("details", "ip_data")
	block.Set(processedBlock, "details", "ip_data")
	fmt.Println(block)
}

func runVmCode(chainCode string, vm *goja.Runtime) {
	_, err := vm.RunString(chainCode)
	if ex, ok := err.(*goja.Exception); ok {
		fmt.Println(ex.String())
	}
}

func guessRespType(resp string) string {
	match, _ := regexp.MatchString(`^\s*<`, resp)
	if match {
		return "xml"
	}
	return "json"
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
	fmt.Println(code)
	return code
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
			return
		}
		log.Fatal().Msg("To convert Postman export to Lama2, try: l2 -p PostmanFile -l Lama2Dir")
	}
	if len(o.LamaDir) > 0 {
		if len(o.PostmanFile) > 0 {
			importer.PostmanImporter(o.PostmanFile, o.LamaDir)
			return
		}
		log.Fatal().Msg("To convert Postman export to Lama2, try: l2 -p PostmanFile -l Lama2Dir")
	}
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
	fmt.Println("***")
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	fmt.Println(parsedAPIblocks)
	fmt.Println("***")
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	for i, block := range parsedAPIblocks {
		fmt.Println(">> ", i)
		fmt.Println(block)
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			b := block.S("value").Data().(*gabs.Container)
			fmt.Println(b)
			script := b.Data().(string)
			runVmCode(script, vm)
		} else if blockType == "Lama2File" {
			expandUrl(block, vm)
			expandHeaders(block, vm)
			expandJSON(block, vm)
			// TODO - replace stuff in headers, and varjson and json as well
			cmd, stdinBody := cmdgen.ConstructCommand(block, o)
			fmt.Println("@@Body", stdinBody)
			retStr, e1 := cmdexec.ExecCommand(cmd, stdinBody, dir)
			fmt.Println("----------xxxxxxxxx-----------")
			fmt.Println(retStr)
			// parsedOutput, e1 := outputmanager.RequestLogParser(retStr)
			if e1 == nil {
				// fmt.Println("ParsedOutput", parsedOutput)
				// poStr := parsedOutput.S("body").Data().(string)
				// poStr = stripansi.Strip(poStr)
				fmt.Println("----------zzzzzzzzz-----------")
				// fmt.Println(poStr)
				// rType := guessRespType(poStr)
				chainCode := generateChainCode(retStr)
				runVmCode(chainCode, vm)
			} else {
				fmt.Println(e1)
			}
		}
	}
	return

	/*
		apiContent, apiDir := preprocess.LamaFile(o.Positional.LamaAPIFile)
		cmdStr := cmdgen.ConstructCommand(parsedAPI, o)
		log.Info().Msg("COMMAND:\n" + cmdStr)
		op := cmdexec.ExecCommand(cmdStr, apiDir)
		if o.Output != "" {
			outputmanager.WriteJSONOutput(op, o.Output)
		}
	*/
}
