package cmdgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog/log"
)

func assembleCmdString(httpv string, url string, jsonObj *gabs.Container, headers *gabs.Container, multipart bool) string {

	command := make([]string, 0)
	log.Info().
		Str("Type", "Construct Command").
		Str("httpv", httpv).
		Str("url", url).
		Bool("multipart", multipart).
		Msg(fmt.Sprint("Construct parameters"))

	fmt.Println("jsonObj", jsonObj)
	fmt.Println("headers", headers)

	var files *gabs.Container
	if multipart {
		if jsonObj.ExistsP("@files") {
			files = jsonObj.S("@files")
			fmt.Println("files", files)
			jsonObj.Delete("@files")
			fmt.Println(jsonObj)
		}
	}

	jsonStr := ""
	if jsonObj != nil {
		dst := &bytes.Buffer{}
		if err := json.Compact(dst, []byte(jsonObj.String())); err != nil {
			log.Fatal().
				Str("Error", err.Error()).
				Msg("Couldn't minify JSON")
		}
		jsonStr = dst.String()
		fmt.Println("##", jsonStr)
	}

	command = append(command, "echo '")
	command = append(command, jsonStr)
	command = append(command, "' | http ")

	command = append(command, httpv+" ")
	command = append(command, url+" ")

	if multipart {
		command = append(command, "--multipart ")

		fmt.Println("key123 = ", files)
		for key, val := range files.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, key+"@"+val.Data().(string)+"  ")
		}
	}

	if headers != nil {
		fmt.Println("header = ", headers)
		for key, val := range headers.Data().(*gabs.Container).ChildrenMap() {
			command = append(command, key+":"+val.Data().(*gabs.Container).Data().(string)+"  ")
		}

	}

	commandStr := strings.Join(command, "")
	fmt.Println(commandStr)
	return commandStr
}

func ConstructCommand(parsedInput *gabs.Container) string {
	fmt.Println("==", parsedInput.S("value", "details", "headers"))

	fmt.Println(parsedInput)
	httpv := parsedInput.S("value", "verb", "value")
	url := parsedInput.S("value", "url", "value")
	jsonObj := parsedInput.S("value", "details", "ip_data")
	headers := parsedInput.S("value", "details", "headers")
	multipart := parsedInput.S("value", "multipart", "value")
	multipartBool := false
	if multipart != nil {
		multipartBool = true
	}
	res := assembleCmdString(httpv.Data().(string), url.Data().(string), jsonObj, headers, multipartBool)
	return res
}
