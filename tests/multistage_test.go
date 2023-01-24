package tests

import (
	"fmt"
	"os"
	"testing"

	controller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/parser"
)

func TestMultiStageCount(t *testing.T) {
	// How many total blocks are there?
	// How many processor blocks are there?
	// How many requestor blocks are there?

	// parse file
	apiContent, _ := os.ReadFile("../examples/0009_processor_basic/0009_processor_basic.l2")
	p := parser.NewLama2Parser()
	parsedAPI, _ := p.Parse(string(apiContent))
	fmt.Println(parsedAPI)

	// get parsed blocks
	parsedAPIblocks := controller.GetParsedAPIBlocks(parsedAPI)
	fmt.Println(parsedAPIblocks)

	procCount := 0
	reqCount := 0
	// loop through
	for _, block := range parsedAPIblocks {
		blockType := block.S("type").Data().(string)
		if blockType == "processor" {
			procCount += 1
		} else if blockType == "Lama2File" {
			reqCount += 1
		}
	}

	// get statistics

	if procCount != 2 {
		t.Fatalf("Expected processor block count = 2")
	}

	if reqCount != 2 {
		t.Fatalf("Expected requestor block count = 2")
	}
}
