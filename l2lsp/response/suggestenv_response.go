package response

import "github.com/HexmosTech/lama2/l2lsp/request"

func CreateEnvironmentVariablesResp(req request.JSONRPCRequest, envs interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      req.ID,
		JSONRPC: "2.0",
		Result:  envs,
	}
}
