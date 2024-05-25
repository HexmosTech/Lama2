//go:build cli

// Package preprocess provides facilities to expand
// environment variables in `.l2` API files and return
// the contents
package preprocess

import (
	"strings"

	"github.com/HexmosTech/gabs/v2"
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

func ExpandJSON(block *gabs.Container, vm *goja.Runtime) {
	log.Debug().Str("JSON block to be expanded", block.String()).Msg("")
	dataBlock := block.S("details", "ip_data")
	if dataBlock == nil {
		return
	}
	dataBlockStr := dataBlock.String()
	dataBlockStr = ExpandEnv(dataBlockStr, vm)
	// dataBlockStr = escapeString(dataBlockStr)
	dataBlockStr = strings.ReplaceAll(dataBlockStr, "\n", "")
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
