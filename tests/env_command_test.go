package tests

import (
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type EnvData struct {
	Src string `json:"src"`
	Val string `json:"val"`
}

func runL2CommandAndParseJSON(t *testing.T, cmdArgs ...string) {
	// Construct the relative path to the l2.go executable
	l2Path := filepath.Join("..", "..", "l2")

	cmd := exec.Command(l2Path, cmdArgs...)
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
}

func TestL2EnvCommand(t *testing.T) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current working directory:", err)
		return
	}

	// Print the current working directory
	fmt.Println("Current working directory:", wd)

	// Perform "ls" command on /home/runner/work/Lama2/Lama2/
	lama2Lama2Dir := "/home/runner/work/Lama2/Lama2/"
	fmt.Println("Contents of /home/runner/work/Lama2/Lama2/:")
	listFilesInDir(lama2Lama2Dir)

	// Your existing code to run the l2 command and parse JSON
	cmdArgs := []string{"-e", "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"}
	runL2CommandAndParseJSON(t, cmdArgs...)
}

func listFilesInDir(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
