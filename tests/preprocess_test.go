package tests

import (
	"testing"

	"github.com/HexmosTech/lama2/preprocess"
	// "github.com/rs/zerolog/log"
)

func TestPreprocessBasic(_ *testing.T) {
	op, _ := preprocess.LamaFile("../elfparser/ElfTestSuite/env1/sample.l2")
	// log.Debug().Str("Preprocessed string", op).Msg("")
}
