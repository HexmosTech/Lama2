package tests

import (
	"testing"

	"github.com/HexmosTech/lama2/codegen"
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
