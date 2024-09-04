package l2lsp

import (
	"github.com/HexmosTech/lama2/l2lsp/methods"
	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/HexmosTech/lama2/l2lsp/response"
)

var isShutdownRequested bool

func HandleMethod(req request.JSONRPCRequest) response.JSONRPCResponse {
	if isShutdownRequested && req.Method != "exit" {
		return response.InvalidReqAfterShutdown(req)
	}

	switch req.Method {
	case "initialize":
		return methods.Initialize(req)

	case "shutdown":
		return methods.Shutdown(req, isShutdownRequested)

	case "exit":
		return methods.Exit(isShutdownRequested)

	case "suggest/environmentVariables":
		return methods.SuggestEnvironmentVariables(req)
	case "executeCommand":
		return methods.ExecuteCommand(req)

	default:
		return response.DefaultResp(req)
	}
}
