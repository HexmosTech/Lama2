// Package preprocess provides facilities to expand
// environment variables in `.l2` API files and return
// the contents
package preprocess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/godotenv"
	"github.com/HexmosTech/lama2/utils"
	"github.com/dop251/goja"
	"github.com/rs/zerolog/log"
)

func ProcessVarsInBlock(block *gabs.Container, vm *goja.Runtime) {
	ExpandURL(block, vm)
	ExpandHeaders(block, vm)
	ExpandJSON(block, vm)
}

func ExpandHeaders(block *gabs.Container, vm *goja.Runtime) {
	headerMap := block.S("details", "headers")
	log.Debug().Str("HeaderMap", headerMap.String()).Msg("")
	if headerMap == nil {
		return
	}
	newHeaderMap := gabs.New()
	for k, v := range headerMap.ChildrenMap() {
		log.Trace().Strs("Header pair", []string{k, " = ", v.String()}).Msg("")
		key := ExpandEnv(k, vm)
		val := ExpandEnv(v.Data().(*gabs.Container).Data().(string), vm)
		valWrap := gabs.New()
		valWrap.Set(val)
		newHeaderMap.Set(valWrap, key)
	}
	block.Delete("details", "headers")
	block.Set(newHeaderMap, "details", "headers")
	log.Debug().Str("Expanded Header block", block.String()).Msg("")
}

func ExpandURL(block *gabs.Container, vm *goja.Runtime) {
	b := block.S("url", "value").Data().(string)
	log.Debug().Str("Url block", b).Msg("")
	url := ExpandEnv(b, vm)
	block.Delete("url", "value")
	block.Set(url, "url", "value")
}

func debugOp(str string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("String written to file successfully.")
}

func escapeString(input string) string {
	output, err := json.Marshal(input)
	if err != nil {
		// handle error
	}
	return string(output)
}

func ExpandJSON(block *gabs.Container, vm *goja.Runtime) {
	log.Debug().Str("JSON block to be expanded", block.String()).Msg("")
	dataBlock := block.S("details", "ip_data")
	if dataBlock == nil {
		return
	}
	dataBlockStr := dataBlock.String()
	dataBlockStr = ExpandEnv(dataBlockStr, vm)
	dataBlockStr = escapeString(dataBlockStr)
	log.Debug().Str("Expanded JSON data block", dataBlockStr).Msg("")
	processedBlock, err := gabs.ParseJSON([]byte(dataBlockStr))
	if err != nil {
		log.Error().Str("Preprocess JSON block issue", "").Msg("")
		return
	}
	block.Delete("details", "ip_data")
	block.Set(processedBlock, "details", "ip_data")
	log.Debug().Str("Processed JSON block", block.String()).Msg("")
}

func LoadElfEnv(l2path string) {
	err := godotenv.Load(l2path)
	if err != nil {
		log.Info().Str("Type", "Preprocess").Msg("Didn't find l2.env in the API directory")
	}
}

func GetLamaFileAsString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		log.Fatal().Str("Type", "Preprocess").Msg(fmt.Sprint("Couldn't read: ", path))
	}
	return string(b)
}

// LamaFile takes in a path to an API file.
// It moves into the API file directory, reads the
// API contents, loads the `l2.env` file if available,
// and finally substitutes environment vars in the API contents
// Once done, it reverts back to the original directory,
// and returns the processed l2 file.
func LamaFile(inputFile string) (string, string) {
	content := GetLamaFileAsString(inputFile)
	_, dir, _ := utils.GetFilePathComponents(inputFile)
	oldDir, _ := os.Getwd()

	utils.ChangeWorkingDir(dir)
	LoadElfEnv("l2.env")
	content = os.ExpandEnv(content)
	utils.ChangeWorkingDir(oldDir)

	return content, dir
}
