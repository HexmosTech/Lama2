// lsp.go
package l2lsp

import (
	"bufio"
	"encoding/json"

	"os"

	outputmanager "github.com/HexmosTech/lama2/outputManager"
	"github.com/rs/zerolog/log"
)

func init() {
	outputmanager.ConfigureZeroLog("INFO")
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
		return Exit()
	case "suggest/environmentVariables":
		return SuggestEnvironmentVariables(request)
	default:
		return DefaultResp(request)
	}
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
