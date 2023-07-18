package codegen

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/atotto/clipboard"
)

//go:embed httpsnippet.js
var snippetcore string

type SnippetArgs struct {
	Language    string
	Library     string
	HARRequest  string
	SnippetCore string
}

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

func GetRequestHARString(parsedAPI *gabs.Container) string {
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	vm := cmdexec.GetJSVm()
	for i, block := range parsedAPIblocks {
		log.Debug().Int("Block num", i).Msg("")
		log.Debug().Str("Block getting processed", block.String()).Msg("")
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			fmt.Println("Skipping processor block")
		} else if blockType == "Lama2File" {
			preprocess.ProcessVarsInBlock(block, vm)
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
	}
	// TODO: Handle multi-stage files properly
	// Better exception handling
	return ""
}

func GenerateTargetCode(targetLangLib string, parsedAPI *gabs.Container) {
	harRequest := GetRequestHARString(parsedAPI)
	snippetArgs := SnippetArgs{}
	lang, lib := SplitLangLib(targetLangLib)
	snippetArgs.Language = lang
	snippetArgs.Library = lib
	snippetArgs.HARRequest = harRequest
	snippetArgs.SnippetCore = snippetcore
	httpsnippetCode := PrepareHTTPSnippetGenerator(snippetArgs)
	vm := cmdexec.GetJSVm()
	_, e := vm.RunString(httpsnippetCode)
	if e != nil {
		log.Fatal().
			Str("Type", "CodeGen").
			Str("Error", e.Error()).
			Msg(fmt.Sprint("Code generator error"))
	}
	// Init returns an error if the package is not ready for use.
	convertedSnippet := vm.Get("convertedSnippet").String()
	fmt.Println(convertedSnippet)
	clipboard.WriteAll(convertedSnippet)
	fmt.Println("Code copied to clipboard")
}
