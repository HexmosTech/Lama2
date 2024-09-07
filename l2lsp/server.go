package l2lsp

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/HexmosTech/lama2/l2lsp/request"
	"github.com/rs/zerolog/log"
)

// LSPServer defines the interface for starting an LSP server
type LSPServer interface {
	StartLspServer()
}

var lspServer LSPServer

// SetLSPServer sets the LSP server implementation
func SetLSPServer(server LSPServer) {
	lspServer = server
}

func init() {
}

// DefaultLSPServer is the default implementation of LSPServer
type DefaultLSPServer struct {}

func (s *DefaultLSPServer) StartLspServer() {
	log.Info().Msg("Started process")

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		handleInput(scanner.Text(), writer)
	}
}

func handleInput(input string, writer *bufio.Writer) {
	log.Info().Msgf("Received input: %s", input)
	var rpcRequest request.JSONRPCRequest
	if err := json.Unmarshal([]byte(input), &rpcRequest); err != nil {
		log.Error().Err(err).Msg("Error decoding JSON-RPC request")
		return
	}

	rpcResponse := HandleMethod(rpcRequest)
	if responseData, err := json.Marshal(rpcResponse); err != nil {
		log.Error().Err(err).Msg("Error encoding JSON-RPC response")
	} else {
		writer.WriteString(string(responseData) + "\n")
		writer.Flush()
	}
}
