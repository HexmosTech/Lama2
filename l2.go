package main

import (
	controller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/l2lsp"
	"github.com/HexmosTech/lama2/lama2cmd"
)

var version string
var lamawasm string

func main() {
	if len(version) == 0 {
		version = "vUnset"
	}
	lama2cmd.SetLSPServer(&l2lsp.DefaultLSPServer{})
	controller.Process(version)
}