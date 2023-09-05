// lsp.go
package l2lsp

import (
	"bufio"
	"encoding/json"
	"path/filepath"
	"strings"

	"os"

	l2envpackege "github.com/HexmosTech/lama2/l2env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var isShutdownRequested bool

func init() {
	setupLogging()
}

func setupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func StartLspServer() {
	log.Info().Msg("Started process")

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		handleInput(scanner.Text(), writer)
	}
}

func handleInput(input string, writer *bufio.Writer) {
	log.Info().Msgf("Received input: %s", input)

	var rpcRequest JSONRPCRequest
	if err := json.Unmarshal([]byte(input), &rpcRequest); err != nil {
		log.Error().Err(err).Msg("Error decoding JSON-RPC request")
		return
	}

	rpcResponse := handleRequest(rpcRequest)
	if responseData, err := json.Marshal(rpcResponse); err != nil {
		log.Error().Err(err).Msg("Error encoding JSON-RPC response")
	} else {
		writer.WriteString(string(responseData) + "\n")
		writer.Flush()
	}
}

func handleRequest(request JSONRPCRequest) JSONRPCResponse {
	if isShutdownRequested && request.Method != "exit" {
		return InvalidReqAfterShutdown(request)
	}

	switch request.Method {
	case "initialize":
		return Initialize(request)
	case "shutdown":
		return Shutdown(request)
	case "exit":
		return handleExit()
	case "suggest/environmentVariables":
		return suggestEnvironmentVariables(request)
	default:
		return DefaultResp(request)
	}
}
func handleExit() JSONRPCResponse {
	exitCode := 1
	if isShutdownRequested {
		exitCode = 0
	}
	os.Exit(exitCode)
	return JSONRPCResponse{} // This line will not be reached but is added for completeness
}

func suggestEnvironmentVariables(request JSONRPCRequest) JSONRPCResponse {
	log.Info().Str("Method", request.Method).Interface("Params", request.Params)

	relevantSearchString := getRequestSearchString(request)
	uri := getRequestURI(request)
	if uri == "" {
		return ErrorResp(request, "Document uri is improper. Ex: 'file:///path/to/workspace/myapi.l2'")
	}
	parentFolder := filepath.Dir(uri)
	res := l2envpackege.ProcessEnvironmentVariables(relevantSearchString, parentFolder)
	return createEnvironmentVariablesResponse(request, res)
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

func getRequestSearchString(request JSONRPCRequest) string {
	if request.Params.RelevantSearchString != nil {
		return *request.Params.RelevantSearchString
	}
	return ""
}
func getRequestURI(request JSONRPCRequest) string {
	if request.Params.TextDocument.Uri == nil {
		return ""
	}
	uri := *request.Params.TextDocument.Uri
	if strings.HasPrefix(uri, "file://") {
		return uri[len("file://"):]
	}
	return ""
}
func createEnvironmentVariablesResponse(request JSONRPCRequest, res interface{}) JSONRPCResponse {
	return JSONRPCResponse{
		ID:      request.ID,
		JSONRPC: "2.0",
		Result:  res,
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

func Shutdown(request JSONRPCRequest) JSONRPCResponse {
	isShutdownRequested = true
	return JSONRPCResponse{
		ID:      request.ID,
		Result:  nil,
		JSONRPC: "2.0",
	}
}

func Initialize(request JSONRPCRequest) JSONRPCResponse {
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
