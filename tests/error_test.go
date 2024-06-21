package tests

import (
	"testing"

	"github.com/HexmosTech/lama2/lama2cmd"
	"github.com/HexmosTech/lama2/preprocess"
)

func TestWrongLama2FileFormat(t *testing.T) {
	ipArgs := []string{
		"l2",
		"xyz",
	}
	defer func() {
		if r := recover(); r != nil {
			// This code will execute if GetLamaFileAsString panics
			t.Logf("Recovered from panic: %v", r)
		}
	}()
	o := *lama2cmd.GetAndValidateCmd(ipArgs)
	_, err := preprocess.GetLamaFileAsString(o.Positional.LamaAPIFile)
	if err == nil {
		t.Error("Expected failure on faulty file format")
	} else {
		t.Log("Correctly failed on faulty file format")
	}
}
