// server.go
package l2lsp

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/HexmosTech/lama2/l2lsp/request"
)

func init() {
}

func StartLspServer() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		handleInput(scanner.Text(), writer)
	}
}

func handleInput(input string, writer *bufio.Writer) {

	var rpcRequest request.JSONRPCRequest
	if err := json.Unmarshal([]byte(input), &rpcRequest); err != nil {
		return
	}

	rpcResponse := HandleMethod(rpcRequest)
	if responseData, err := json.Marshal(rpcResponse); err != nil {
	} else {
		writer.WriteString(string(responseData) + "\n")
		writer.Flush()
	}
}
