// lsp_shutdown.go
package l2lsp

import (
	"os"

	"github.com/rs/zerolog/log"
)

var isShutdownRequested bool

func Initialize(request JSONRPCRequest) JSONRPCResponse {
	// {"jsonrpc": "2.0","id": 1,"method": "initialize","params": { "processId": null, "clientInfo": { "name": "MyEditor", "version": "1.0.0" }, "rootUri": "file:///path/to/workspace", "capabilities": { "workspace": { "applyEdit": true, "workspaceEdit": { "documentChanges": true },didChangeConfiguration": { "dynamicRegistration": true } }, "textDocument": { "publishDiagnostics": { "relatedInformation": true }}}}}
	log.Info().Msg("L2 LSP initialized")

	serverCapabilities := ServerCapabilities{
		TextDocumentSync: 0,
		SuggestL2Envs:    true,
	}

	return JSONRPCResponse{
		ID: request.ID,
		Result: map[string]interface{}{
			"capabilities": serverCapabilities,
		},
		JSONRPC: "2.0",
	}
}

func Shutdown(request JSONRPCRequest) JSONRPCResponse {
	// { "jsonrpc": "2.0", "id": 1, "method": "shutdown"}
	isShutdownRequested = true
	return JSONRPCResponse{
		ID:      request.ID,
		Result:  nil,
		JSONRPC: "2.0",
	}
}

func Exit() JSONRPCResponse {
	// { "jsonrpc": "2.0", "method": "exit"}
	exitCode := 1
	if isShutdownRequested {
		exitCode = 0
	}
	os.Exit(exitCode)
	return JSONRPCResponse{}
}
