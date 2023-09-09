package response

type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JSONRPCResponse struct {
	ID      int           `json:"id"`
	Result  interface{}   `json:"result"`
	JSONRPC string        `json:"jsonrpc"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

const (
	ErrInvalidAfterShutdown = -32000
	ErrUnsupportedFeature   = -32001
	ErrInvalidURI           = -32002
	ErrUnexpectedURIScheme  = -32003
	ErrInvalidRequest       = -32600
	ErrMethodNotFound       = -32601
	ErrInternalError        = -32603
)
