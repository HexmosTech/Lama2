// lsp_lifecycle.go
package methods

import (
	"os"

	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
	"github.com/rs/zerolog/log"
)

func Initialize(req request.JSONRPCRequest) response.JSONRPCResponse {
	/*
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "initialize",
			"params": {
				"processId": null,
				"clientInfo": {
					"name": "MyEditor",
					"version": "1.0.0"
				},
				"rootUri": "file:///path/to/workspace",
				"capabilities": {
					"workspace": {
						"applyEdit": true,
						"workspaceEdit": {
							"documentChanges": true
						},
						"didChangeConfiguration": {
							"dynamicRegistration": true
						}
					},
					"textDocument": {
						"publishDiagnostics": {
							"relatedInformation": true
						}
					}
				}
			}
		}
	*/
	log.Info().Msg("L2 LSP initialized")

	serverCapabilities := response.ServerCapabilities{
		TextDocumentSync: 0,
		SuggestL2Envs:    true,
		HoverProvider:    false,
	}
	res := map[string]interface{}{
		"capabilities": serverCapabilities,
	}
	return response.CreateSuccessResponse(req.ID, res)
}

// Shutdown is a way to gracefully terminate the server.
// The server can perform cleanup operations, like closing open files, releasing resources, or saving state.
func Shutdown(req request.JSONRPCRequest, isShutdownRequested bool) response.JSONRPCResponse {
	/*
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "shutdown"
		}
	*/

	log.Info().Msg("L2 LSP shutdown requested")

	isShutdownRequested = true
	return response.CreateSuccessResponse(req.ID, nil)
}

// Exit terminates the server process, with the exit code depending on whether a shutdown was requested.
func Exit(isShutdownRequested bool) response.JSONRPCResponse {
	/*
		{
			"jsonrpc": "2.0",
			"method": "exit"
		}
	*/

	log.Info().Msg("L2 LSP exit requested")

	exitCode := 1
	if isShutdownRequested {
		exitCode = 0
	}
	os.Exit(exitCode)
	return response.JSONRPCResponse{}
}
