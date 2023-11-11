package codegen

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
	"syscall/js"
	"github.com/rs/zerolog/log"

	"github.com/HexmosTech/gabs/v2"
	// "github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/preprocess"
)

//go:embed httpsnippet.js
var snippetcore string

type SnippetArgs struct {
	Language    string
	Library     string
	HARRequest  string
	SnippetCore string
}

// var globalVM *goja.Runtime

// func initialize() {
// 	globalVM = cmdexec.GetJSVm()
// }

func PrepareHTTPSnippetGenerator(snippetArgs SnippetArgs) string {
	var templOutput bytes.Buffer
	templStr := `{{.SnippetCore}} 

	const snippet = new window.HTTPSnippet({{.HARRequest}});
	
	let convertedSnippet = snippet.convert('{{.Language}}'{{if .Library }}, '{{.Library}}'{{end}});
	`
	tmpl, _ := template.New("httpsnippet").Parse(templStr)
	tmpl.Execute(&templOutput, snippetArgs)
	return templOutput.String()
}

// takes in the headers in L2 format, and generates
// HAR compatible
func GetHARHeadersCookies(headers *gabs.Container) (*gabs.Container, *gabs.Container) {
	headersData := gabs.New()
	headersData.Array()

	cookiesData := gabs.New()
	cookiesData.Array()
	for key, val := range headers.Data().(*gabs.Container).ChildrenMap() {
		if strings.ToLower(key) == "cookie" {
			cookieList := strings.Split(val.Data().(*gabs.Container).Data().(string), ";")
			for _, c := range cookieList {
				r := strings.Split(c, "=")
				ck := r[0]
				cv := r[1]
				cookie := gabs.New()
				cookie.Set(ck, "name")
				cookie.Set(cv, "value")
				cookiesData.ArrayAppend(cookie)
			}
			continue
		}
		header := gabs.New()
		header.Set(key, "name")
		header.Set(val, "value")
		headersData.ArrayAppend(header)
	}

	cookiesData.Array()
	for key, val := range headers.Data().(*gabs.Container).ChildrenMap() {
		if strings.ToLower(key) == "cookie" {
			cookieList := strings.Split(val.Data().(*gabs.Container).Data().(string), ";")
			for _, c := range cookieList {
				r := strings.Split(c, "=")
				ck := r[0]
				cv := r[1]
				cookie := gabs.New()
				cookie.Set(ck, "name")
				cookie.Set(cv, "value")
				cookiesData.ArrayAppend(cookie)
			}
			continue
		}
		header := gabs.New()
		header.Set(key, "name")
		header.Set(val, "value")
		headersData.ArrayAppend(header)
	}
	return headersData, cookiesData
}

func GetRequestHARString(block *gabs.Container) string {
	preprocess.ProcessVarsInBlock(block)
	httpv := block.S("verb", "value")
	url := block.S("url", "value")
	jsonObj := block.S("details", "ip_data")
	headers := block.S("details", "headers")
	/*
		TODO: Handle multipart case

		multipart := block.S("multipart", "value")
		multipartBool := false
		if multipart != nil {
			multipartBool = true
		}
	*/
	harObj := gabs.New()

	if jsonObj != nil {
		postData := gabs.New()
		postData.Set("application/json", "mimeType")
		postData.Set(jsonObj.String(), "text")
		harObj.Set(postData, "postData")
	}

	if headers != nil {
		headersData, cookiesData := GetHARHeadersCookies(headers)
		if cookiesData.String() != "[]" {
			harObj.Set(cookiesData, "cookies")
		}
		harObj.Set(headersData, "headers")
	}

	harObj.Set(httpv, "method")
	harObj.Set(url, "url")

	res := harObj.String()
	return res
}

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) {
	// initialize()
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)

	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			snippet := block.S("value").Data().(*gabs.Container).Data().(string)
			convertedSnippetList = append(convertedSnippetList, snippet)
		} else if blockType == "Lama2File" {
			harRequest := GetRequestHARString(block)
			snippetArgs := SnippetArgs{}
			lang, lib := SplitLangLib(targetLangLib)
			snippetArgs.Language = lang
			snippetArgs.Library = lib
			snippetArgs.HARRequest = harRequest
			snippetArgs.SnippetCore = snippetcore
			httpsnippetCode := PrepareHTTPSnippetGenerator(snippetArgs)

			// vm := cmdexec.GetJSVm()
			// _, e := vm.RunString(httpsnippetCode)
			js.Global().Call("eval", httpsnippetCode) 

			// if e != nil {
			// 	log.Fatal().
			// 		Str("Type", "CodeGen").
			// 		Str("Error", e.Error()).
			// 		Msg("Code generator error")
			// }
			// Init returns an error if the package is not ready for use.
			// convertedSnippet := vm.Get("convertedSnippet").String()

			convertedSnippet :=  js.Global().Get("convertedSnippet").String()
			convertedSnippetList = append(convertedSnippetList, convertedSnippet)
		}
	}
	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	fmt.Println(convertedSnippetFinal)
}
