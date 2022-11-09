package contoller

import (
	"os"

	"github.com/HexmosTech/lama2/lama2cmd"
)

func Process() {
	lama2cmd.GetAndValidateCmd(os.Args)
}
