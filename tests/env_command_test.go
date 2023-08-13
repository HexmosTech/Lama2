package tests

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

func TestL2EnvCommand(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=", fpath}
	envMap := runL2CommandAndParseJSON(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)

	// Check the "BHOST" key
	checkBHost(t, envMap)
}

func TestL2EnvForAHOenv(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=A", fpath}
	envMap := runL2CommandAndParseJSON(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)

	// Check the "BHOST" key is absent
	checkBHostDoesNotExist(t, envMap)
}

func TestL2EnvForBHOenv(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=B", fpath}
	envMap := runL2CommandAndParseJSON(t, cmdArgs...)

	// Check the "BHOST" key is present
	checkBHost(t, envMap)

	// Check the "AHOST" key is absent
	checkAHostDoesNotExist(t, envMap)
}

func TestL2EnvCommandVerbose(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	cmdArgs := []string{"-e=", "-v", fpath}
	envMap := runL2CommandAndParseJSON(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)

	// Check the "BHOST" key is present
	checkBHost(t, envMap)
}

func TestL2EnvWithoutL2config(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/no_l2config/api/y_0021_no_l2config.l2"
	cmdArgs := []string{"-e=", fpath}
	envMap := runL2CommandAndParseJSON(t, cmdArgs...)

	// Check the "AHOST" key is present
	checkAHost(t, envMap)
}

func TestL2EnvWithoutL2env(t *testing.T) {
	fpath := "../elfparser/ElfTestSuite/no_l2env/api/y_0022_no_l2env.l2"
	cmdArgs := []string{"-e=", fpath}
	envMap := runL2CommandAndParseJSON(t, cmdArgs...)

	// Check the "BHOST" key is present
	checkBHost(t, envMap)
}

func runL2CommandAndParseJSON(t *testing.T, cmdArgs ...string) map[string]EnvData {
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

// checkL2BinaryExists checks if the l2 binary file exists in the specified path
func checkL2BinaryExists(l2BinPath string) error {
	// Check if the l2 binary file exists
	if _, err := os.Stat(l2BinPath); os.IsNotExist(err) {
		return fmt.Errorf("l2 binary not found in the build folder %s, please change the path", l2BinPath)
	}
	return nil
}

// checkAHost checks the "AHOST" key in the JSON map
func checkAHost(t *testing.T, envMap map[string]EnvData) {
	if ahost, ok := envMap["AHOST"]; !ok {
		t.Error("Expected 'AHOST' key in the JSON, but it was not found")
	} else {
		// Example assertion: Check the "AHOST" src and val values
		if ahost.Src != "l2env" {
			t.Errorf(`Expected "src" value to be "l2env" for "AHOST", but got: %v`, ahost.Src)
		}
		if ahost.Val != "`echo http://httpbin.org`" {
			t.Errorf(`Expected "val" value to be "echo http://httpbin.org" for "AHOST", but got: %v`, ahost.Val)
		}
	}
}

// checkBHost checks the "BHOST" key in the JSON map
func checkBHost(t *testing.T, envMap map[string]EnvData) {
	if bhost, ok := envMap["BHOST"]; !ok {
		t.Error("Expected 'BHOST' key in the JSON, but it was not found")
	} else {
		// Example assertion: Check the "BHOST" src and val values
		if bhost.Src != "l2configenv" {
			t.Errorf(`Expected "src" value to be "l2configenv" for "BHOST", but got: %v`, bhost.Src)
		}
		if bhost.Val != "https://httpbin.org" {
			t.Errorf(`Expected "val" value to be "https://httpbin.org" for "BHOST", but got: %v`, bhost.Val)
		}
	}
}

func checkBHostDoesNotExist(t *testing.T, envMap map[string]EnvData) {
	if _, ok := envMap["BHOST"]; ok {
		t.Error("Expected 'BHOST' key not to be present in the JSON, but it was found")
	}
}

func checkAHostDoesNotExist(t *testing.T, envMap map[string]EnvData) {
	if _, ok := envMap["AHOST"]; ok {
		t.Error("Expected 'AHOST' key not to be present in the JSON, but it was found")
	}
}
