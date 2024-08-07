package codegen

import (
	_ "embed"
	"fmt"
	"net/url"
	"strings"

	"github.com/HexmosTech/gabs/v2"
)

type SnippetArgs struct {
	Language    string
	Library     string
	HARRequest  string
	SnippetCore string
}

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
