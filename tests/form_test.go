package tests

import (
	"strings"
	"testing"

	testutils "github.com/HexmosTech/lama2/tests/utils"
)

func TestFormPositive(t *testing.T) {
	fpath := "formtests/test_form_positive.l2"
	cmdArgs := []string{fpath}

	output, err := testutils.RunL2CommandAndGetOutput(cmdArgs...)
	if err != nil {
		t.Errorf("Error running L2 command: %v", err)
		return
	}

	expectedOutputPart := "\"form\": {\n    \"'first'\": \"'second'\", \n    \"'third'\": \"'fourth'\"\n  }"
	if !strings.Contains(output, expectedOutputPart) {
		t.Errorf("Expected output to contain %q, but got %q", expectedOutputPart, output)
	}
}

func TestFormNegative(t *testing.T) {
	fpath := "formtests/test_form_negative.l2"
	cmdArgs := []string{fpath}

	_, err := testutils.RunL2CommandAndGetOutput(cmdArgs...)
	if err == nil {
		t.Errorf("Expected failure on faulty code")
		return
	}
}
