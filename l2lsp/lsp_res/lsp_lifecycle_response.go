package lsp_res

import "github.com/HexmosTech/lama2/l2lsp/lsp_req"

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
}

type ServerCapabilities struct {
	TextDocumentSync int                `json:"textDocumentSync"`
	Completion       *CompletionOptions `json:"completion,omitempty"`
	HoverProvider    bool               `json:"hoverProvider,omitempty"`
	SuggestL2Envs    bool               `json:"suggestL2Env,omitempty"`
}

type CompletionOptions struct {
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
	ResolveProvider   bool     `json:"resolveProvider,omitempty"`
}

func InvalidReqAfterShutdown(request lsp_req.JSONRPCRequest) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Error: &JSONRPCError{
			Code:    ErrInvalidAfterShutdown,
			Message: "Invalid request after shutdown.",
		},
	}
}
