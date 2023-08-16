package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/rs/zerolog/log"
)

type EnvData struct {
	Src string `json:"src"`
	Val string `json:"val"`
}

func RunL2CommandAndGetOutput(t *testing.T, cmdArgs ...string) string {

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

func RunL2CommandAndParseJSON(t *testing.T, cmdArgs ...string) map[string]EnvData {
	// Get the full path to the l2 binary
	l2BinPath := "../build/l2"

	// Check if the l2 binary file exists
	if err := checkL2BinaryExists(l2BinPath); err != nil {
		t.Error(err)
		return make(map[string]EnvData)
	}

	// Your existing code to run the l2 command and parse JSON
	cmd := exec.Command(l2BinPath, cmdArgs...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Execute the command
	err := cmd.Run()
	if err != nil {
		// Handle the error if needed
		t.Errorf("Error running l2 command: %v\n", err)
		return make(map[string]EnvData)
	}

	// Retrieve the captured stdout
	stdoutOutput := stdout.String()

	log.Debug().Str("Test env_command", stdoutOutput).Msg("output from command")

	// Convert the stdoutOutput string to []byte slice
	outputBytes := []byte(stdoutOutput)

	envMap := make(map[string]EnvData)
	err = json.Unmarshal(outputBytes, &envMap)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON env: %v\nOutput:\n%s", err, stdoutOutput)
	}

	return envMap
}

func checkL2BinaryExists(l2BinPath string) error {
	// Check if the l2 binary file exists
	if _, err := os.Stat(l2BinPath); os.IsNotExist(err) {
		return fmt.Errorf("l2 binary not found in the build folder %s, please change the path", l2BinPath)
	}
	return nil
}
