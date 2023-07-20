package tests

import (
	// "encoding/json"
	"os/exec"
	"testing"
	"fmt"
	"bytes"
	"encoding/json"
)

type EnvData struct {
	Src string `json:"src"`
	Val string `json:"val"`
}

func runL2CommandAndParseJSON(t *testing.T, cmdArgs ...string)  {
	cmd := exec.Command("../build/l2", cmdArgs...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Execute the command
	err := cmd.Run()
	if err != nil {
		// Handle the error if needed
		fmt.Printf("Error running l2 command: %v\n", err)
	}

	// Retrieve the captured stdout
	stdoutOutput := stdout.String()

	fmt.Println(stdoutOutput)

	// Convert the stdoutOutput string to []byte slice
	outputBytes := []byte(stdoutOutput)

	envMap := make(map[string]EnvData)
	err = json.Unmarshal(outputBytes, &envMap)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON: %v\nOutput:\n%s", err, stdoutOutput)
	}
	// Example assertion: Check the "AHOST" key
	if ahost, ok := envMap["AHOST"]; !ok {
		t.Error("Expected 'AHOST' key in the JSON, but it was not found")
	} else {
		// Example assertion: Check the "AHOST" src and val values
		if ahost.Src != "l2env" {
			t.Errorf(`Expected "src" value to be "l2env" for "AHOST", but got: %v`, ahost.Src)
		}
		if ahost.Val != "`echo http://httpbin.org`" {
			t.Errorf("Expected \"val\" value to be \"`echo http://httpbin.org`\" for \"AHOST\", but got: %v", ahost.Val)
		}
	}
	// Example assertion: Check the "BHOST" key
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


func TestL2EnvCommand(t *testing.T) {
	cmdArgs := []string{"-e", "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"}
	runL2CommandAndParseJSON(t, cmdArgs...)
}


func TestL2EnvCommandVerbose(t *testing.T) {
	cmdArgs := []string{"-ev", "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"}
	 runL2CommandAndParseJSON(t, cmdArgs...)
}