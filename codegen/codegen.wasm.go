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

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) string {
	convertedSnippetFinal := generateConvertedSippet(targetLangLib, parsedAPI)
	fmt.Println("This is the converted snippet:", convertedSnippetFinal)
	return convertedSnippetFinal
}

func asyncTask(resultChan0 chan<- string) {
	time.Sleep(2 * time.Second)
	resultChan0 <- "Task Completed"
}

func evaluateJSCode(jsCode string) (string, error) {
	global := js.Global()
	eval := global.Get("eval")
	eval.Invoke(jsCode)

	resultChan0 := make(chan string)
	go asyncTask(resultChan0)
	results := <-resultChan0
	fmt.Println(results)

	evalAsync := global.Get("evaluateAsync")
	if !evalAsync.Truthy() {
		return "", fmt.Errorf("evaluateAsync function not found")
	}
	resultChan := make(chan string)

	callback := js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		resultChan <- p[0].String()
		return nil
	})
	evalAsync.Invoke(jsCode, callback)
	result := <-resultChan
	return result, nil
}

func generateConvertedSippet(targetLangLib string, parsedAPI *gabs.Container) string {
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)

	for _, block := range parsedAPIblocks {
		fmt.Println("Block type:", block.S("type").Data().(string))
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
			convertedSnippet, err := evaluateJSCode(httpsnippetCode)
			if err != nil {
				log.Fatal().
					Str("Type", "CodeGen").
					Str("Error", err.Error()).
					Msg("Code generator error")
			}
			convertedSnippetWithPostProcessing := postprocessURL(convertedSnippet, flag)
			flag = 0
			convertedSnippetList = append(convertedSnippetList, convertedSnippetWithPostProcessing)
		}
	}
	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	return convertedSnippetFinal
}

func PrepareHTTPSnippetGenerator(snippetArgs SnippetArgs) string {
	var templOutput bytes.Buffer
	templStr := `{{.SnippetCore}} 

	const snippet = new window.HTTPSnippet({{.HARRequest}});
	
	window.convertedSnippet = snippet.convert('{{.Language}}'{{if .Library }}, '{{.Library}}'{{end}});
	
	console.log("convertedSnippet results from httpsnippet:",convertedSnippet)

	function evaluateAsync(jsCode, callback) {
    new Promise((resolve, reject) => {
        setTimeout(() => {
            eval(jsCode);
            resolve(window.convertedSnippet);
        }, 1000); 
    }).then(result => {
        callback(result);
    }).catch(error => {
        callback('Error: ' + error.message);
    });
	}
	`
	tmpl, _ := template.New("httpsnippet").Parse(templStr)
	tmpl.Execute(&templOutput, snippetArgs)
	return templOutput.String()
}
