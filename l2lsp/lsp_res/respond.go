package lsp_res

import "github.com/HexmosTech/lama2/l2lsp/lsp_req"

// Utility function to create a general success response
func CreateSuccessResponse(requestID int, result interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      requestID,
		JSONRPC: "2.0",
		Result:  result,
	}
}

func DefaultResp(request lsp_req.JSONRPCRequest) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Error: &JSONRPCError{
			Code:    ErrMethodNotFound,
			Message: "Method not supported by the server. Method: " + request.Method,
		},
	}
}

// Utility function to create a general error response
func ErrorResp(request lsp_req.JSONRPCRequest, errorCode int, errorMsg string) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Error: &JSONRPCError{
			Code:    errorCode,
			Message: errorMsg,
		},
	}
}
