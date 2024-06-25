package codegen

import (
	"bytes"
	_ "embed"
	"text/template"
)

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
// func GetHARHeadersCookies(headers *gabs.Container) (*gabs.Container, *gabs.Container) {
// 	headersData := gabs.New()
// 	headersData.Array()

// 	cookiesData := gabs.New()
// 	cookiesData.Array()
// 	for key, val := range headers.Data().(*gabs.Container).ChildrenMap() {
// 		if strings.ToLower(key) == "cookie" {
// 			cookieList := strings.Split(val.Data().(*gabs.Container).Data().(string), ";")
// 			for _, c := range cookieList {
// 				r := strings.Split(c, "=")
// 				ck := r[0]
// 				cv := r[1]
// 				cookie := gabs.New()
// 				cookie.Set(ck, "name")
// 				cookie.Set(cv, "value")
// 				cookiesData.ArrayAppend(cookie)
// 			}
// 			continue
// 		}
// 		header := gabs.New()
// 		header.Set(key, "name")
// 		header.Set(val, "value")
// 		headersData.ArrayAppend(header)
// 	}

// 	cookiesData.Array()
// 	for key, val := range headers.Data().(*gabs.Container).ChildrenMap() {
// 		if strings.ToLower(key) == "cookie" {
// 			cookieList := strings.Split(val.Data().(*gabs.Container).Data().(string), ";")
// 			for _, c := range cookieList {
// 				r := strings.Split(c, "=")
// 				ck := r[0]
// 				cv := r[1]
// 				cookie := gabs.New()
// 				cookie.Set(ck, "name")
// 				cookie.Set(cv, "value")
// 				cookiesData.ArrayAppend(cookie)
// 			}
// 			continue
// 		}
// 		header := gabs.New()
// 		header.Set(key, "name")
// 		header.Set(val, "value")
// 		headersData.ArrayAppend(header)
// 	}
// 	return headersData, cookiesData
// }

// func GetRequestHARString(block *gabs.Container) string {
// 	preprocess.ProcessVarsInBlock(block, globalVM)
// 	httpv := block.S("verb", "value")
// 	url := block.S("url", "value")
// 	jsonObj := block.S("details", "ip_data")
// 	headers := block.S("details", "headers")
// 	/*
// 		TODO: Handle multipart case

// 		multipart := block.S("multipart", "value")
// 		multipartBool := false
// 		if multipart != nil {
// 			multipartBool = true
// 		}
// 	*/
// 	harObj := gabs.New()

// 	if jsonObj != nil {
// 		postData := gabs.New()
// 		postData.Set("application/json", "mimeType")
// 		postData.Set(jsonObj.String(), "text")
// 		harObj.Set(postData, "postData")
// 	}

// 	if headers != nil {
// 		headersData, cookiesData := GetHARHeadersCookies(headers)
// 		if cookiesData.String() != "[]" {
// 			harObj.Set(cookiesData, "cookies")
// 		}
// 		harObj.Set(headersData, "headers")
// 	}

// 	harObj.Set(httpv, "method")
// 	harObj.Set(url, "url")

// 	res := harObj.String()
// 	return res
// }
