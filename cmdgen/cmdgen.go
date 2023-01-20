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
	"github.com/rs/zerolog/log"
)

func assembleCmdString(httpv string, url string, jsonObj *gabs.Container, headers *gabs.Container, multipart bool, o *lama2cmd.Opts) ([]string, string) {
	command := make([]string, 0)
	log.Info().
		Str("Type", "Construct Command").
		Str("httpv", httpv).
		Str("url", url).
		Bool("multipart", multipart).
		Msg(fmt.Sprint("Construct parameters"))

	log.Debug().
		Str("JSONObj", jsonObj.String()).
		Str("Headers", headers.String()).Msg("")

	var files *gabs.Container
	if multipart {
		if jsonObj.ExistsP("@files") {
			files = jsonObj.S("@files")
			log.Debug().Str("Files", files.String()).Msg("")
			jsonObj.Delete("@files")
			log.Trace().Str("Shortened JsonObj", jsonObj.String()).Msg("")
		}
	}

	jsonStr := ""
	if jsonObj != nil && !multipart {
		dst := &bytes.Buffer{}
		if err := json.Compact(dst, []byte(jsonObj.String())); err != nil {
			log.Fatal().
				Str("Error", err.Error()).
				Msg("Couldn't minify JSON")
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
	if o.Nocolor {
		command = append(command, "--pretty=none ")
	}
	if multipart {
		command = append(command, "--multipart ")
	}

	command = append(command, httpv+" ")
	command = append(command, url+" ")

	if multipart {
		for key, val := range jsonObj.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, "'"+key+"'='"+val.Data().(string)+"'  ")
		}
		for key, val := range files.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, "'"+key+"'@'"+val.Data().(string)+"'  ")
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
	return cleanCommand, jsonStr

	/*
		commandStr := strings.Join(command, "")
		fmt.Println("### ", commandStr)
		log.Info().Str("Generated command", commandStr)
		return commandStr
	*/
}

// ConstructCommand extracts the HTTP verb, url and other
// API file inputs, figures out the type of target command
// and finally generates a string representing the generated
// command
func ConstructCommand(parsedInput *gabs.Container, o *lama2cmd.Opts) ([]string, string) {
	log.Info().Str("ParsedInput", parsedInput.String()).Msg("")
	fmt.Println(parsedInput)
	httpv := parsedInput.S("verb", "value")
	url := parsedInput.S("url", "value")
	jsonObj := parsedInput.S("details", "ip_data")
	headers := parsedInput.S("details", "headers")
	multipart := parsedInput.S("multipart", "value")
	multipartBool := false
	if multipart != nil {
		multipartBool = true
	}

	res, stdinBody := assembleCmdString(httpv.Data().(string), url.Data().(string), jsonObj, headers, multipartBool, o)
	return res, stdinBody
}
