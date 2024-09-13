//go:build cli

package codegen

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/atotto/clipboard"
	"github.com/dop251/goja"
	"github.com/rs/zerolog/log"
)

//go:embed httpsnippet.js
var snippetcore string

var globalVM interface{}

func initialize() {
	globalVM = cmdexec.GetJSVm()
}

var flag = 0

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) (string) {
	convertedSnippetFinal := generateConvertedSippet(targetLangLib, parsedAPI)
	clipboard.WriteAll(convertedSnippetFinal)
	log.Info().Msg("Code copied to clipboard")
	return convertedSnippetFinal
}

func PrepareHTTPSnippetGenerator(snippetArgs SnippetArgs) string {
	var templOutput bytes.Buffer
	templStr := `{{.SnippetCore}} 

	const snippet = new window.HTTPSnippet({{.HARRequest}});
	
	convertedSnippet = snippet.convert('{{.Language}}'{{if .Library }}, '{{.Library}}'{{end}});
	`
	tmpl, _ := template.New("httpsnippet").Parse(templStr)
	tmpl.Execute(&templOutput, snippetArgs)
	return templOutput.String()
}

func generateConvertedSippet(targetLangLib string, parsedAPI *gabs.Container) string {
	initialize()
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)
	blockLength := len(parsedAPIblocks)
	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			snippet := block.S("value").Data().(*gabs.Container).Data().(string)
			log.Debug().Str("Processor block incoming block", block.String()).Msg("")
			convertedSnippetList = append(convertedSnippetList, snippet)
		} else if blockType == "Lama2File" {
			preprocess.ProcessVarsInBlock(block, globalVM)
			harRequest, flag := GetRequestHARString(block, targetLangLib)
			snippetArgs := SnippetArgs{}
			lang, lib := SplitLangLib(targetLangLib)
			snippetArgs.Language = lang
			snippetArgs.Library = lib
			snippetArgs.HARRequest = harRequest
			snippetArgs.SnippetCore = snippetcore
			httpsnippetCode := PrepareHTTPSnippetGenerator(snippetArgs)
			vm := cmdexec.GetJSVm()
			_, e := vm.(*goja.Runtime).RunString(httpsnippetCode)
			if e != nil {
				log.Fatal().
					Str("Type", "CodeGen").
					Str("Error", e.Error()).
					Msg("Code generator error")
			}
			convertedSnippet := vm.(*goja.Runtime).Get("convertedSnippet").String()
			convertedSnippetWithPostProcessing := postprocessURL(convertedSnippet, flag)
			flag = 0
			convertedSnippetList = append(convertedSnippetList, convertedSnippetWithPostProcessing)
		}
		if i != blockLength-1 {
			convertedSnippetList = append(convertedSnippetList, "\n---\n")
		}
	}

	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	return convertedSnippetFinal
}
