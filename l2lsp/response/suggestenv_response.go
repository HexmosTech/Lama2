package response

import "github.com/HexmosTech/lama2/l2lsp/request"

func CreateEnvironmentVariablesResp(req request.JSONRPCRequest, envs interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      req.ID,
		JSONRPC: "2.0",
		Result:  envs,
	}
}

//  {"jsonrpc":"2.0","id":12,"method":"executeCommand","params":{"filePath":"/home/lince/hexmos/apihub/apihub/fb_backend_v3/search.l2"}}
