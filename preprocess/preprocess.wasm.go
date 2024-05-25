//go:build wasm

// Package preprocess provides facilities to expand
// environment variables in `.l2` API files and return
// the contents
package preprocess

import (
	"fmt"
	"strings"

	"github.com/HexmosTech/gabs/v2"
	// "github.com/rs/zerolog/log"
)

func ProcessVarsInBlock(block *gabs.Container) {
	ExpandURL(block)
	ExpandHeaders(block)
	ExpandJSON(block)
}

func ExpandHeaders(block *gabs.Container) {
	headerMap := block.S("details", "headers")
	if headerMap == nil {
		return
	}
	newHeaderMap := gabs.New()
	for k, v := range headerMap.ChildrenMap() {
		key := ExpandEnv(k)
		val := ExpandEnv(v.Data().(*gabs.Container).Data().(string))
		valWrap := gabs.New()
		valWrap.Set(val)
		newHeaderMap.Set(valWrap, key)
	}
	block.Delete("details", "headers")
	block.Set(newHeaderMap, "details", "headers")
}

func ExpandURL(block *gabs.Container) {
	b := block.S("url", "value").Data().(string)
	url := ExpandEnv(b)
	block.Delete("url", "value")
	block.Set(url, "url", "value")
}

func ExpandJSON(block *gabs.Container) {
	dataBlock := block.S("details", "ip_data")
	if dataBlock == nil {
		return
	}
	dataBlockStr := dataBlock.String()
	dataBlockStr = ExpandEnv(dataBlockStr)
	dataBlockStr = strings.ReplaceAll(dataBlockStr, "\n", "")
	processedBlock, err := gabs.ParseJSON([]byte(dataBlockStr))
	if err != nil {
		fmt.Println("Preprocess JSON block issue)")
		return
	}
	block.Delete("details", "ip_data")
	block.Set(processedBlock, "details", "ip_data")
}
