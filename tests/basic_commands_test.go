// basic_commands_test.go
package tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/rs/zerolog/log"
)

func runL2CommandAndGetOutput(t *testing.T, cmdArgs ...string) string {

	// Get the full path to the l2 binary
	l2BinPath := "../build/l2"

	// Check if the l2 binary file exists
	if err := checkL2BinaryExists(l2BinPath); err != nil {
		t.Error(err)
		return ""
	}

	// Code to run the l2 command
	cmd := exec.Command(l2BinPath, cmdArgs...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		log.Error().Str("Error", stderr.String()).Msg("Error running l2 command")
		t.Errorf("Error running l2 command: %v\n", err)
		return ""
	}

	// Retrieve the captured stdout
	stdoutOutput := stdout.String()
	log.Info().Str("Test env_command", stdoutOutput).Msg("output from command")
	return stdoutOutput
}

func TestNormalExecution(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/y_0000_basic_get.l2"
	cmdArgs := []string{fpath}
	output := runL2CommandAndGetOutput(t, cmdArgs...)

	expectedOutputPart := "\"url\": \"http://httpbin.org/get\""
	if !strings.Contains(output, expectedOutputPart) {
		t.Errorf("Expected output to contain %q, but got %q", expectedOutputPart, output)
	}
}
