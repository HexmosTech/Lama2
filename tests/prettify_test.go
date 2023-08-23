package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/prettify"
)

func TestPrettifyMultiStage(t *testing.T) {
	lamaInputAPIFile := "../examples/0010_ugly_json/ugly_json_source1.l2"
	lamaOutputAPIFile := "../examples/0010_ugly_json/ugly_json_target1.l2"
	lamaOutputVerifyFile := "../examples/0010_ugly_json/ugly_json_target1_verify.l2"
	apiContent, _ := os.ReadFile(lamaInputAPIFile)
	fmt.Println(apiContent)
	p := parser.NewLama2Parser()
	parsedAPI, _ := p.Parse(string(apiContent))
	fmt.Println(parsedAPI)
	prettify.Prettify(parsedAPI, p.Context, p.MarkRange, string(apiContent), lamaOutputAPIFile)

	outputByte, _ := os.ReadFile(lamaOutputAPIFile)
	verifyByte, _ := os.ReadFile(lamaOutputVerifyFile)

	outputStr := string(outputByte)
	verifyStr := string(verifyByte)

	if err := os.Remove(lamaOutputAPIFile); err != nil {
		t.Error(err)
	}

	if verifyStr != outputStr {
		t.Errorf("Multistage prettification output not as expected")
	}
}

func TestPrettifySingleStage(t *testing.T) {
	lamaInputAPIFile := "../examples/0010_ugly_json/ugly_json_source2.l2"
	lamaOutputAPIFile := "../examples/0010_ugly_json/ugly_json_target2.l2"
	lamaOutputVerifyFile := "../examples/0010_ugly_json/ugly_json_target2_verify.l2"
	apiContent, _ := os.ReadFile(lamaInputAPIFile)
	fmt.Println(apiContent)
	p := parser.NewLama2Parser()
	parsedAPI, _ := p.Parse(string(apiContent))
	fmt.Println(parsedAPI)
	prettify.Prettify(parsedAPI, p.Context, p.MarkRange, string(apiContent), lamaOutputAPIFile)

	outputByte, _ := os.ReadFile(lamaOutputAPIFile)
	verifyByte, _ := os.ReadFile(lamaOutputVerifyFile)

	outputStr := string(outputByte)
	verifyStr := string(verifyByte)

	if err := os.Remove(lamaOutputAPIFile); err != nil {
		t.Error(err)
	}

	if verifyStr != outputStr {
		t.Errorf("Multistage prettification output not as expected")
	}
}
