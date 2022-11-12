package contoller

import (
	"fmt"
	"os"

	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/rs/zerolog/log"
)

func Process() {
	o := lama2cmd.GetAndValidateCmd(os.Args)
	apiContent := preprocess.PreprocessLamaFile(o.Positional.LamaAPIFile)
	p := parser.NewLama2Parser()
	parsedAPI, e := p.Parse(apiContent)
	if e != nil {
		log.Fatal().
			Str("Type", "Controller").
			Str("LamaFile", o.Positional.LamaAPIFile).
			Str("Error", e.Error()).
			Msg(fmt.Sprint("Parse Error"))
	}
	fmt.Println(parsedAPI)
	cmdStr := cmdgen.ConstructCommand(parsedAPI, o)
	res := cmdexec.ExecCommand(cmdStr)
	fmt.Println(res)
}
