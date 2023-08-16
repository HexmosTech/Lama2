// basic_commands_test.go
package tests

import (
	"strings"
	"testing"

	testutils "github.com/HexmosTech/lama2/tests/utils"
)

func TestNormalExecution(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/y_0000_basic_get.l2"
	cmdArgs := []string{fpath}
	output := testutils.RunL2CommandAndGetOutput(t, cmdArgs...)

	expectedOutputPart := "\"url\": \"http://httpbin.org/get\""
	if !strings.Contains(output, expectedOutputPart) {
		t.Errorf("Expected output to contain %q, but got %q", expectedOutputPart, output)
	}
}
