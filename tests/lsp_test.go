// lsp_test.go
package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"testing"

	"github.com/HexmosTech/lama2/l2lsp"
	testutils "github.com/HexmosTech/lama2/tests/utils"
)

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

func TestLSPInitialization(t *testing.T) {
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()

	request := `{ "jsonrpc": "2.0", "id": 1, "method": "initialize", "params": { "processId": 449931, "rootUri": "file:///home/lovestaco/repos/apihub", "workspace": { "workspaceFolders": { "supported": false, "changeNotifications": false } }, "clientInfo": { "name": "Visual Studio Code", "version": "1.81.1" } } }`
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

	// Now unmarshal the raw Result into InitializeResult
	var initResult l2lsp.InitializeResult
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
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()

	// Constructing the shutdown request
	request := `{ "jsonrpc": "2.0", "id": 1, "method": "shutdown" }`
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
	stdin, stdout, err := startLSPServer()
	if err != nil {
		t.Fatalf("Failed to start LSP server: %v", err)
	}
	defer stdin.Close()

	// Constructing the exit request
	request := `{ "jsonrpc": "2.0", "method": "exit" }`
	_, err = stdin.Write([]byte(request + "\n"))
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
