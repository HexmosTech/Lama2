package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"

	controller "github.com/HexmosTech/lama2/controller"
	"github.com/HexmosTech/lama2/parser"
	testutils "github.com/HexmosTech/lama2/tests/utils"
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
			procCount++
		} else if blockType == "Lama2File" {
			reqCount++
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

func TestUnquotedVars(t *testing.T) {
	apiContent, _ := os.ReadFile("../examples/0024_vars_multistage/0024_vars_multistage.l2")
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
			procCount++
		} else if blockType == "Lama2File" {
			reqCount++
		}
	}

	// get statistics

	if procCount != 1 {
		t.Fatalf("Expected processor block count = 1")
	}

	if reqCount != 2 {
		t.Fatalf("Expected requestor block count = 2")
	}
}

func TestMultiStageResults(t *testing.T) {
	fpath := "../examples/0024_vars_multistage/0024_vars_multistage_acess_token.l2"
	cmdArgs := []string{fpath}

	output, err := testutils.RunL2CommandAndGetOutput(cmdArgs...)
	fmt.Println(output)
	if err != nil {
		t.Errorf("Error running L2 command: %v", err)
		return
	}

	expectedOutputPart := "\"authenticated\": true"
	if !strings.Contains(output, expectedOutputPart) {
		t.Errorf("Expected output to contain %q, but got %q", expectedOutputPart, output)
	}
}
