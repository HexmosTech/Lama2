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
	// log.Logger = log.Output(io.Discard)

	// os.Stdout.Write([]byte{})
	// os.Stderr.Write([]byte{})

	lama2cmd.SetLSPServer(&l2lsp.DefaultLSPServer{})
	controller.Process(version)
}