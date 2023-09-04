package l2lsp

import (
	"bufio"
	"encoding/json"
	"strings"

	"os"

	l2envpackege "github.com/HexmosTech/lama2/l2env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var isShutdownRequested bool

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
		log.Info().Msg("Received input: " + input)

		var rpcRequest JSONRPCRequest
		err := json.Unmarshal([]byte(input), &rpcRequest)
		if err != nil {
			log.Error().Err(err).Msg("Error decoding JSON-RPC request")
			continue
		}

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
	case "suggest/environmentVariables":
		log.Info().Str("Method", request.Method).Interface("Params", request.Params)

		// Check if relevantSearchString is present in params and get the value if available
		var relevantSearchString string
		if request.Params.RelevantSearchString != nil {
			relevantSearchString = *request.Params.RelevantSearchString
		} else {
			relevantSearchString = ""
		}
		log.Info().Str("Method", request.Method).Msg("relevantSearchString: " + relevantSearchString)

		var uri string
		if request.Params.TextDocument.Uri == nil {
			msg := "No document uri found"
			return ErrorResp(request, msg)
		} else {
			uri = *request.Params.TextDocument.Uri
			if strings.HasPrefix(uri, "file://") {
				uri = uri[len("file://"):]
			} else {
				msg := "Document uri is improper. Ex: 'file:///path/to/workspace/myapi.l2' "
				return ErrorResp(request, msg)
			}
		}

		res := l2envpackege.ProcessEnvironmentVariables(relevantSearchString, uri)
		// log.Info().Msgf("Processed Environment Variables Result: %v", res)

		return JSONRPCResponse{
			ID:      request.ID,
			JSONRPC: "1555.0",
			Result:  res,
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

func ErrorResp(request JSONRPCRequest, msg string) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Error: map[string]string{
			"code":    "-32601",
			"message": msg,
		},
	}
}

func DefaultResp(request JSONRPCRequest) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Error: map[string]string{
			"code":    "-32601",
			"message": "Method not supported by the server. Method: " + request.Method,
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
	log.Info().Msg("L2 LSP initialized")

	serverCapabilities := ServerCapabilities{
		TextDocumentSync: 0,
		SuggestL2Envs:    true,
	}

	return JSONRPCResponse{
		ID: request.ID,
		Result: map[string]interface{}{
			"capabilities": serverCapabilities,
		},
		JSONRPC: "2.0",
	}
}

// {"jsonrpc": "2.0","id": 1,"method": "initialize","params": {  "processId": null,  "clientInfo": {    "name": "MyEditor",    "version": "1.0.0"  },  "rootUri": "file:///path/to/workspace",  "capabilities": {    "workspace": {      "applyEdit": true,      "workspaceEdit": {        "documentChanges": true      },      "didChangeConfiguration": {        "dynamicRegistration": true      }    },    "textDocument": {      "publishDiagnostics": {        "relatedInformation": true      }    }  }}}
// {"jsonrpc":"2.0","id":1,"method":"suggest/environmentVariables","params":{"textDocument":{"uri":"file:////home/lovestaco/repos/apihub/karma/karma_admin/login/one_logout.l2"},"position":{"line":1,"character":45},"relevantSearchString":"karma"}}
// {  "jsonrpc": "2.0",  "id": 1,  "method": "shutdown"}
// {  "jsonrpc": "2.0",  "method": "exit"}
