package codegen

import (
	"bytes"
	_ "embed"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"text/template"
	"time"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/dop251/goja"
	"github.com/rs/zerolog/log"
)

//go:embed httpsnippet.js
var snippetcore string

type SnippetArgs struct {
	Language    string
	Library     string
	HARRequest  string
	SnippetCore string
}

var globalVM *goja.Runtime

func initialize() {
	globalVM = cmdexec.GetJSVm()
}

var flag = 0

func GetRequestHARString(block *gabs.Container) (string, int) {
	httpv := block.S("verb", "value")
	url := block.S("url", "value")
	flag := preprocessURL(url)
	jsonObj := block.S("details", "ip_data")
	headers := block.S("details", "headers")
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
	return res, flag
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

func preprocessURL(url *gabs.Container) int {
	urls := url.String()
	flag := 0
	if !strings.HasPrefix(urls, "https://") && !strings.HasPrefix(urls, "http://") {
		http_string := "https://"
		newURL := http_string + urls
		url.Set(newURL)
		flag = 1
		return flag
	}
	// Remove the outermost "${}" to isolate the placeholder content
	return flag
}

func postprocessURL(convertedSnippet string, flag int) string {
	if flag == 1 {
		convertedSnippet = strings.Replace(convertedSnippet, "https://", "", 1)
		decodedURL, err := url.QueryUnescape(convertedSnippet)
		if err != nil {
			fmt.Println("Error decoding URL:", err)
			return convertedSnippet
		}
		decodedURL = strings.Replace(decodedURL, `""`, `"`, -1)
		return decodedURL
		// return convertedSnippet
	}
	return convertedSnippet
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000)
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
			_, e := vm.RunString(httpsnippetCode)
			if e != nil {
				log.Fatal().
					Str("Type", "CodeGen").
					Str("Error", e.Error()).
					Msg("Code generator error")
			}
			convertedSnippet := vm.Get("convertedSnippet").String()
			convertedSnippetWithPostProcessing := postprocessURL(convertedSnippet, flag)
			flag = 0
			convertedSnippetList = append(convertedSnippetList, convertedSnippetWithPostProcessing)
		}
	}

	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	return convertedSnippetFinal
}
