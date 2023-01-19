package tests

import (
	"testing"

	"github.com/HexmosTech/lama2/cmdgen"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
	"github.com/rs/zerolog/log"
)

// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0012_varjson_multipart.l2")

func TestConstruct(t *testing.T) {
	opts := lama2cmd.Opts{}
	s, _ := FileToString("../elfparser/ElfTestSuite/y_0012_varjson_multipart.l2")
	// s, _ := FileToString("../elfparser/ElfTestSuite/y_0014_at_equal_ambiguity.l2")
	lp := parser.NewLama2Parser()
	res, e := lp.Parse(s)
	if e != nil {
		t.Fatalf("Error on parsing")
	}
	r2 := cmdgen.ConstructCommand(res, &opts)
	log.Debug().Strs("Constructed command", r2).Msg("")
}
