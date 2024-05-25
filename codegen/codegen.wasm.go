//go:build wasm

package codegen

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"syscall/js"
	"text/template"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/preprocess"
)

var snippetcore string

type SnippetArgs struct {
	Language    string
	Library     string
	HARRequest  string
	SnippetCore string
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
	parsedAPIblocks := parsedAPI.S("value").Data().(*gabs.Container).Children()
	convertedSnippetList := make([]string, 0)

	for i, block := range parsedAPIblocks {
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
			js.Global().Call("eval", httpsnippetCode)
			convertedSnippet := js.Global().Get("convertedSnippet").String()
			convertedSnippetList = append(convertedSnippetList, convertedSnippet)
		}
	}
	convertedSnippetFinal := strings.Join(convertedSnippetList, "\n")
	fmt.Println(convertedSnippetFinal)
}
