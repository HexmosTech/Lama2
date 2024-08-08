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

// func evaluateJSCode(jsCode string) (string, error) {
// 	global := js.Global()
// 	eval := global.Get("eval")
// 	eval.Invoke(jsCode)

// 	resultChan0 := make(chan string,10)
// 	go asyncTask(resultChan0)
// 	results := <-resultChan0
// 	fmt.Println(results)

// 	evalAsync := global.Get("evaluateAsync")
// 	if !evalAsync.Truthy() {
// 		return "", fmt.Errorf("evaluateAsync function not found")
// 	}
// 	resultChan := make(chan string,10)

// 	callback := js.FuncOf(func(this js.Value, p []js.Value) interface{} {
// 		resultChan <- p[0].String()
// 		return nil
// 	})
// 	evalAsync.Invoke(jsCode, callback)
// 	result := <-resultChan
// 	return result, nil
// }

// func evaluateJSCode (code string) (string, error) {
// 	// Create a new JavaScript function to run the code
// 	result := js.Global().Call("eval", code)
// 	// if result.Type() == js.TypeString {
// 	// 	return result.String(), nil
// 	// }
// 	return result.String(), nil
// }

func generateConvertedSippet(targetLangLib string, parsedAPI *gabs.Container) string {
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)
	blockLength := len(parsedAPIblocks)
	fmt.Println("Parsed API", parsedAPIblocks)
	for i, block := range parsedAPIblocks {
		blockType := block.S("type").Data().(string)
		fmt.Println("Block type:", blockType)
		if blockType == "processor" {
			fmt.Println("Block type is processor")
			snippet := block.S("value").Data().(*gabs.Container).Data().(string)
			log.Debug().Str("Processor block incoming block", block.String()).Msg("")
			convertedSnippetList = append(convertedSnippetList, snippet)
		} else if blockType == "Lama2File" {
			fmt.Println("Block type is L2File")
			harRequest, flag := GetRequestHARString(block)
			fmt.Println("HAR request:", harRequest)
			snippetArgs := SnippetArgs{}
			lang, lib := SplitLangLib(targetLangLib)
			fmt.Println("Target lang:", lang)
			snippetArgs.Language = lang
			snippetArgs.Library = lib
			snippetArgs.HARRequest = harRequest
			snippetArgs.SnippetCore = snippetcore
			fmt.Println("Snippet args:", snippetArgs)
			httpsnippetCode := PrepareHTTPSnippetGenerator(snippetArgs)
			fmt.Println("httpsnippetCode:", httpsnippetCode)
			convertedSnippet := js.Global().Call("eval", httpsnippetCode)
			fmt.Println("convertedSnippet data:", convertedSnippet.String())
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
