package l2lsp

import (
	"fmt"
	"net/http"
)

type LSPService struct{}

type LSPRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

func (h *LSPService) Process(r *http.Request, request *LSPRequest, response *string) error {
	fmt.Printf("Decoded request: %+v\n", request)

	// Log HTTP and LSP Methods.
	fmt.Printf("HTTP Method: %s\n", r.Method)
	fmt.Printf("LSP Method: %s\n", request.Method)

	*response = startBasicLSPServer(request.Method)
	return nil
}

func startBasicLSPServer(request string) string {
	fmt.Printf("Processing LSP Method: %s\n", request) // Logging the LSP method being processed

	var rsp string
	switch request {
	case "initialize":
		rsp = "{\"jsonrpc\": \"2.0\", \"result\": {\"capabilities\": {}}, \"id\": 1}"
		// Expand this with actual capabilities
	default:
		rsp = "{\"jsonrpc\": \"2.0\", \"error\": {\"code\": -32000, \"message\": \"Unknown method.\"}, \"id\": null}"
	}
	return rsp
}
