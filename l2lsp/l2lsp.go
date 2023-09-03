package l2lsp

import (
	"bufio"
	"encoding/json"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var isShutdownRequested bool

type InitializeParams struct {
	ProcessID             *int64             `json:"processId"`
	ClientInfo            *ClientInfo        `json:"clientInfo,omitempty"`
	Locale                *string            `json:"locale,omitempty"`
	RootPath              *string            `json:"rootPath,omitempty"`
	RootURI               *string            `json:"rootUri,omitempty"`
	InitializationOptions *interface{}       `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities"`
	Trace                 *string            `json:"trace,omitempty"`
	WorkspaceFolders      *[]WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

type LSPServer struct {
	IsShutdownRequested bool
}

type ClientInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

type TextDocumentClientCapabilities struct {
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`
	Hover      *HoverClientCapabilities      `json:"hover,omitempty"`
	// ... Add other text document capabilities as needed
}
 
type CompletionClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type HoverClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type WorkspaceFolder struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
}

type ClientCapabilities struct {
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
	// ... Add other capability sections (e.g., workspace) as needed
}

type ServerCapabilities struct {
	TextDocumentSync int                `json:"textDocumentSync"`
	Completion       *CompletionOptions `json:"completion,omitempty"`
	HoverProvider    bool               `json:"hoverProvider,omitempty"`
	// ... Add other capabilities as needed
}

type CompletionOptions struct {
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
	ResolveProvider   bool     `json:"resolveProvider,omitempty"`
}

type JSONRPCRequest struct {
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Params  *InitializeParams `json:"params,omitempty"`
	JSONRPC string            `json:"jsonrpc"`
}

type JSONRPCResponse struct {
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
	JSONRPC string      `json:"jsonrpc"`
	Error   interface{} `json:"error,omitempty"`
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func StartLspServer() {
	log.Info().Msg("Started process")

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		input := scanner.Text()

		var rpcRequest JSONRPCRequest
		err := json.Unmarshal([]byte(input), &rpcRequest)
		if err != nil {
			log.Error().Err(err).Msg("Error decoding JSON-RPC request")
			continue
		}

		// Log the received method
		log.Info().Str("Method", rpcRequest.Method).Msg("Received JSON-RPC method")

		// Handle request and prepare response
		rpcResponse := handleRequest(rpcRequest)

		responseData, err := json.Marshal(rpcResponse)
		if err != nil {
			log.Error().Err(err).Msg("Error encoding JSON-RPC response")
			continue
		}

		// Write response to stdout
		writer.WriteString(string(responseData) + "\n")
		writer.Flush()
	}
}

func handleRequest(request JSONRPCRequest) JSONRPCResponse {
	var response JSONRPCResponse

	if isShutdownRequested && request.Method != "exit" {
		// Invalid Request
		return InvalidReqAfterShutdown(request)
	}

	switch request.Method {

	case "initialize":
		// 1 indicates full synchronization
		response = Initilize(request)

	case "shutdown":
		response = ShutDown(request)

	case "exit":
		if isShutdownRequested {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

	default:
		// Respond with an error for unhandled methods
		// Method not found
		response = DefaultResp(request)
	}

	return response
}

func InvalidReqAfterShutdown(request JSONRPCRequest) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Error: map[string]string{
			"code":    "-32600",
			"message": "Invalid request after shutdown.",
		},
	}
}

func DefaultResp(request JSONRPCRequest) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Error: map[string]string{
			"code":    "-32601",
			"message": "Method not supported by the server.",
		},
	}
}

func ShutDown(request JSONRPCRequest) JSONRPCResponse {
	isShutdownRequested = true
	return JSONRPCResponse{
		ID:      request.ID,
		Result:  nil,
		JSONRPC: "2.0",
	}
}


func Initilize(request JSONRPCRequest) JSONRPCResponse {
	log.Info().Msg("LSP initialized")

	serverCapabilities := ServerCapabilities{
		TextDocumentSync: ,
	}

	if request.Params != nil && request.Params.Capabilities.TextDocument != nil {
		if request.Params.Capabilities.TextDocument.Completion != nil {
			serverCapabilities.Completion = &CompletionOptions{
				TriggerCharacters: []string{".", ",", "("},
				ResolveProvider:   true,
			}
		}
		if request.Params.Capabilities.TextDocument.Hover != nil {
			serverCapabilities.HoverProvider = true
		}
	}

	return JSONRPCResponse{
		ID: request.ID,
		Result: map[string]interface{}{
			"capabilities": serverCapabilities,
		},
		JSONRPC: "2.0",
	}
}

// {"id": 1, "method": "sayHi", "params": {}, "jsonrpc": "2.0"}
// {"jsonrpc": "2.0","id": 1,"method": "initialize","params": {  "processId": null,  "clientInfo": {    "name": "MyEditor",    "version": "1.0.0"  },  "rootUri": "file:///path/to/workspace",  "capabilities": {    "workspace": {      "applyEdit": true,      "workspaceEdit": {        "documentChanges": true      },      "didChangeConfiguration": {        "dynamicRegistration": true      }    },    "textDocument": {      "publishDiagnostics": {        "relatedInformation": true      }    }  }}}
//  All the events come in between these things.
// {  "jsonrpc": "2.0",  "id": 1,  "method": "shutdown"}
// {  "jsonrpc": "2.0",  "method": "exit"}
