// lsp_suggestEnv.go
package l2lsp

import (
	"path/filepath"
	"strings"

	l2envpackege "github.com/HexmosTech/lama2/l2env"
	"github.com/rs/zerolog/log"
)

func getRelevantSearchString(request JSONRPCRequest) string {
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

func suggestEnvironmentVariables(request JSONRPCRequest) JSONRPCResponse {
	// { "jsonrpc": "2.0", "id": 2, "method": "suggest/environmentVariables", "params": { "textDocument": { "uri": "file:///home/lovestaco/repos/Lama2/elfparser/ElfTestSuite/root_variable_override/api/y_0020_root_override.l2" }, "position": { "line": 1, "character": 2 }, "relevantSearchString": "" } }
	log.Info().Str("Method", request.Method).Interface("Params", request.Params)

	relevantSearchString := getRelevantSearchString(request)
	uri := getRequestURI(request)
	if uri == "" {
		return ErrorResp(request, "Document uri is improper. Ex: 'file:///path/to/workspace/myapi.l2'")
	}
	parentFolder := filepath.Dir(uri)
	res := l2envpackege.ProcessEnvironmentVariables(relevantSearchString, parentFolder)
	return createEnvironmentVariablesResponse(request, res)
}
