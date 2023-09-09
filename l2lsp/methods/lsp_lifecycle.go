// lsp_lifecycle.go
package methods

import (
	"os"

	"github.com/HexmosTech/lama2/l2lsp/lsp_req"
	"github.com/HexmosTech/lama2/l2lsp/lsp_res"
	"github.com/rs/zerolog/log"
)

func Initialize(request lsp_req.JSONRPCRequest) lsp_res.JSONRPCResponse {
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

	serverCapabilities := lsp_res.ServerCapabilities{
		TextDocumentSync: 0,
		SuggestL2Envs:    true,
		HoverProvider:    false,
	}
	res := map[string]interface{}{
		"capabilities": serverCapabilities,
	}
	return lsp_res.CreateSuccessResponse(request.ID, res)
}

func Shutdown(request lsp_req.JSONRPCRequest, isShutdownRequested bool) lsp_res.JSONRPCResponse {
	/*
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "shutdown"
		}
	*/

	log.Info().Msg("L2 LSP shutdown requested")

	isShutdownRequested = true
	return lsp_res.CreateSuccessResponse(request.ID, nil)
}

func Exit(isShutdownRequested bool) lsp_res.JSONRPCResponse {
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
	return lsp_res.JSONRPCResponse{}
}
