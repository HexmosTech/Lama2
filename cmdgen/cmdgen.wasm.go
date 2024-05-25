//go:build wasm

package cmdgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/HexmosTech/gabs/v2"
)

func assembleCmdString(httpv string, url string, jsonObj *gabs.Container, headers *gabs.Container, multipart bool, form bool) ([]string, string) {
	command := make([]string, 0)
	var files *gabs.Container
	if multipart {
		if jsonObj.ExistsP("@files") {
			files = jsonObj.S("@files")
			jsonObj.Delete("@files")
		}
	}

	jsonStr := ""
	if jsonObj != nil && !multipart && !form {
		dst := &bytes.Buffer{}
		if err := json.Compact(dst, []byte(jsonObj.String())); err != nil {
		}
		jsonStr = dst.String()
	}

	command = append(command, "ht ")
	if multipart || form {
		command = append(command, "--ignore-stdin", "--form")
	}

	command = append(command, httpv+" ")
	command = append(command, url+" ")

	if multipart {
		for key, val := range jsonObj.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, "'"+key+"'='"+val.Data().(string)+"'  ")
		}
		for key, val := range files.ChildrenMap() {
			command = append(command, key+"@"+val.Data().(string))
		}
	}

	if form {
		for key, val := range jsonObj.Data().(*gabs.Container).ChildrenMap() {
			keyValuePair := fmt.Sprintf("'%s'='%s'  ", key, val.Data().(string))
			command = append(command, keyValuePair)
		}
	}

	if headers != nil {
		for key, val := range headers.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, key+":"+val.Data().(*gabs.Container).Data().(string))
		}
	}
	cleanCommand := make([]string, 0)
	for _, c := range command {
		cleanC := strings.TrimSpace(c)
		cleanCommand = append(cleanCommand, cleanC)
	}
	if multipart || form {
		return cleanCommand, ""
	}
	return cleanCommand, jsonStr
}

// func assembleCmdString(httpv string, url string, jsonObj *gabs.Container, headers *gabs.Container, multipart bool, form bool, o *lama2cmd.Opts) ([]string, string) {
// 	return assembleCmdStringHelper(httpv, url, jsonObj, headers, multipart, form, o.Nocolor)
// }

func ConstructCommand(parsedInput *gabs.Container) ([]string, string) {
	httpv := parsedInput.S("verb", "value")
	url := parsedInput.S("url", "value")
	jsonObj := parsedInput.S("details", "ip_data")
	headers := parsedInput.S("details", "headers")
	multipart := parsedInput.S("multipart", "value")
	multipartBool := false
	if multipart != nil {
		multipartBool = true
	}
	form := parsedInput.S("form", "value")
	formBool := form != nil

	res, stdinBody := assembleCmdString(httpv.Data().(string), url.Data().(string), jsonObj, headers, multipartBool, formBool)
	return res, stdinBody
}
