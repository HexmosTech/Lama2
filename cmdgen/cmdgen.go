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

func assembleCmdString(httpv string, url string, jsonObj *gabs.Container, headers *gabs.Container, multipart bool, o *lama2cmd.Opts) string {
	command := make([]string, 0)
	log.Info().
		Str("Type", "Construct Command").
		Str("httpv", httpv).
		Str("url", url).
		Bool("multipart", multipart).
		Msg(fmt.Sprint("Construct parameters"))

	log.Debug().
		Str("JSONObj", jsonObj.String()).
		Str("Headers", headers.String())

	var files *gabs.Container
	if multipart {
		if jsonObj.ExistsP("@files") {
			files = jsonObj.S("@files")
			log.Debug().Str("Files", files.String())
			jsonObj.Delete("@files")
			log.Trace().Str("Shortened JsonObj", jsonObj.String())
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

	if !multipart && jsonStr != "" {
		command = append(command, "echo '")
		command = append(command, jsonStr)
		command = append(command, "' |")
	}

	command = append(command, "http ")
	if o.Nocolor {
		command = append(command, "--pretty=none ")
	}
	if multipart {
		command = append(command, "--multipart ")
	}

	command = append(command, httpv+" ")
	command = append(command, url+" ")

	if multipart {
		// fmt.Println("key123 = ", files)
		for key, val := range jsonObj.Data().(*gabs.Container).ChildrenMap() {
			fmt.Println("processing now", val)
			command = append(command, "'"+key+"'='"+val.Data().(string)+"'  ")
		}
		for key, val := range files.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, "'"+key+"'@'"+val.Data().(string)+"'  ")
		}
	}

	if headers != nil {
		// fmt.Println("header = ", headers)
		for key, val := range headers.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, "'"+key+":"+val.Data().(*gabs.Container).Data().(string)+"'  ")
		}

	}

	commandStr := strings.Join(command, "")
	log.Info().Str("Generated command", commandStr)
	return commandStr
}

func ConstructCommand(parsedInput *gabs.Container, o *lama2cmd.Opts) string {
	log.Info().Str("ParsedInput", parsedInput.String())
	httpv := parsedInput.S("value", "verb", "value")
	url := parsedInput.S("value", "url", "value")
	jsonObj := parsedInput.S("value", "details", "ip_data")
	headers := parsedInput.S("value", "details", "headers")
	multipart := parsedInput.S("value", "multipart", "value")
	multipartBool := false
	if multipart != nil {
		multipartBool = true
	}
	fmt.Println("httpv", httpv, httpv.Data().(string))
	fmt.Println("url", url, url.Data().(string))

	res := assembleCmdString(httpv.Data().(string), url.Data().(string), jsonObj, headers, multipartBool, o)
	return res
}
