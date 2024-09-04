package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/HexmosTech/lama2/codegen"
	"github.com/HexmosTech/lama2/parser"
	"github.com/HexmosTech/lama2/preprocess"
	"github.com/HexmosTech/lama2/utils"
)

func TestSplitLangLib1(t *testing.T) {
	lang, lib := codegen.SplitLangLib("python.requests")
	if lang != "python" {
		t.Errorf("Expected lang = python")
	}

	if lib != "requests" {
		t.Errorf("Expected lib = requests")
	}
}

func TestSplitLangLib2(t *testing.T) {
	lang, lib := codegen.SplitLangLib("python")
	if lang != "python" {
		t.Errorf("Expected lang = python")
	}

	if lib != "" {
		t.Errorf("Expected lib = ''")
	}
}

func TestGenerateMultiStage( _ *testing.T) {
	l2Path := "../examples/0009_processor_basic/0009_processor_basic.l2"
	apiContent, _ := os.ReadFile(l2Path)
	_, dir, _ := utils.GetFilePathComponents(l2Path)
	nowPwd, _ := os.Getwd()
	utils.ChangeWorkingDir(dir)
	defer utils.ChangeWorkingDir(nowPwd)
	preprocess.LoadEnvironments(dir)
	p := parser.NewLama2Parser()
	parsedAPI, _ := p.Parse(string(apiContent))
	fmt.Println(parsedAPI)

	codegen.GenerateTargetCode("python.requests", parsedAPI)
}
