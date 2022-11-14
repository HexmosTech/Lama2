package tests

import (
	"testing"

	"github.com/HexmosTech/lama2/cmdexec"
	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/rs/zerolog/log"
)

func TestExec(t *testing.T) {
	opts := lama2cmd.Opts{}
	l2File := "../elfparser/ElfTestSuite/y_0012_varjson_multipart.l2"
	// s, _ := FileToString()
	// l2File := "../elfparser/ElfTestSuite/y_0014_at_equal_ambiguity.l2"
	// l2File := "../elfparser/ElfTestSuite/y_0000_basic_get.l2"
	pp := preprocess.PreprocessLamaFile(l2File)
	lp := parser.NewLama2Parser()
	res, e := lp.Parse(pp)
	if e != nil {
		t.Fatalf("Error on parsing")
	}
	r2 := cmdgen.ConstructCommand(res, &opts)
	log.Debug().Str("Generated command", r2)
	r3 := cmdexec.ExecCommand(r2)
	log.Debug().Str("Execution result", r3)
}
