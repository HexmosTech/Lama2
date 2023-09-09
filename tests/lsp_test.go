// lsp_test.go
package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/HexmosTech/lama2/l2lsp/lsp_res"
	testutils "github.com/HexmosTech/lama2/tests/utils"
)

type LSPResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Capabilities struct {
			TextDocumentSync int  `json:"textDocumentSync"`
			SuggestL2Env     bool `json:"suggestL2Env"`
		} `json:"capabilities"`
	} `json:"result"`
}

type RawJSONRPCResponse struct {
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	JSONRPC string          `json:"jsonrpc"`
	Error   interface{}     `json:"error,omitempty"`
}

var stdin io.WriteCloser
var stdout io.ReadCloser

func TestMain(m *testing.M) {
	// Setup: Start the LSP server
	var err error
	stdin, stdout, err = startLSPServer()
	if err != nil {
		panic(err)
	}

	// Run all tests
	exitVal := m.Run()

	// Teardown: Close the server (this will send an "exit" to your LSP server)
	stdin.Write([]byte(`{ "jsonrpc": "2.0", "method": "exit" }` + "\n"))
	stdin.Close()

	// Exit with the result of the tests
	os.Exit(exitVal)
}

func startLSPServer() (io.WriteCloser, io.ReadCloser, error) {
	l2BinPath, err := testutils.GetLocalL2BinaryPath()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get L2 binary path: %v", err)
	}

	cmd := exec.Command(l2BinPath, "--lsp")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get stdin pipe: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get stdout pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("Failed to start LSP server: %v", err)
	}

	return stdin, stdout, nil
}

// func TestAbsolutePathUtility(t *testing.T) {
// 	relPath := "../elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2"
// 	absPath, err := testutils.GetAbsolutePath(relPath)
// 	if err != nil {
// 		t.Fatalf("Failed to get absolute path: %v", err)
// 	}

// 	// To view the result in logs, intentionally fail the test
// 	t.Fatalf("Absolute path: %s", absPath)
// }

func TestLSPInitialization(t *testing.T) {
	request := `{ "jsonrpc": "2.0", "id": 1, "method": "initialize", "params": { "processId": 449931, "rootUri": "file:///home/lovestaco/repos/apihub", "workspace": { "workspaceFolders": { "supported": false, "changeNotifications": false } }, "clientInfo": { "name": "Visual Studio Code", "version": "1.81.1" } } }`
	_, err := stdin.Write([]byte(request + "\n"))
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

	// Now unmarshal the raw Result into InitializeResult
	var initResult lsp_res.InitializeResult
	err = json.Unmarshal(rawResponse.Result, &initResult)
	if err != nil {
		t.Fatalf("Failed to unmarshal LSP result: %v", err)
	}

	// Validate response
	if rawResponse.ID != 1 {
		t.Fatalf("Expected response ID to be 1, got %v", rawResponse.ID)
	}
	if rawResponse.JSONRPC != "2.0" {
		t.Fatalf("Expected jsonrpc version to be 2.0, got %v", rawResponse.JSONRPC)
	}

	// Continue your validations on initResult
	if !initResult.Capabilities.SuggestL2Envs {
		t.Fatalf("Expected SuggestL2Envs to be true")
	}
}

func TestLSPShutdown(t *testing.T) {
	// Constructing the shutdown request
	request := `{ "jsonrpc": "2.0", "id": 1, "method": "shutdown" }`
	_, err := stdin.Write([]byte(request + "\n"))
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

	// Validate response
	if rawResponse.ID != 1 {
		t.Fatalf("Expected response ID to be 1, got %v", rawResponse.ID)
	}
	if rawResponse.JSONRPC != "2.0" {
		t.Fatalf("Expected jsonrpc version to be 2.0, got %v", rawResponse.JSONRPC)
	}
	if string(rawResponse.Result) != "null" {
		t.Fatalf("Expected result to be 'null', got %v", string(rawResponse.Result))
	}
}

func TestLSPExit(t *testing.T) {
	// Constructing the exit request
	request := `{ "jsonrpc": "2.0", "method": "exit" }`
	_, err := stdin.Write([]byte(request + "\n"))
	if err != nil {
		t.Fatalf("Failed to write to LSP server stdin: %v", err)
	}

	// Since the server shuts down on receiving the "exit" request,
	// attempting to read from stdout should result in an EOF (End of File) error
	buffer := make([]byte, 2048)
	_, err = stdout.Read(buffer)
	if err != io.EOF {
		t.Fatalf("Expected an EOF error after sending 'exit', but received: %v", err)
	}
}