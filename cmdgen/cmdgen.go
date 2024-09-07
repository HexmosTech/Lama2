// Package `cmdgen` provides an API to generate
// API request commands (by default based on HTTPie)
// based on the parsed API file contents and the `l2`
// command invocation parameters
package cmdgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/lama2/lama2cmd"
)

// The assembleCmdString function constructs a httpie
// command string for an HTTP request based on the provided
// HTTP method, URL, JSON object, headers, and options. It also
//
//	handle multipart and form data.
func assembleCmdString(httpv string, url string, jsonObj *gabs.Container, headers *gabs.Container, multipart bool, form bool, o *lama2cmd.Opts) ([]string, string) {
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
			fmt.Println(err)
		}
		jsonStr = dst.String()
	}

	/*
		if !multipart && jsonStr != "" {
			command = append(command, "echo '")
			command = append(command, jsonStr)
			command = append(command, "' |")
		}*/

	command = append(command, "ht ")
	if o != nil && o.Nocolor {
		command = append(command, "--pretty=none ")
	}
	if multipart || form {
		command = append(command, "--ignore-stdin", "--form")
	}

	command = append(command, httpv+" ")
	command = append(command, url+" ")

	if multipart {
		for key, val := range jsonObj.Data().(*gabs.Container).ChildrenMap() {
			keyValuePair := fmt.Sprintf("%s=%s", key, val.Data().(string))
			command = append(command, keyValuePair)
		}
		for key, val := range files.ChildrenMap() {
			command = append(command, key+"@"+val.Data().(string))
		}
	}

	if form {
		for key, val := range jsonObj.Data().(*gabs.Container).ChildrenMap() {
			keyValuePair := fmt.Sprintf("%s=%s", key, val.Data().(string))
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

// ConstructCommand extracts the HTTP verb, url and other
// API file inputs, figures out the type of target command
// and finally generates a string representing the generated
// command
func ConstructCommandHelper(parsedInput *gabs.Container) (string, string, *gabs.Container, *gabs.Container, bool, bool) {
	httpv := parsedInput.S("verb", "value")
	url := parsedInput.S("url", "value")
	jsonObj := parsedInput.S("details", "ip_data")
	headers := parsedInput.S("details", "headers")
	multipart := parsedInput.S("multipart", "value")
	form := parsedInput.S("form", "value")
	multipartBool := false
	if multipart != nil {
		multipartBool = true
	}
	formBool := form != nil
	return httpv.Data().(string), url.Data().(string), jsonObj, headers, multipartBool, formBool
}
