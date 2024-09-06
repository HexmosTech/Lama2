package response

import "github.com/HexmosTech/lama2/l2lsp/request"

func CreateExecuteCommandResponse(req request.JSONRPCRequest, response interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result:  response,
	}
}

// {"jsonrpc":"2.0","id":12,"method":"executeCommand","params":{"filePath":"/home/lince/hexmos/apihub/apihub/fb_backend_v3/search.l2"}}