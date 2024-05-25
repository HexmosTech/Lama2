//go:wasm cli

package codegen

import (
	_ "embed"

	"github.com/HexmosTech/lama2/cmdexec"
)

func initialize() {
	globalVM = cmdexec.GetJSVm()
}
