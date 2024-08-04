//go:build wasm

package codegen

import (
	_ "embed"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog/log"
)

var snippetcore string

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) string {
	convertedSnippetFinal := generateConvertedSippet(targetLangLib, parsedAPI)
	fmt.Println("This is the converted snippet:", convertedSnippetFinal)
	return convertedSnippetFinal
}

func generateConvertedSippet(targetLangLib string, parsedAPI *gabs.Container) string {
	fmt.Println("targetLangLib:",targetLangLib)
	// parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)
	fmt.Println("convertedSnippetList:", convertedSnippetList)
	parsedAPIblocks := parsedAPI.S("value").Children()
	fmt.Println("parsedAPIblocks:", parsedAPIblocks)
	done := make(chan struct{})
	worker := js.Global().Get("Worker").New("worker.js")

	worker.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		message := args[0].Get("data")
		if message.Get("error").Truthy() {
			log.Fatal().
				Str("Type", "CodeGen").
				Str("Error", message.Get("error").String()).
				Msg("Code generator error")
		} else {
			convertedSnippet := message.Get("convertedSnippet").String()
			convertedSnippetWithPostProcessing := postprocessURL(convertedSnippet, 0)
			convertedSnippetList = append(convertedSnippetList, convertedSnippetWithPostProcessing)
		}
		done <- struct{}{}
		return nil
	}))

	for i, block := range parsedAPIblocks {
		// log.Debug().Int("Block num", i).Msg("")
		// log.Debug().Str("Block getting processed", block.String()).Msg("")
		fmt.Println("Block type:", block)
		fmt.Println("Block num:", i)
		blockType := block.S("type").Data().(string) 
		if blockType == "processor" {
			snippet := block.S("value").Data().(string)
			log.Debug().Str("Processor block incoming block", block.String()).Msg("")
			convertedSnippetList = append(convertedSnippetList, snippet)
		} else if blockType == "Lama2File" {
			harRequest, _ := GetRequestHARString(block)
			snippetArgs := SnippetArgs{}
			lang, lib := SplitLangLib(targetLangLib)
			snippetArgs.Language = lang
			snippetArgs.Library = lib
			snippetArgs.HARRequest = harRequest
			snippetArgs.SnippetCore = snippetcore
			httpsnippetCode := PrepareHTTPSnippetGenerator(snippetArgs)

			worker.Call("postMessage", httpsnippetCode)
			<-done
		}
	}

	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	return convertedSnippetFinal
}
