// env_command_test.go
package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	testutils "github.com/HexmosTech/lama2/tests/utils"
)

func TestForEmptyRelevantString(t *testing.T) {
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	absPath, err := testutils.GetAbsolutePath(fpath)
	if err != nil {
		t.Fatalf("Failed to get the absolute path: %v", err)
	}
	// t.Fatalf("  absolute path: %v", absPath)

	relevantSearchString := ""

	request := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"suggest/environmentVariables","params":{"textDocument":{"uri":"file://%s"},"position":{"line":1,"character":45},"relevantSearchString":"%s"}}`, absPath, relevantSearchString)
	_, err = stdin.Write([]byte(request + "\n"))
	if err != nil {
		t.Fatalf("Failed to write to LSP server stdin: %v", err)
	}

	buffer := make([]byte, 2048)
	n, err := stdout.Read(buffer)
	if err != nil {
		t.Fatalf("Failed to read from LSP server stdout: %v", err)
	}

	var rawResponse RawJSONRPCResponse
	err = json.Unmarshal(buffer[:n], &rawResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP raw response: %v", err)
	}

	if rawResponse.ID != 1 {
		t.Fatalf("Expected response ID to be 1, got %v", rawResponse.ID)
	}
	if rawResponse.JSONRPC != "2.0" {
		t.Fatalf("Expected jsonrpc version to be 2.0, got %v", rawResponse.JSONRPC)
	}

	// Parse the Result into a map
	var envMap map[string]testutils.EnvData
	err = json.Unmarshal(rawResponse.Result, &envMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP response result: %v", err)
	}

	// Use the helper functions to check the AHOST and BHOST values
	checkAHost(t, envMap)
	checkBHost(t, envMap)
}

func TestL2SuggestEnvForNoL2Config(t *testing.T) {
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()
	fpath := "../elfparser/ElfTestSuite/no_l2config/api/y_0021_no_l2config.l2"
	absPath, err := testutils.GetAbsolutePath(fpath)

	if err != nil {
		t.Fatalf("Failed to get the absolute path: %v", err)
	}
	relevantSearchString := ""

	request := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"suggest/environmentVariables","params":{"textDocument":{"uri":"file://%s"},"position":{"line":1,"character":45},"relevantSearchString":"%s"}}`, absPath, relevantSearchString)
	_, err = stdin.Write([]byte(request + "\n"))
	if err != nil {
		t.Fatalf("Failed to write to LSP server stdin: %v", err)
	}

	buffer := make([]byte, 2048)
	n, err := stdout.Read(buffer)
	if err != nil {
		t.Fatalf("Failed to read from LSP server stdout: %v", err)
	}

	var rawResponse RawJSONRPCResponse
	err = json.Unmarshal(buffer[:n], &rawResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP raw response: %v", err)
	}

	if rawResponse.ID != 1 {
		t.Fatalf("Expected response ID to be 1, got %v", rawResponse.ID)
	}
	if rawResponse.JSONRPC != "2.0" {
		t.Fatalf("Expected jsonrpc version to be 2.0, got %v", rawResponse.JSONRPC)
	}

	// Parse the Result into a map
	var envMap map[string]testutils.EnvData
	err = json.Unmarshal(rawResponse.Result, &envMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP response result: %v", err)
	}

	// Use the helper functions to check the AHOST and BHOST values
	checkAHost(t, envMap)
}

func TestL2RelevantEnvForAString(t *testing.T) {
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	absPath, err := testutils.GetAbsolutePath(fpath)
	if err != nil {
		t.Fatalf("Failed to get the absolute path: %v", err)
	}
	relevantSearchString := "A"

	request := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"suggest/environmentVariables","params":{"textDocument":{"uri":"file://%s"},"position":{"line":1,"character":45},"relevantSearchString":"%s"}}`, absPath, relevantSearchString)
	_, err = stdin.Write([]byte(request + "\n"))
	if err != nil {
		t.Fatalf("Failed to write to LSP server stdin: %v", err)
	}

	buffer := make([]byte, 2048)
	n, err := stdout.Read(buffer)
	if err != nil {
		t.Fatalf("Failed to read from LSP server stdout: %v", err)
	}

	var rawResponse RawJSONRPCResponse
	err = json.Unmarshal(buffer[:n], &rawResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP raw response: %v", err)
	}

	if rawResponse.ID != 1 {
		t.Fatalf("Expected response ID to be 1, got %v", rawResponse.ID)
	}
	if rawResponse.JSONRPC != "2.0" {
		t.Fatalf("Expected jsonrpc version to be 2.0, got %v", rawResponse.JSONRPC)
	}

	// Parse the Result into a map
	var envMap map[string]testutils.EnvData
	err = json.Unmarshal(rawResponse.Result, &envMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP response result: %v", err)
	}

	// Use the helper functions to check the AHOST and BHOST values
	// Check the "AHOST" key is present
	checkAHost(t, envMap)

	// Check the "BHOST" key is absent
	checkBHostDoesNotExist(t, envMap)
}

func TestL2RelevantEnvForBString(t *testing.T) {
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()
	fpath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
	absPath, err := testutils.GetAbsolutePath(fpath)
	if err != nil {
		t.Fatalf("Failed to get the absolute path: %v", err)
	}
	relevantSearchString := "B"

	request := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"suggest/environmentVariables","params":{"textDocument":{"uri":"file://%s"},"position":{"line":1,"character":45},"relevantSearchString":"%s"}}`, absPath, relevantSearchString)
	_, err = stdin.Write([]byte(request + "\n"))
	if err != nil {
		t.Fatalf("Failed to write to LSP server stdin: %v", err)
	}

	buffer := make([]byte, 2048)
	n, err := stdout.Read(buffer)
	if err != nil {
		t.Fatalf("Failed to read from LSP server stdout: %v", err)
	}

	var rawResponse RawJSONRPCResponse
	err = json.Unmarshal(buffer[:n], &rawResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP raw response: %v", err)
	}

	if rawResponse.ID != 1 {
		t.Fatalf("Expected response ID to be 1, got %v", rawResponse.ID)
	}
	if rawResponse.JSONRPC != "2.0" {
		t.Fatalf("Expected jsonrpc version to be 2.0, got %v", rawResponse.JSONRPC)
	}

	// Parse the Result into a map
	var envMap map[string]testutils.EnvData
	err = json.Unmarshal(rawResponse.Result, &envMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP response result: %v", err)
	}

	// Use the helper functions to check the AHOST and BHOST values
	// Check the "BHOST" key is present
	checkBHost(t, envMap)

	// Check the "AHOST" key is absent
	checkAHostDoesNotExist(t, envMap)
}

func TestL2EnvWithoutL2config(t *testing.T) {
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()
	fpath := "../elfparser/ElfTestSuite/no_l2env/api/y_0022_no_l2env.l2"

	absPath, err := testutils.GetAbsolutePath(fpath)
	// t.Fatalf(string(absPath))

	if err != nil {
		t.Fatalf("Failed to get the absolute path: %v", err)
	}
	relevantSearchString := ""

	request := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"suggest/environmentVariables","params":{"textDocument":{"uri":"file://%s"},"position":{"line":1,"character":45},"relevantSearchString":"%s"}}`, absPath, relevantSearchString)
	_, err = stdin.Write([]byte(request + "\n"))
	if err != nil {
		t.Fatalf("Failed to write to LSP server stdin: %v", err)
	}

	buffer := make([]byte, 2048)
	n, err := stdout.Read(buffer)
	if err != nil {
		t.Fatalf("Failed to read from LSP server stdout: %v", err)
	}

	var rawResponse RawJSONRPCResponse
	err = json.Unmarshal(buffer[:n], &rawResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP raw response: %v", err)
	}

	if rawResponse.ID != 1 {
		t.Fatalf("Expected response ID to be 1, got %v", rawResponse.ID)
	}
	if rawResponse.JSONRPC != "2.0" {
		t.Fatalf("Expected jsonrpc version to be 2.0, got %v", rawResponse.JSONRPC)
	}

	// Parse the Result into a map
	var envMap map[string]testutils.EnvData
	err = json.Unmarshal(rawResponse.Result, &envMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP response result: %v", err)
	}

	// Use the helper functions to check the AHOST and BHOST values
	// Check the "AHOST" key is present
	checkBHost(t, envMap)
}

// checkAHost checks the "AHOST" key in the JSON map
func checkAHost(t *testing.T, envMap map[string]testutils.EnvData) {
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
func checkBHost(t *testing.T, envMap map[string]testutils.EnvData) {
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

func checkBHostDoesNotExist(t *testing.T, envMap map[string]testutils.EnvData) {
	if _, ok := envMap["BHOST"]; ok {
		t.Error("Expected 'BHOST' key not to be present in the JSON, but it was found")
	}
}

func checkAHostDoesNotExist(t *testing.T, envMap map[string]testutils.EnvData) {
	if _, ok := envMap["AHOST"]; ok {
		t.Error("Expected 'AHOST' key not to be present in the JSON, but it was found")
	}
}
