package tests

import (
	"testing"

	"github.com/HexmosTech/lama2/preprocess"
)

func TestPreprocessBasic(t *testing.T) {
	op := preprocess.PreprocessLamaFile("../elfparser/ElfTestSuite/env1/sample.l2")
	// fmt.Println(op)
}
