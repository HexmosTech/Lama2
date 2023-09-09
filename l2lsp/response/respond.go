package response

import "github.com/HexmosTech/lama2/l2lsp/request"

// Utility function to create a general success response
func CreateSuccessResponse(requestID int, result interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      requestID,
		JSONRPC: "2.0",
		Result:  result,
	}
}

func DefaultResp(req request.JSONRPCRequest) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      req.ID,
		JSONRPC: "2.0",
		Error: &JSONRPCError{
			Code:    ErrMethodNotFound,
			Message: "Method not supported by the server. Method: " + req.Method,
		},
	}
}

// Utility function to create a general error response
func ErrorResp(req request.JSONRPCRequest, errorCode int, errorMsg string) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      req.ID,
		JSONRPC: "2.0",
		Error: &JSONRPCError{
			Code:    errorCode,
			Message: errorMsg,
		},
	}
}
