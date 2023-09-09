package l2lsp

import (
	"github.com/HexmosTech/lama2/l2lsp/lsp_req"
	"github.com/HexmosTech/lama2/l2lsp/lsp_res"
	"github.com/HexmosTech/lama2/l2lsp/methods"
)

var isShutdownRequested bool

func HandleMethod(request lsp_req.JSONRPCRequest) lsp_res.JSONRPCResponse {
	if isShutdownRequested && request.Method != "exit" {
		return lsp_res.InvalidReqAfterShutdown(request)
	}

	switch request.Method {
	case "initialize":
		return methods.Initialize(request)

	case "shutdown":
		return methods.Shutdown(request, isShutdownRequested)

	case "exit":
		return methods.Exit(isShutdownRequested)

	case "suggest/environmentVariables":
		return methods.SuggestEnvironmentVariables(request)

	default:
		return lsp_res.DefaultResp(request)
	}
}
