//go:build cli

package codegen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/atotto/clipboard"
	"github.com/rs/zerolog/log"
)

//go:embed httpsnippet.js
var snippetcore string

var globalVM *goja.Runtime

func initialize() {
	globalVM = cmdexec.GetJSVm()
}

var flag = 0

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) {
	convertedSnippetFinal := generateConvertedSippet(targetLangLib, parsedAPI)
	fmt.Println("This is the converted snippet:", convertedSnippetFinal)
	clipboard.WriteAll(convertedSnippetFinal)
	fmt.Println("Code copied to clipboard")
}

func generateConvertedSippet(targetLangLib string, parsedAPI *gabs.Container) string {
	initialize()
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)

	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			snippet := block.S("value").Data().(*gabs.Container).Data().(string)
			log.Debug().Str("Processor block incoming block", block.String()).Msg("")
			convertedSnippetList = append(convertedSnippetList, snippet)
		} else if blockType == "Lama2File" {
			harRequest, flag := GetRequestHARString(block)
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
	}

	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	return convertedSnippetFinal
}
