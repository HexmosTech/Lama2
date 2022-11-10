package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/HexmosTech/lama2/lama2cmd"
)

func TestCmdBasic(t *testing.T) {
	ipArgs := []string{
		"l2",
		"-v",
		"../elfparser/ElfTestSuite/y_0000_basic_get.l2",
	}
	o := *lama2cmd.GetAndValidateCmd(ipArgs)

	expected := lama2cmd.Opts{
		Verbose: []bool{true}, Sort: false, Prettify: false, Nocolor: false, Positional: struct {
			LamaAPIFile string "required:\"yes\""
		}{LamaAPIFile: "../elfparser/ElfTestSuite/y_0000_basic_get.l2"}}

	if !reflect.DeepEqual(o, expected) {
		fmt.Println(reflect.DeepEqual(o.Verbose, expected.Verbose))
		t.Errorf("Unsuccessful parsing basic CLI options")
	}
}
