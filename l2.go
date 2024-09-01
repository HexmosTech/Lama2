//go:build cli

package main

import (
	contoller "github.com/HexmosTech/lama2/controller"
)

var version string
var lamawasm string

func main() {
	if len(version) == 0 {
		version = "vUnset"
	}
	contoller.Process(version)
}