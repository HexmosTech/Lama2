//go:build wasm

package codegen

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"syscall/js"
	"text/template"
	"time"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog/log"
)

//go:embed httpsnippet.js
var snippetcore string

var flag = 0

//go:noinline
func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) string {
	fmt.Println("WASM: Generating target code")
	convertedSnippetFinal := generateConvertedSippet(targetLangLib, parsedAPI)
	fmt.Println("WASM: This is the converted snippet:", convertedSnippetFinal)
	return convertedSnippetFinal
}

//go:noinline
func asyncTask(resultChan0 chan<- string) {
	time.Sleep(2 * time.Second)
	resultChan0 <- "Task Completed"
}

//go:noinline
func generateConvertedSippet(targetLangLib string, parsedAPI *gabs.Container) string {
	fmt.Println("WASM: Generating converted snippet")
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)
	blockLength := len(parsedAPIblocks)
	fmt.Println("WASM: Parsed API", parsedAPIblocks)
	for i, block := range parsedAPIblocks {
		blockType := block.S("type").Data().(string)
		fmt.Println("WASM: Block type:", blockType)
		if blockType == "processor" {
			fmt.Println("WASM: Block type is processor")
			snippet := block.S("value").Data().(*gabs.Container).Data().(string)
			log.Debug().Str("Processor block incoming block", block.String()).Msg("")
			convertedSnippetList = append(convertedSnippetList, snippet)
		} else if blockType == "Lama2File" {
			fmt.Println("WASM: Block type is L2File")
			harRequest, flag := GetRequestHARString(block, targetLangLib)
			fmt.Println("WASM: HAR request:", harRequest)
			snippetArgs := SnippetArgs{}
			lang, lib := SplitLangLib(targetLangLib)
			fmt.Println("WASM: Target lang:", lang)
			snippetArgs.Language = lang
			snippetArgs.Library = lib
			snippetArgs.HARRequest = harRequest
			snippetArgs.SnippetCore = snippetcore
			fmt.Println("WASM: Snippet args:", snippetArgs)
			httpsnippetCode := PrepareHTTPSnippetGenerator(snippetArgs)
			convertedSnippet := js.Global().Call("eval", httpsnippetCode)
			fmt.Println("WASM: convertedSnippet data:", convertedSnippet.String())
			// convertedSnippet, err := evaluateJSCode(httpsnippetCode)
			// if err != nil {
			// 	log.Fatal().
			// 		Str("Type", "CodeGen").
			// 		Str("Error", err.Error()).
			// 		Msg("Code generator error")
			// }
			convertedSnippetWithPostProcessing := postprocessURL(convertedSnippet.String(), flag)
			flag = 0
			convertedSnippetList = append(convertedSnippetList, convertedSnippetWithPostProcessing)
		}
		if i != blockLength-1 {
			convertedSnippetList = append(convertedSnippetList, "\n---\n")
		}
	}

	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	fmt.Println("WASM: Converted snippet final:", convertedSnippetFinal)
	return convertedSnippetFinal
}

// func PrepareHTTPSnippetGenerator(snippetArgs SnippetArgs) string {
// 	var templOutput bytes.Buffer
// 	templStr := `{{.SnippetCore}}

// 	const snippet = new window.HTTPSnippet({{.HARRequest}});

// 	window.convertedSnippet = snippet.convert('{{.Language}}'{{if .Library }}, '{{.Library}}'{{end}});

// 	console.log("convertedSnippet results from httpsnippet:",convertedSnippet)

// 	function evaluateAsync(jsCode, callback) {
//     new Promise((resolve, reject) => {
//         setTimeout(() => {
//             eval(jsCode);
//             resolve(window.convertedSnippet);
//         }, 1000);
//     }).then(result => {
//         callback(result);
//     }).catch(error => {
//         callback('Error: ' + error.message);
//     });
// 	}
// 	`
// 	tmpl, _ := template.New("httpsnippet").Parse(templStr)
// 	tmpl.Execute(&templOutput, snippetArgs)
// 	return templOutput.String()
// }

//go:noinline
func PrepareHTTPSnippetGenerator(snippetArgs SnippetArgs) string {
	fmt.Println("WASM: Snippet args:", snippetArgs)
	var templOutput bytes.Buffer
	templStr := `{{.SnippetCore}} 

	const snippet = new window.HTTPSnippet({{.HARRequest}});
	
	convertedSnippet = snippet.convert('{{.Language}}'{{if .Library }}, '{{.Library}}'{{end}});
	`
	tmpl, _ := template.New("httpsnippet").Parse(templStr)
	tmpl.Execute(&templOutput, snippetArgs)
	fmt.Println("WASM: Templ output:", templOutput.String())
	return templOutput.String()
}
