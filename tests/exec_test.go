package tests

import (
	"testing"

	controller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
)

func TestExec(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/y_0012_varjson_multipart.l2"
	fdir := "../elfparser/ElfTestSuite/"
	cmd := []string{"l2", fpath}
	opts := lama2cmd.GetAndValidateCmd(cmd)
	s, _ := FileToString(fpath)
	lp := parser.NewLama2Parser()
	res, e := lp.Parse(s)
	if e != nil {
		t.Fatalf("Error on parsing")
	}
	controller.HandleParsedFile(res, opts, fdir)
	// log.Debug().Strs("Generated command", r2).Msg("")
	// r3, _ := cmdexec.ExecCommand(r2, body, apiDir)
	// log.Debug().Str("Execution result", r3).Msg("")
}
