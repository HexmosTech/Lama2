package lsp_res

import "github.com/HexmosTech/lama2/l2lsp/lsp_req"

func CreateEnvironmentVariablesResp(request lsp_req.JSONRPCRequest, envs interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Result:  envs,
	}
}
