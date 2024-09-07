package codegen

import (
	// embed is used to embed the contents of the template files into the binary
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

func GetRequestHARString(block *gabs.Container, targetLang string) (string, int) {
	httpv := block.S("verb", "value").String()
	httpv = strings.Trim(httpv, `"`)
	httpv = strings.Trim(httpv, `'`)
	url := block.S("url", "value")
	flag := preprocessURL(url)
	jsonObj := block.S("details", "ip_data")
	headers := block.S("details", "headers")
	harObj := gabs.New()

	if strings.Contains(strings.ToLower(targetLang), "python") || strings.Contains(strings.ToLower(targetLang), "shell") || strings.Contains(strings.ToLower(targetLang), "php") {
		httpv = strings.ToUpper(httpv)
	}

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
	fmt.Println("HAR request:", res)
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
	fmt.Println("URL:", urls)
	urls = strings.Trim(urls, `"`)
	urls = strings.Trim(urls, `'`)
	fmt.Println("URL Updated:", urls)
	if !strings.HasPrefix(urls, "https://") && !strings.HasPrefix(urls, "http://") {
		fmt.Println("URL does not start with 'https://' or 'http://'")
		newURL := "https://" + urls
		// if !strings.Contains(newURL, ".com") {
		// 	parts := strings.SplitN(newURL, "://", 2)
		// 	protocol = parts[0] + "://"
		// 	if len(parts) == 2 {
		// 		domain := parts[1]
		// 		newURL = parts[0] + "://" + domain + ".com"
		// 	} else {
		// 		newURL += ".com"
		// 	}
		// }
		url.Set(newURL)
		flag = 1
	}

	// Remove the outermost "${}" to isolate the placeholder content
	fmt.Print("URL:", urls)
	return flag
}

func postprocessURL(convertedSnippet string, flag int) string {
	fmt.Println("Converted snippet to postprocess:", convertedSnippet)
	if flag == 1 {
		convertedSnippet = strings.Replace(convertedSnippet, "https://", "", 1)
		convertedSnippet = strings.Replace(convertedSnippet, ".com", "", 1)
		decodedURL, err := url.QueryUnescape(convertedSnippet)
		if err != nil {
			fmt.Println("Error decoding URL:", err)
			return convertedSnippet
		}
		// decodedURL = strings.Replace(decodedURL, `""`, `"`, -1) removed for python
		return decodedURL
		// return convertedSnippet
	}
	return convertedSnippet
}
