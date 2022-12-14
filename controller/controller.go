// Package controller coordinates all the other
// components in the `Lama2` project. The high
// level overview of command execution is easily
// understood from this package
package contoller

import (
	"fmt"
	"os"

	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/importer"
	"github.com/HexmosTech/lama2/lama2cmd"
	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/utils"
	"github.com/rs/zerolog/log"
)

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
	apiContent, apiDir := preprocess.LamaFile(o.Positional.LamaAPIFile)
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
	cmdStr := cmdgen.ConstructCommand(parsedAPI, o)
	log.Info().Msg("COMMAND:\n" + cmdStr)
	op := cmdexec.ExecCommand(cmdStr, apiDir)
	if o.Output != "" {
		outputmanager.WriteJSONOutput(op, o.Output)
	}
}
