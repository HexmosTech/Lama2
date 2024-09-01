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
	"github.com/HexmosTech/lama2/argparse"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/codegen"
	helpers "github.com/HexmosTech/lama2/helpers"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/options"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/prettify"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

func HandleParsedFile(parsedAPI *gabs.Container, o *options.Opts, dir string) {
	parsedAPIblocks := helpers.GetParsedAPIBlocks(parsedAPI)
	vm := cmdexec.GetJSVm()
	var resp httpie.ExResponse
	var responseTime []outputmanager.ResponseTime
	var statusCodes []outputmanager.StatusCode
	var contentSizes []outputmanager.ContentSize
	var timeInMs int64
	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			helpers.ExecuteProcessorBlock(block, vm)
		} else if blockType == "Lama2File" {
			resp, timeInMs = helpers.ExecuteRequestorBlock(block, vm, o, dir)
			log.Info().Str("ResponseTime", fmt.Sprintf("%dms", timeInMs)).Msg("")
			responseTime, statusCodes, contentSizes = helpers.CalculateMetrics(resp, timeInMs, responseTime, statusCodes, contentSizes)
		}
	}
	if o.Output != "" {
		outputmanager.WriteJSONOutput(resp, responseTime, statusCodes, contentSizes, o.Output)
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
	argparse.ArgParsing(o, version)

	apiContent := preprocess.GetLamaFileAsString(o.Positional.LamaAPIFile)
	_, dir, _ := utils.GetFilePathComponents(o.Positional.LamaAPIFile)
	oldDir, _ := os.Getwd()
	utils.ChangeWorkingDir(dir)

	preprocess.LoadEnvironments(dir)
	utils.ChangeWorkingDir(oldDir)
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	if o.Convert != "" {
		codegen.GenerateTargetCode(o.Convert, parsedAPI)
		return
	}

	if o.Prettify {
		prettify.Prettify(parsedAPI, p.Context, p.MarkRange, apiContent, o.Positional.LamaAPIFile)
		return
	}

	if e != nil {
		log.Fatal().
			Str("Type", "Controller").
			Str("LamaFile", o.Positional.LamaAPIFile).
			Str("Error", e.Error()).
			Msg("Parse Error")
	}
	log.Debug().Str("Parsed API", parsedAPI.String()).Msg("")
	HandleParsedFile(parsedAPI, o, dir)
}
