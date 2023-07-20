package tests

import (
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type EnvData struct {
	Src string `json:"src"`
	Val string `json:"val"`
}

func runL2CommandAndParseJSON(t *testing.T, cmdArgs ...string) {
// Store the original working directory
originalPath, err := os.Getwd()
if err != nil {
		fmt.Println("Error getting the current working directory:", err)
		return
}

// Change the current working directory to lama2Path
l2BinPath := "/home/runner/work/Lama2/Lama2/build"
err = os.Chdir(l2BinPath)
if err != nil {
		fmt.Println("Error changing the current working directory:", err)
		return
}

// Print the current working directory to verify
wd, err := os.Getwd()
if err != nil {
		fmt.Println("Error getting the current working directory:", err)
		return
}
fmt.Println("Current working directory after changing:", wd)

// Now let's check if the l2 binary is present in the current working directory
fileInfo, err := os.Stat("./l2")
if err != nil {
		fmt.Println("Error checking l2 binary:", err)
		return
}
fmt.Println("Is l2 binary present:", fileInfo.Mode().IsRegular())

// Change back to the original working directory after debugging
defer func() {
		err := os.Chdir(originalPath)
		if err != nil {
				fmt.Println("Error changing back to the original working directory:", err)
		}
}()

// Your existing code to run the l2 command and parse JSON
cmd := exec.Command("./l2", cmdArgs...) // Use "./l2" as the executable name

var stdout bytes.Buffer
cmd.Stdout = &stdout

// Execute the command
err = cmd.Run() // Uncomment this line to execute the "l2" command
if err != nil {
		// Handle the error if needed
		fmt.Printf("Error running l2 command: %v\n", err)
		return
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
	err = os.Chdir(originalPath)
		if err != nil {
			fmt.Println("Error changing back to the original working directory:", err)
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
 
 

	// Your existing code to run the l2 command and parse JSON
	cmdArgs := []string{"-e", "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"}
	runL2CommandAndParseJSON(t, cmdArgs...)
}
 


func listFilesInDir(dirPath string) {
	fmt.Println("Contents of", dirPath+":")
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}