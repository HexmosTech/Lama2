package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/HexmosTech/lama2/cmdgen"
	contoller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/parser"
	// "github.com/rs/zerolog/log"
)

// matchFiles, _ := getDataFiles("../elfparser/ElfTestSuite", "y_0012_varjson_multipart.l2")

func TestConstruct(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/y_0012_varjson_multipart.l2"
	s, _ := FileToString(fpath)
	cmd := []string{"l2", fpath}
	opts := lama2cmd.GetAndValidateCmd(cmd)
	fmt.Println(os.Getwd())
	lp := parser.NewLama2Parser()
	res, e := lp.Parse(s)
	if e != nil {
		t.Fatalf("Error on parsing")
	}
	blocks := contoller.GetParsedAPIBlocks(res)
	for _, block := range blocks {
		r2, body := cmdgen.ConstructCommand(block, opts)
		// log.Debug().Strs("Constructed command", r2).Msg("")
		// log.Debug().Str("Constructed body: ", body).Msg("")
	}
}
