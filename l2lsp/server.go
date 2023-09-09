// server.go
package l2lsp

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/HexmosTech/lama2/l2lsp/request"
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
